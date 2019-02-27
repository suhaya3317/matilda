package controllers

import (
	"matilda/backend/usecase"
	"os"
	"testing"
)

var TargetMovie MovieController

func TestMain(m *testing.M) {
	TargetMovie = MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &MockMovieMuxRepository{},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &MockMovieAPIRepository{},
		},
		LogInterceptor: usecase.LogInterceptor{
			LogRepository: &MockLogRepository{},
		},
	}
	code := m.Run()
	os.Exit(code)
}
