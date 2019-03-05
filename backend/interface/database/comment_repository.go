package database

import (
	"net/http"

	"github.com/mjibson/goon"

	"google.golang.org/appengine/datastore"
)

type DatastoreCommentRepository struct {
	DatastoreHandler
}

func (repo *DatastoreCommentRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return repo.Put(r, src)
}

func (repo *DatastoreCommentRepository) FindKey(r *http.Request, src interface{}) *datastore.Key {
	return repo.GetKey(r, src)
}

func (repo *DatastoreCommentRepository) FindMulti(r *http.Request, src interface{}) error {
	return repo.GetMulti(r, src)
}

func (repo *DatastoreCommentRepository) RunQuery(r *http.Request, query *datastore.Query) *goon.Iterator {
	return repo.Run(r, query)
}

func (repo *DatastoreCommentRepository) NextQuery(it *goon.Iterator) (*datastore.Key, error) {
	return repo.Next(it)
}
