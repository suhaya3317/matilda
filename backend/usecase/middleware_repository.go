package usecase

import "context"

type MiddlewareRepository interface {
	Check(context.Context, string) (int, error)
}
