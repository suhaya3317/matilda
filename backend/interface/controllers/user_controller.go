package controllers

import (
	"matilda/backend/interface/database"
	"matilda/backend/interface/logging"
	"matilda/backend/usecase"
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
