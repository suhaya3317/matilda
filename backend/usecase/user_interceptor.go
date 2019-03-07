package usecase

import (
	"context"
	"matilda/backend/domain/entity"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserInterceptor struct {
	DatastoreUserRepository DatastoreUserRepository
}

func (interceptor *DatastoreUserInterceptor) Put(r *http.Request, src interface{}) (*datastore.Key, error) {
	return interceptor.DatastoreUserRepository.Store(r, src)
}

func (interceptor *DatastoreUserInterceptor) GetMulti(r *http.Request, src []*entity.User) error {
	return interceptor.DatastoreUserRepository.FindMulti(r, src)
}

type LogUserInterceptor struct {
	LogUserRepository LogUserRepository
}

func (interceptor *LogUserInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogUserRepository.Output(ctx, format, args)
}
