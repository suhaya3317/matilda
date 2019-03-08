package infrastructure

import (
	"matilda/backend/interface/movie"
	"net/http"
	"os"
)

type MovieAPIHandler struct {
	APIKey string
}

func NewMovieAPIHandler() movie.MovieAPIHandler {
	movieAPIHandler := new(MovieAPIHandler)
	movieAPIHandler.APIKey = os.Getenv("MOVIE_DB_API_KEY")
	return movieAPIHandler
}

func (handler *MovieAPIHandler) GetPopularMovies(client *http.Client, page string) (*http.Response, error) {
	res, err := client.Get("https://api.themoviedb.org/3/movie/popular?api_key=" + handler.APIKey + "&page=" + page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (handler *MovieAPIHandler) GetMovie(client *http.Client, id string) (*http.Response, error) {
	res, err := client.Get("https://api.themoviedb.org/3/movie/" + id + "?api_key=" + handler.APIKey)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (handler *MovieAPIHandler) GetMovieInformation(client *http.Client, id string) (*http.Response, error) {
	res, err := client.Get("https://api.themoviedb.org/3/movie/" + id + "?api_key=" + handler.APIKey + "&append_to_response=credits")
	if err != nil {
		return nil, err
	}
	return res, nil
}
