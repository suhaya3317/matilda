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

func (interceptor *MovieAPIInterceptor) GetPopularMovies(client *http.Client, page string) (*http.Response, error) {
	return interceptor.MovieAPIRepository.FindAll(client, page)
}

func (interceptor *MovieAPIInterceptor) GetMovie(client *http.Client, id string) (*http.Response, error) {
	return interceptor.MovieAPIRepository.Find(client, id)
}

func (interceptor *MovieAPIInterceptor) GetMovieInformation(client *http.Client, id string) (*http.Response, error) {
	return interceptor.MovieAPIRepository.FindInfo(client, id)
}

type LogMovieInterceptor struct {
	LogMovieRepository LogMovieRepository
}

func (interceptor *LogMovieInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogMovieRepository.Output(ctx, format, args)
}
