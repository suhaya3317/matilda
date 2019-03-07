package database

import (
	"matilda/backend/domain/entity"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserRepository struct {
	DatastoreHandler
}

func (repo *DatastoreUserRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return repo.Put(r, src)
}

func (repo *DatastoreUserRepository) FindMulti(r *http.Request, src []*entity.User) error {
	return repo.GetMulti(r, src)
}
