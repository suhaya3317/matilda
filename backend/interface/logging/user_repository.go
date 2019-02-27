package logging

import "context"

type LogUserRepository struct {
	LogHandler
}

func (repo *LogUserRepository) Output(ctx context.Context, format string, args interface{}) {
	repo.LogInfo(ctx, format, args)
}
