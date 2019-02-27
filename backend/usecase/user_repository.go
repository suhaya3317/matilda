package usecase

import (
	"context"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
}

type LogUserRepository interface {
	Output(context.Context, string, interface{})
}
