package logging

import "context"

type LogMovieRepository struct {
	LogHandler
}

func (repo *LogMovieRepository) Output(ctx context.Context, format string, args interface{}) {
	repo.LogInfo(ctx, format, args)
}
