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
