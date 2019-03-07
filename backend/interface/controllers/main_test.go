package controllers

import (
	"encoding/json"
	"errors"
	"matilda/backend/infrastructure"
	"matilda/backend/interface/database"
	"matilda/backend/interface/logging"
	"matilda/backend/interface/movie"
	"matilda/backend/usecase"
	"os"
	"reflect"
	"testing"
)

var TargetMovie *MovieController
var TargetComment *CommentController

type publicKey struct {
	First  string `json:"3bbd28edc3d10b929f575a2ca68549fca6d88993"`
	Second string `json:"456299d7a268145bd9bb206f8889dac028468fe4"`
}

func TestMain(m *testing.M) {
	movieAPIHandler := infrastructure.NewMovieAPIHandler()
	datastoreHandler := infrastructure.NewDatastoreHandler()
	logHandler := infrastructure.NewLogHandler()

	Common = &InternalController{
		FirebaseInternalInterceptor: usecase.FirebaseInternalInterceptor{
			FirebaseInternalRepository: &MockFirebaseInternalRepository{},
		},
	}
	UserHandler = &UserController{
		DatastoreUserInterceptor: usecase.DatastoreUserInterceptor{
			DatastoreUserRepository: &MockDatastoreUserRepository{},
		},
		LogUserInterceptor: usecase.LogUserInterceptor{
			LogUserRepository: &MockLogUserRepository{},
		},
	}

	TargetMovie = &MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &MockMovieMuxRepository{},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &movie.MovieAPIRepository{
				MovieAPIHandler: movieAPIHandler,
			},
		},
		LogInterceptor: usecase.LogMovieInterceptor{
			LogMovieRepository: &logging.LogMovieRepository{
				LogHandler: logHandler,
			},
		},
	}

	TargetComment = &CommentController{
		MuxCommentInterceptor: usecase.MuxCommentInterceptor{
			MuxCommentRepository: &MockMuxCommentRepository{},
		},
		DatastoreCommentInterceptor: usecase.DatastoreCommentInterceptor{
			DatastoreCommentRepository: &database.DatastoreCommentRepository{
				DatastoreHandler: datastoreHandler,
			},
		},
		LogCommentInterceptor: usecase.LogCommentInterceptor{
			LogCommentRepository: &MockLogCommentRepository{},
		},
	}
	code := m.Run()
	os.Exit(code)
}

func IsEqualJSON(a, b string) error {
	err, ok := DeepEqualJSON(a, b)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	err = errors.New("not Equal")
	return err
}

func DeepEqualJSON(j1, j2 string) (error, bool) {
	var err error

	var d1 interface{}
	err = json.Unmarshal([]byte(j1), &d1)

	if err != nil {
		return err, false
	}

	var d2 interface{}
	err = json.Unmarshal([]byte(j2), &d2)

	if err != nil {
		return err, false
	}

	if reflect.DeepEqual(d1, d2) {
		return nil, true
	} else {
		return nil, false
	}
}
