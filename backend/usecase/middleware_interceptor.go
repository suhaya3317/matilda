package usecase

import "context"

type MiddlewareInterceptor struct {
	MiddlewareRepository MiddlewareRepository
}

func (interceptor *MiddlewareInterceptor) Auth(ctx context.Context, idToken string) (int, error) {
	return interceptor.MiddlewareRepository.Check(ctx, idToken)
}
