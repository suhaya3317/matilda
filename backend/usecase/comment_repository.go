package usecase

import (
	"context"
	"net/http"

	"github.com/mjibson/goon"

	"google.golang.org/appengine/datastore"
)

type MuxCommentRepository interface {
	Find(*http.Request, string) string
}

type DatastoreCommentRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
	FindKey(*http.Request, interface{}) *datastore.Key
	FindMulti(*http.Request, interface{}) error
	RunQuery(*http.Request, *datastore.Query) *goon.Iterator
	NextQuery(*goon.Iterator) (*datastore.Key, error)
}

type LogCommentRepository interface {
	Output(context.Context, string, interface{})
}
