package usecase

import (
	"context"
	"net/http"
)

type MovieMuxRepository interface {
	Find(*http.Request, string) string
}

type MovieAPIRepository interface {
	FindAll(*http.Client, string) (*http.Response, error)
	Find(*http.Client, string) (*http.Response, error)
	FindInfo(*http.Client, string) (*http.Response, error)
}

type LogMovieRepository interface {
	Output(context.Context, string, interface{})
}
