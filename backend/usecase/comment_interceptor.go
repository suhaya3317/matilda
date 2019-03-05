package usecase

import (
	"context"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreCommentInterceptor struct {
	DatastoreCommentRepository DatastoreCommentRepository
}

func (interceptor *DatastoreCommentInterceptor) Put(r *http.Request, src interface{}) (*datastore.Key, error) {
	return interceptor.DatastoreCommentRepository.Store(r, src)
}

type LogCommentInterceptor struct {
	LogCommentRepository LogCommentRepository
}

func (interceptor *LogCommentInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogCommentRepository.Output(ctx, format, args)
}
