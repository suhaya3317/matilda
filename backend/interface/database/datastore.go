package database

import (
	"net/http"

	"github.com/mjibson/goon"

	"google.golang.org/appengine/datastore"
)

type DatastoreHandler interface {
	Put(*http.Request, interface{}) (*datastore.Key, error)
	GetKey(*http.Request, interface{}) *datastore.Key
	GetMulti(*http.Request, interface{}) error
	Run(*http.Request, *datastore.Query) *goon.Iterator
	Next(*goon.Iterator) (*datastore.Key, error)
}
