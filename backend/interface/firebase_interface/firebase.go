package firebase_interface

import "context"

type FirebaseHandler interface {
	Auth(context.Context, string) (int, error)
}
