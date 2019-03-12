package unit_test

import (
	"matilda/backend/infrastructure"
	"matilda/backend/interface/controllers"
	"matilda/backend/interface/database"
	"matilda/backend/interface/logging"
	"matilda/backend/interface/movie"
	"matilda/backend/usecase"
	"os"
	"testing"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/appengine/aetest"
)

var AuthToken string
var TargetUser *controllers.UserController
var TargetComment *controllers.CommentController
var TargetMovie *controllers.MovieController
var DatastoreHandler database.DatastoreHandler
var MovieAPiHandler movie.MovieAPIHandler
var LogHandler logging.LogHandler

func TestMain(m *testing.M) {
	AuthToken = authFirebase()

	firebaseHandler := infrastructure.NewFirebaseHandler()
	controllers.Common = controllers.NewInternalController(firebaseHandler)
	controllers.Common.FirebaseInternalInterceptor.FirebaseInternalRepository = &MockFirebaseInternalRepository{}
	DatastoreHandler = infrastructure.NewDatastoreHandler()
	MovieAPiHandler = infrastructure.NewMovieAPIHandler()
	LogHandler = infrastructure.NewLogHandler()

	controllers.UserHandler = &controllers.UserController{
		DatastoreUserInterceptor: usecase.DatastoreUserInterceptor{
			DatastoreUserRepository: &database.DatastoreUserRepository{
				DatastoreHandler: DatastoreHandler,
			},
		},
		LogUserInterceptor: usecase.LogUserInterceptor{
			LogUserRepository: &logging.LogUserRepository{
				LogHandler: LogHandler,
			},
		},
	}

	TargetUser = controllers.UserHandler

	TargetComment = &controllers.CommentController{
		MuxCommentInterceptor: usecase.MuxCommentInterceptor{
			MuxCommentRepository: &MockMuxCommentRepository{},
		},
		DatastoreCommentInterceptor: usecase.DatastoreCommentInterceptor{
			DatastoreCommentRepository: &database.DatastoreCommentRepository{
				DatastoreHandler: DatastoreHandler,
			},
		},
		LogCommentInterceptor: usecase.LogCommentInterceptor{
			LogCommentRepository: &logging.LogCommentRepository{
				LogHandler: LogHandler,
			},
		},
	}

	TargetMovie = &controllers.MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &MockMovieMuxRepository{},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &movie.MovieAPIRepository{
				MovieAPIHandler: MovieAPiHandler,
			},
		},
		LogInterceptor: usecase.LogMovieInterceptor{
			LogMovieRepository: &logging.LogMovieRepository{
				LogHandler: LogHandler,
			},
		},
	}
	code := m.Run()
	os.Exit(code)
}

func authFirebase() string {
	ctx, _, err := aetest.NewContext()
	if err != nil {
		panic(err)
	}
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	token, err := client.CustomToken(ctx, "some-uid")
	if err != nil {
		panic(err)
	}
	return token
}

func Equal(a, b []string) bool {
	var check bool
	if len(a) != len(b) {
		return false
	}
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				check = true
				break
			}
			check = false
		}
		if check == false {
			break
		}
	}
	return check
}
