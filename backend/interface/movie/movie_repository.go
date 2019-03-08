package movie

import (
	"net/http"
)

type MovieAPIRepository struct {
	MovieAPIHandler
}

func (repo *MovieAPIRepository) FindAll(client *http.Client, page string) (*http.Response, error) {
	return repo.GetPopularMovies(client, page)
}

func (repo *MovieAPIRepository) Find(client *http.Client, id string) (*http.Response, error) {
	return repo.GetMovie(client, id)
}

func (repo *MovieAPIRepository) FindInfo(client *http.Client, id string) (*http.Response, error) {
	return repo.GetMovieInformation(client, id)
}
