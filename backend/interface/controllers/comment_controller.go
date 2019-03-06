package controllers

import (
	"errors"
	"matilda/backend/domain"
	"matilda/backend/domain/entity"
	"matilda/backend/interface/database"
	"matilda/backend/interface/gorilla_mux"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
	"net/http"
	"strconv"
	"time"

	errors2 "github.com/pkg/errors"
	"google.golang.org/appengine/datastore"

	"google.golang.org/appengine"
)

type CommentController struct {
	MuxCommentInterceptor       usecase.MuxCommentInterceptor
	DatastoreCommentInterceptor usecase.DatastoreCommentInterceptor
	LogCommentInterceptor       usecase.LogCommentInterceptor
}

func NewCommentController(gorillaMuxHandler gorilla_mux.GorillaMuxHandler, datastoreHandler database.DatastoreHandler, logHandler logging.LogHandler) *CommentController {
	return &CommentController{
		MuxCommentInterceptor: usecase.MuxCommentInterceptor{
			MuxCommentRepository: &gorilla_mux.MuxCommentRepository{
				GorillaMuxHandler: gorillaMuxHandler,
			},
		},
		DatastoreCommentInterceptor: usecase.DatastoreCommentInterceptor{
			DatastoreCommentRepository: &database.DatastoreCommentRepository{
				DatastoreHandler: datastoreHandler,
			},
		},
		LogCommentInterceptor: usecase.LogCommentInterceptor{
			LogCommentRepository: &logging.LogCommentRepository{
				LogHandler: logHandler,
			},
		},
	}
}

func (controller *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)

	sub, err := getUserID(r, ctx)
	if err != nil {
		return appErrorf(err, "getUserID() error: %v", err)
	}
	if sub == "" {
		err := errors.New("could not get sub")
		return appErrorf(err, "controller.getUserID() error: %v", err)
	}

	u := &entity.User{UserID: sub}

	userKey := controller.DatastoreCommentInterceptor.GetKey(r, u)
	if userKey == nil {
		err := errors.New("could not get key")
		return appErrorf(err, "controller.DatastoreCommentInterceptor.GetKey() error: %v", err)
	}

	var c entity.Comment
	err = mappingJsonToStruct(r, &c)
	if err != nil {
		return appErrorf(err, "mappingJsonToStruct() error: %v", err)
	}

	movieID, err := strconv.Atoi(controller.MuxCommentInterceptor.Get(r, "movieID"))
	if err != nil {
		return appErrorf(err, "strconv.Atoi() error: %v", err)
	}
	c.MovieID = movieID
	c.UserKey = userKey
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	_, err = controller.DatastoreCommentInterceptor.Put(r, &c)
	if err != nil {
		return appErrorf(err, "controller.DatastoreCommentInterceptor.Put() error: %v", err)
	}

	err = setResponseWriter(w, 202, err)
	if err != nil {
		return appErrorf(err, "setResponseWriter() error: %v", err)
	}

	controller.LogCommentInterceptor.LogInfo(ctx, "CreateComment() user_id: %v", sub)
	return nil
}

func (controller *CommentController) GetComments(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)

	sub, err := getUserID(r, ctx)
	if err != nil {
		return appErrorf(err, "getUserID() error: %v", err)
	}

	commentKeys, err := controller.getCommentKeys(r)
	if err != nil {
		return appErrorf(err, "controller.getCommentKeys() error: %v", err)
	}

	var commentIDs []int64
	for _, k := range commentKeys {
		commentIDs = append(commentIDs, k.IntID())
	}

	var entityComments []*entity.Comment
	for i := range commentIDs {
		commentID := commentIDs[i]
		entityComments = append(entityComments, &entity.Comment{
			CommentID: commentID,
		})
	}

	entityComments, err = controller.getMultiComments(r, entityComments)
	if err != nil {
		return appErrorf(err, "controller.getMultiComments() error: %v", err)
	}

	var comments []*domain.Comment
	var mine bool
	for i := range entityComments {
		if entityComments[i].UserKey.StringID() == sub {
			mine = true
		}
		comments = append(comments, &domain.Comment{
			CommentID:   entityComments[i].CommentID,
			CommentText: entityComments[i].CommentText,
			Mine:        mine,
			CreatedAt:   entityComments[i].CreatedAt,
			UpdatedAt:   entityComments[i].UpdatedAt,
			UserKey:     entityComments[i].UserKey,
			UserID:      entityComments[i].UserKey.StringID(),
		})
	}

	userKeys := controller.getCommentUserKeys(comments)

	var userIDs []string
	for i := range userKeys {
		userIDs = append(userIDs, userKeys[i].StringID())
	}

	var users []*entity.User
	for i := range userIDs {
		userID := userIDs[i]
		users = append(users, &entity.User{
			UserID: userID,
		})
	}

	users, err = controller.getMultiUsers(r, users)
	if err != nil {
		return appErrorf(err, "controller.getMultiUsers() error: %v", err)
	}

	for i := range users {
		comments[i].Name = users[i].Name
		comments[i].IconPath = users[i].IconPath
	}

	err = setResponseWriter(w, 200, &comments)
	if err != nil {
		return appErrorf(err, "setResponseWriter() error: %v", err)
	}
	controller.LogCommentInterceptor.LogInfo(ctx, "GetComments() user_id: %v", sub)
	return nil
}

func (controller *CommentController) getCommentKeys(r *http.Request) ([]*datastore.Key, error) {
	movieID, err := strconv.Atoi(controller.MuxCommentInterceptor.Get(r, "movieID"))
	if err != nil {
		err = errors2.Wrap(err, "strconv.Atoi()")
		return nil, err
	}
	var commentKeys []*datastore.Key
	q := datastore.NewQuery("Comment").KeysOnly().Filter("deleted =", false).Filter("movie_id =", movieID).Order("-created_at").Limit(10)
	it := controller.DatastoreCommentInterceptor.Run(r, q)
	for {
		k, err := controller.DatastoreCommentInterceptor.Next(it)
		if err == datastore.Done {
			break
		}
		if err != nil {
			err = errors2.Wrap(err, "controller.DatastoreCommentInterceptor.Next()")
			return nil, err
		}
		commentKeys = append(commentKeys, k)
	}
	return commentKeys, nil
}

func (controller *CommentController) getMultiComments(r *http.Request, comments []*entity.Comment) ([]*entity.Comment, error) {
	err := controller.DatastoreCommentInterceptor.GetMulti(r, comments)
	mErr, _ := err.(appengine.MultiError)
	for _, e := range mErr {
		if e == nil {
			// entityが存在する
			continue
		}
		if e == datastore.ErrNoSuchEntity {
			// entityが存在しないけど正常系とみなすのでスルーする
			continue
		}
		// datastore.ErrNoSuchEntity以外のエラー
		return nil, err
	}
	return comments, nil
}

func (controller *CommentController) getCommentUserKeys(comments []*domain.Comment) []*datastore.Key {
	var userKeys []*datastore.Key
	for i := range comments {
		userKeys = append(userKeys, comments[i].UserKey)
	}
	return userKeys
}

func (controller *CommentController) getMultiUsers(r *http.Request, users []*entity.User) ([]*entity.User, error) {
	err := controller.DatastoreCommentInterceptor.GetMulti(r, users)
	mErr, _ := err.(appengine.MultiError)
	for _, e := range mErr {
		if e == nil {
			// entityが存在する
			continue
		}
		if e == datastore.ErrNoSuchEntity {
			// entityが存在しないけど正常系とみなすのでスルーする
			continue
		}
		// datastore.ErrNoSuchEntity以外のエラー
		return nil, err
	}
	return users, nil
}

