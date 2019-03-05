package controllers

import (
	"errors"
	"matilda/backend/domain/entity"
	"matilda/backend/interface/database"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

type UserController struct {
	DatastoreUserInterceptor usecase.DatastoreUserInterceptor
	LogUserInterceptor       usecase.LogUserInterceptor
}

func NewUserController(datastoreHandler database.DatastoreHandler, logHandler logging.LogHandler) *UserController {
	return &UserController{
		DatastoreUserInterceptor: usecase.DatastoreUserInterceptor{
			DatastoreUserRepository: &database.DatastoreUserRepository{
				DatastoreHandler: datastoreHandler,
			},
		},
		LogUserInterceptor: usecase.LogUserInterceptor{
			LogUserRepository: &logging.LogUserRepository{
				LogHandler: logHandler,
			},
		},
	}
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)

	sub, err := getUserID(r, ctx)
	if err != nil {
		return appErrorf(err, "getUserID() error: %v", err)
	}
	if sub == "" {
		err := errors.New("could not get sub")
		return appErrorf(err, "controller.getUserID() error: %v", err)
	}

	var u entity.User
	u.UserID = sub
	u.Name = "test name"
	u.IconPath = "test path"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	_, err = controller.DatastoreUserInterceptor.Put(r, &u)
	if err != nil {
		return appErrorf(err, "controller.DatastoreUserInterceptor.Put() error: %v", err)
	}

	err = setResponseWriter(w, 202, err)
	if err != nil {
		return appErrorf(err, "setResponseWriter() error: %v", err)
	}

	controller.LogUserInterceptor.LogInfo(ctx, "CreateUser() user_id: %v", sub)
	return nil
}
