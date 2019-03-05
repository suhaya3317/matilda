package controllers

import (
	"errors"
	"matilda/backend/domain/entity"
	"matilda/backend/interface/database"
	"matilda/backend/interface/gorilla_mux"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
	"net/http"
	"strconv"
	"time"

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

	sub, err := getUserID(r, ctx, Common)
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
