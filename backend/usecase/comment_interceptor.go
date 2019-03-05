package usecase

import (
	"context"
	"net/http"

	"github.com/mjibson/goon"

	"google.golang.org/appengine/datastore"
)

type MuxCommentInterceptor struct {
	MuxCommentRepository MuxCommentRepository
}

func (interceptor *MuxCommentInterceptor) Get(r *http.Request, key string) string {
	return interceptor.MuxCommentRepository.Find(r, key)
}

type DatastoreCommentInterceptor struct {
	DatastoreCommentRepository DatastoreCommentRepository
}

func (interceptor *DatastoreCommentInterceptor) Put(r *http.Request, src interface{}) (*datastore.Key, error) {
	return interceptor.DatastoreCommentRepository.Store(r, src)
}

func (interceptor *DatastoreCommentInterceptor) GetKey(r *http.Request, src interface{}) *datastore.Key {
	return interceptor.DatastoreCommentRepository.FindKey(r, src)
}

func (interceptor *DatastoreCommentInterceptor) GetMulti(r *http.Request, src interface{}) error {
	return interceptor.DatastoreCommentRepository.FindMulti(r, src)
}

func (interceptor *DatastoreCommentInterceptor) Run(r *http.Request, query *datastore.Query) *goon.Iterator {
	return interceptor.DatastoreCommentRepository.RunQuery(r, query)
}

func (interceptor *DatastoreCommentInterceptor) Next(it *goon.Iterator) (*datastore.Key, error) {
	return interceptor.DatastoreCommentRepository.NextQuery(it)
}

type LogCommentInterceptor struct {
	LogCommentRepository LogCommentRepository
}

func (interceptor *LogCommentInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogCommentRepository.Output(ctx, format, args)
}
