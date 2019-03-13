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

	entityComments, err := controller.setCommentIdToEntityComments(r)
	if err != nil {
		return appErrorf(err, "controller.setCommentIdToEntityComments() error: %v", err)
	}

	err = controller.DatastoreCommentInterceptor.GetMulti(r, entityComments)
	err = checkGetMultiErr(err)
	if err != nil {
		return appErrorf(err, "checkGetMultiErr() error: %v", err)
	}

	comments := controller.setEntityCommentsToComments(entityComments, sub)

	entityUsers := controller.setUserIdtoEntityUsers(comments)

	err = UserHandler.DatastoreUserInterceptor.GetMulti(r, entityUsers)
	err = checkGetMultiErr(err)
	if err != nil {
		return appErrorf(err, "checkGetMultiErr() error: %v", err)
	}

	controller.setEntityUsersToComments(comments, entityUsers)

	err = setResponseWriter(w, 200, &comments)
	if err != nil {
		return appErrorf(err, "setResponseWriter() error: %v", err)
	}
	controller.LogCommentInterceptor.LogInfo(ctx, "GetComments() user_id: %v", sub)
	return nil
}

func (controller *CommentController) setCommentIdToEntityComments(r *http.Request) ([]*entity.Comment, error) {
	commentKeys, err := controller.getCommentKeys(r)
	if err != nil {
		err = errors2.Wrap(err, "controller.getCommentKeys()")
		return nil, err
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
	return entityComments, nil
}

func (controller *CommentController) getCommentKeys(r *http.Request) ([]*datastore.Key, error) {
	movieID, err := strconv.Atoi(controller.MuxCommentInterceptor.Get(r, "movieID"))
	if err != nil {
		err = errors2.Wrap(err, "strconv.Atoi(): movieID")
		return nil, err
	}
	var commentKeys []*datastore.Key

	page, err := strconv.Atoi(controller.MuxCommentInterceptor.Get(r, "page"))
	if err != nil {
		err = errors2.Wrap(err, "strconv.Atoi(): page")
		return nil, err
	}

	// TODO: offsetではなくcursorで解決したい
	q := datastore.NewQuery("Comment").KeysOnly().Filter("deleted =", false).Filter("movie_id =", movieID).Order("-created_at").Limit(10).Offset(page*10 - 10)

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

func (controller *CommentController) setEntityCommentsToComments(entityComments []*entity.Comment, sub string) []*domain.Comment {
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
	return comments
}

func (controller *CommentController) setUserIdtoEntityUsers(comments []*domain.Comment) []*entity.User {
	userKeys := controller.getCommentUserKeys(comments)

	var userIDs []string
	for i := range userKeys {
		userIDs = append(userIDs, userKeys[i].StringID())
	}

	var entityUsers []*entity.User
	for i := range userIDs {
		userID := userIDs[i]
		entityUsers = append(entityUsers, &entity.User{
			UserID: userID,
		})
	}
	return entityUsers
}

func (controller *CommentController) getCommentUserKeys(comments []*domain.Comment) []*datastore.Key {
	var userKeys []*datastore.Key
	for i := range comments {
		userKeys = append(userKeys, comments[i].UserKey)
	}
	return userKeys
}

func (controller *CommentController) setEntityUsersToComments(comments []*domain.Comment, entityUsers []*entity.User) {
	for i := range entityUsers {
		comments[i].Name = entityUsers[i].Name
		comments[i].IconPath = entityUsers[i].IconPath
	}
}
