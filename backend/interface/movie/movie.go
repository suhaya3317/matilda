package movie

import (
	"net/http"
)

type MovieAPIHandler interface {
	GetPopularMovies(*http.Client, string) (*http.Response, error)
	GetMovie(*http.Client, string) (*http.Response, error)
	GetMovieInformation(*http.Client, string) (*http.Response, error)
}
