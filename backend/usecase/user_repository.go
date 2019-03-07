package usecase

import (
	"context"
	"matilda/backend/domain/entity"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
	FindMulti(*http.Request, []*entity.User) error
}

type LogUserRepository interface {
	Output(context.Context, string, interface{})
}
