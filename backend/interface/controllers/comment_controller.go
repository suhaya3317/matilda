package controllers

import (
	"context"
	"errors"
	"matilda/backend/domain/entity"
	"matilda/backend/interface/database"
	"matilda/backend/interface/firebase_interface"
	"matilda/backend/interface/gorilla_mux"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
	"net/http"
	"strconv"
	"time"

	errors2 "github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

type CommentController struct {
	MuxCommentInterceptor       usecase.MuxCommentInterceptor
	FirebaseCommentInterceptor  usecase.FirebaseCommentInterceptor
	DatastoreCommentInterceptor usecase.DatastoreCommentInterceptor
	LogCommentInterceptor       usecase.LogCommentInterceptor
}

func NewCommentController(gorillaMuxHandler gorilla_mux.GorillaMuxHandler, firebaseHandler firebase_interface.FirebaseHandler, datastoreHandler database.DatastoreHandler, logHandler logging.LogHandler) *CommentController {
	return &CommentController{
		MuxCommentInterceptor: usecase.MuxCommentInterceptor{
			MuxCommentRepository: &gorilla_mux.MuxCommentRepository{
				GorillaMuxHandler: gorillaMuxHandler,
			},
		},
		FirebaseCommentInterceptor: usecase.FirebaseCommentInterceptor{
			FirebaseCommentRepository: &firebase_interface.FirebaseCommentRepository{
				FirebaseHandler: firebaseHandler,
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

	sub, err := controller.getUserID(r, ctx)
	if err != nil {
		return appErrorf(err, "controller.getUserID() error: %v", err)
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

func (controller *CommentController) getUserID(r *http.Request, ctx context.Context) (string, error) {
	idToken := getIDToken(r)

	client := urlfetch.Client(ctx)
	resp, err := controller.FirebaseCommentInterceptor.GetPublicKey(client)
	if err != nil {
		err = errors2.Wrap(err, "controller.FirebaseCommentInterceptor.GetPublicKey()")
		return "", err
	}

	keys, err := decodePublicKeys(resp)
	if err != nil {
		err = errors2.Wrap(err, "decodePublicKeys()")
		return "", err
	}

	parsedToken, err := controller.FirebaseCommentInterceptor.ParseJWT(idToken, keys)
	if err != nil {
		err = errors2.Wrap(err, "controller.FirebaseCommentInterceptor.ParseJWT()")
		return "", err
	}

	sub, ok := controller.FirebaseCommentInterceptor.GetSub(parsedToken)
	if ok == false {
		err := errors.New("controller.FirebaseCommentInterceptor.GetSub(): could not get sub")
		return "", err
	}
	return sub, nil
}
