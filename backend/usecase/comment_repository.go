package usecase

import (
	"context"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreCommentRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
}

type LogCommentRepository interface {
	Output(context.Context, string, interface{})
}
