package controllers

import (
	"matilda/backend/usecase"
	"os"
	"testing"
)

var TargetUser UserController
var TargetMovie MovieController
var TargetComment CommentController

type publicKey struct {
	First  string `json:"3bbd28edc3d10b929f575a2ca68549fca6d88993"`
	Second string `json:"456299d7a268145bd9bb206f8889dac028468fe4"`
}

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

	TargetComment = CommentController{
		MuxCommentInterceptor: usecase.MuxCommentInterceptor{
			MuxCommentRepository: &MockMuxCommentRepository{},
		},
		FirebaseCommentInterceptor: usecase.FirebaseCommentInterceptor{
			FirebaseCommentRepository: &MockFirebaseCommentRepository{},
		},
		DatastoreCommentInterceptor: usecase.DatastoreCommentInterceptor{
			DatastoreCommentRepository: &MockDatastoreCommentRepository{},
		},
		LogCommentInterceptor: usecase.LogCommentInterceptor{
			LogCommentRepository: &MockLogCommentRepository{},
		},
	}
	code := m.Run()
	os.Exit(code)
}
