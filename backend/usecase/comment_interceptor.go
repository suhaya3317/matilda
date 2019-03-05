package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

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

type LogCommentInterceptor struct {
	LogCommentRepository LogCommentRepository
}

func (interceptor *LogCommentInterceptor) LogInfo(ctx context.Context, format string, args interface{}) {
	interceptor.LogCommentRepository.Output(ctx, format, args)
}

type FirebaseCommentInterceptor struct {
	FirebaseCommentRepository FirebaseCommentRepository
}

func (interceptor *FirebaseCommentInterceptor) GetPublicKey(client *http.Client) (*http.Response, error) {
	return interceptor.FirebaseCommentRepository.FindPublicKey(client)
}

func (interceptor *FirebaseCommentInterceptor) ParseJWT(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return interceptor.FirebaseCommentRepository.ParseToken(idToken, keys)
}

func (interceptor *FirebaseCommentInterceptor) GetSub(parsedToken *jwt.Token) (string, bool) {
	return interceptor.FirebaseCommentRepository.FindSub(parsedToken)
}
