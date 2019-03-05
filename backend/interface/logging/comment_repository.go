package logging

import "context"

type LogCommentRepository struct {
	LogHandler
}

func (repo *LogCommentRepository) Output(ctx context.Context, format string, args interface{}) {
	repo.LogInfo(ctx, format, args)
}
