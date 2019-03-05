package database

import (
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreHandler interface {
	Put(*http.Request, interface{}) (*datastore.Key, error)
	GetKey(*http.Request, interface{}) *datastore.Key
}
