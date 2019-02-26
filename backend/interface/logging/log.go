package logging

import "context"

type LogHandler interface {
	LogInfo(context.Context, string, interface{})
}
