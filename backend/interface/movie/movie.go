package movie

import (
	"context"
	"net/http"
)

type MovieAPIHandler interface {
	GetPopularMovies(context.Context, string) (*http.Response, error)
}