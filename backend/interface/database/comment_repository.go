package database

import (
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreCommentRepository struct {
	DatastoreHandler
}

func (repo *DatastoreCommentRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return repo.Put(r, src)
}
