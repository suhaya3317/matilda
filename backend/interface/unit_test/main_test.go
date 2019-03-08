package unit_test

import (
	"matilda/backend/infrastructure"
	"matilda/backend/interface/controllers"
	"matilda/backend/interface/logging"
	"matilda/backend/interface/movie"
	"matilda/backend/usecase"
	"os"
	"testing"
)

var TargetMovie *controllers.MovieController
var MovieAPiHandler movie.MovieAPIHandler
var LogHandler logging.LogHandler

func TestMain(m *testing.M) {
	MovieAPiHandler = infrastructure.NewMovieAPIHandler()
	LogHandler = infrastructure.NewLogHandler()
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
