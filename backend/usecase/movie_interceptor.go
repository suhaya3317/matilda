package usecase

import (
	"context"
	"net/http"
)

type MovieMuxInterceptor struct {
	MovieMuxRepository MovieMuxRepository
}

func (interceptor *MovieMuxInterceptor) Get(r *http.Request, key string) string {
	return interceptor.MovieMuxRepository.Find(r, key)
}

type MovieAPIInterceptor struct {
	MovieAPIRepository MovieAPIRepository
}

func (interceptor *MovieAPIInterceptor) GetPopularMovies(ctx context.Context, page string) (*http.Response, error) {
	return interceptor.MovieAPIRepository.FindAll(ctx, page)
}

func (interceptor *MovieAPIInterceptor) GetMovie(ctx context.Context, id string) (*http.Response, error) {
	return interceptor.MovieAPIRepository.Find(ctx, id)
}

type LogInterceptor struct {
	LogRepository LogRepository
}

func (interceptor *LogInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogRepository.Output(ctx, format, args)
}
