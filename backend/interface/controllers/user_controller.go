package controllers

import (
	"errors"
	"matilda/backend/domain/entity"
	"matilda/backend/interface/database"
	"matilda/backend/interface/firebase_interface"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
	"net/http"
	"time"

	"google.golang.org/appengine/urlfetch"

	"google.golang.org/appengine"
)

type UserController struct {
	DatastoreUserInterceptor usecase.DatastoreUserInterceptor
	LogUserInterceptor       usecase.LogUserInterceptor
	FirebaseUserInterceptor  usecase.FirebaseUserInterceptor
}

func NewUserController(datastoreHandler database.DatastoreHandler, logHandler logging.LogHandler, firebaseHandler firebase_interface.FirebaseHandler) *UserController {
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
		FirebaseUserInterceptor: usecase.FirebaseUserInterceptor{
			FirebaseUserRepository: &firebase_interface.FirebaseUserRepository{
				FirebaseHandler: firebaseHandler,
			},
		},
	}
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)

	idToken := getIDToken(r)

	client := urlfetch.Client(ctx)
	resp, err := controller.FirebaseUserInterceptor.GetPublicKey(client)
	if err != nil {
		return appErrorf(err, "controller.FirebaseUserInterceptor.GetPublicKey() error: %v", err)
	}

	keys, err := decodePublicKeys(resp)
	if err != nil {
		return appErrorf(err, "decodePublicKeys() error: %v", err)
	}

	parsedToken, err := controller.FirebaseUserInterceptor.ParseJWT(idToken, keys)
	if err != nil {
		return appErrorf(err, "controller.FirebaseUserInterceptor.ParseJWT() error: %v", err)
	}

	sub, ok := controller.FirebaseUserInterceptor.GetSub(parsedToken)
	if ok == false {
		err := errors.New("could not get sub")
		return appErrorf(err, "getSub() ok: %v", err)
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
