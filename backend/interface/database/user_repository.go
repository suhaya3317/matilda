package database

import (
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserRepository struct {
	DatastoreHandler
}

func (repo *DatastoreUserRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return repo.Put(r, src)
}
