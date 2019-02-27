package usecase

import (
	"context"
	"net/http"
)

type MovieMuxRepository interface {
	Find(*http.Request, string) string
}

type MovieAPIRepository interface {
	FindAll(context.Context, string) (*http.Response, error)
	Find(context.Context, string) (*http.Response, error)
	FindInfo(context.Context, string) (*http.Response, error)
}

type LogMovieRepository interface {
	Output(context.Context, string, interface{})
}
