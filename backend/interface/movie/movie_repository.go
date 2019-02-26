package movie

import (
	"context"
	"net/http"
)

type MovieAPIRepository struct {
	MovieAPIHandler
}

func (repo *MovieAPIRepository) FindAll(ctx context.Context, page string) (*http.Response, error) {
	return repo.GetPopularMovies(ctx, page)
}

func (repo *MovieAPIRepository) Find(ctx context.Context, id string) (*http.Response, error) {
	return repo.GetMovie(ctx, id)
}

func (repo *MovieAPIRepository) FindInfo(ctx context.Context, id string) (*http.Response, error) {
	return repo.GetMovieInformation(ctx, id)
}
