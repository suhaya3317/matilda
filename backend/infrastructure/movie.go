package infrastructure

import (
	"context"
	"matilda/backend/interface/movie"
	"net/http"
	"os"

	"google.golang.org/appengine/urlfetch"
)

type MovieAPIHandler struct {
	APIKey string
	Client func(ctx context.Context) *http.Client
}

func NewMovieAPIHandler() movie.MovieAPIHandler {
	movieAPIHandler := new(MovieAPIHandler)
	movieAPIHandler.APIKey = os.Getenv("MOVIE_DB_API_KEY")
	movieAPIHandler.Client =
		func(ctx context.Context) *http.Client {
			return urlfetch.Client(ctx)
		}
	return movieAPIHandler
}

func (handler *MovieAPIHandler) GetPopularMovies(ctx context.Context, page string) (*http.Response, error) {
	res, err := handler.Client(ctx).Get("https://api.themoviedb.org/3/movie/popular?api_key=" + handler.APIKey + "&page=" + page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (handler *MovieAPIHandler) GetMovie(ctx context.Context, id string) (*http.Response, error) {
	res, err := handler.Client(ctx).Get("https://api.themoviedb.org/3/movie/" + id + "?api_key=" + handler.APIKey)
	if err != nil {
		return nil, err
	}
	return res, err
}
