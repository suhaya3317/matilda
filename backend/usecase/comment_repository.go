package usecase

import (
	"context"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type MuxCommentRepository interface {
	Find(*http.Request, string) string
}

type DatastoreCommentRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
	FindKey(*http.Request, interface{}) *datastore.Key
}

type LogCommentRepository interface {
	Output(context.Context, string, interface{})
}
