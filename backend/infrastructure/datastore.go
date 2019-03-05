package infrastructure

import (
	"matilda/backend/interface/database"
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

type DatastoreHandler struct {
	Conn func(r *http.Request) *goon.Goon
}

func NewDatastoreHandler() database.DatastoreHandler {
	datastoreHandler := new(DatastoreHandler)
	datastoreHandler.Conn =
		func(r *http.Request) *goon.Goon {
			g := goon.NewGoon(r)
			return g
		}
	return datastoreHandler
}

func (handler *DatastoreHandler) Put(r *http.Request, src interface{}) (*datastore.Key, error) {
	return handler.Conn(r).Put(src)
}

func (handler *DatastoreHandler) GetKey(r *http.Request, src interface{}) *datastore.Key {
	return handler.Conn(r).Key(src)
}

func (handler *DatastoreHandler) GetMulti(r *http.Request, src interface{}) error {
	return handler.Conn(r).GetMulti(src)
}

func (handler *DatastoreHandler) Run(r *http.Request, query *datastore.Query) *goon.Iterator {
	return handler.Conn(r).Run(query)
}

func (handler *DatastoreHandler) Next(it *goon.Iterator) (*datastore.Key, error) {
	return it.Next(nil)
}
