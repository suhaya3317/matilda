package infrastructure

import (
	"context"
	"matilda/backend/interface/logging"

	"google.golang.org/appengine/log"
)

type LogHandler struct {
}

func NewLogHandler() logging.LogHandler {
	logHandler := new(LogHandler)
	return logHandler
}

func (handler *LogHandler) LogInfo(ctx context.Context, format string, args interface{}) {
	log.Infof(ctx, format, args)
}
