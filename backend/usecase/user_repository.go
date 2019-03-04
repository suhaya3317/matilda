package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/appengine/datastore"
)

type DatastoreUserRepository interface {
	Store(*http.Request, interface{}) (*datastore.Key, error)
}

type LogUserRepository interface {
	Output(context.Context, string, interface{})
}

type FirebaseUserRepository interface {
	FindPublicKey(*http.Client) (*http.Response, error)
	ParseToken(string, map[string]*json.RawMessage) (*jwt.Token, error)
	FindSub(*jwt.Token) (string, bool)
}
