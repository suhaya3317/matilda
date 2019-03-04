package controllers

import (
	"matilda/backend/usecase"
	"os"
	"testing"
)

var TargetUser UserController
var TargetMovie MovieController

func TestMain(m *testing.M) {
	TargetUser = UserController{
		DatastoreUserInterceptor: usecase.DatastoreUserInterceptor{
			DatastoreUserRepository: &MockDatastoreUserRepository{},
		},
		LogUserInterceptor: usecase.LogUserInterceptor{
			LogUserRepository: &MockLogUserRepository{},
		},
		FirebaseUserInterceptor: usecase.FirebaseUserInterceptor{
			FirebaseUserRepository: &MockFirebaseUserRepository{},
		},
	}

	TargetMovie = MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &MockMovieMuxRepository{},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &MockMovieAPIRepository{},
		},
		LogInterceptor: usecase.LogMovieInterceptor{
			LogMovieRepository: &MockLogRepository{},
		},
	}
	code := m.Run()
	os.Exit(code)
}
