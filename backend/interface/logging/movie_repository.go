package logging

import "context"

type LogRepository struct {
	LogHandler
}

func (repo *LogRepository) Output(ctx context.Context, format string, args interface{}) {
	repo.LogInfo(ctx, format, args)
}
