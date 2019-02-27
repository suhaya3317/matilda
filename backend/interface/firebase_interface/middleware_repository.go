package firebase_interface

import "context"

type MiddlewareRepository struct {
	FirebaseHandler
}

func (fire *MiddlewareRepository) Check(ctx context.Context, idToken string) (int, error) {
	return fire.Auth(ctx, idToken)
}
