package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/appengine/datastore"
)

type MuxCommentRepository interface {
	Find(*http.Request, string) string
}

type DatastoreCommentRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
	FindKey(*http.Request, interface{}) *datastore.Key
}

type LogCommentRepository interface {
	Output(context.Context, string, interface{})
}

type FirebaseCommentRepository interface {
	FindPublicKey(*http.Client) (*http.Response, error)
	ParseToken(string, map[string]*json.RawMessage) (*jwt.Token, error)
	FindSub(*jwt.Token) (string, bool)
}
