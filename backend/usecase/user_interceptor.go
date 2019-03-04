package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserInterceptor struct {
	DatastoreUserRepository DatastoreUserRepository
}

func (interceptor *DatastoreUserInterceptor) Put(r *http.Request, src interface{}) (*datastore.Key, error) {
	return interceptor.DatastoreUserRepository.Store(r, src)
}

type LogUserInterceptor struct {
	LogUserRepository LogUserRepository
}

func (interceptor *LogUserInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogUserRepository.Output(ctx, format, args)
}

type FirebaseUserInterceptor struct {
	FirebaseUserRepository FirebaseUserRepository
}

func (interceptor *FirebaseUserInterceptor) GetPublicKey(client *http.Client) (*http.Response, error) {
	return interceptor.FirebaseUserRepository.FindPublicKey(client)
}

func (interceptor *FirebaseUserInterceptor) ParseJWT(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return interceptor.FirebaseUserRepository.ParseToken(idToken, keys)
}

func (interceptor *FirebaseUserInterceptor) GetSub(parsedToken *jwt.Token) (string, bool) {
	return interceptor.FirebaseUserRepository.FindSub(parsedToken)
}
