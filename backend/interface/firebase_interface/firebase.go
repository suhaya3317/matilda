package firebase_interface

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FirebaseHandler interface {
	Auth(context.Context, string) (int, error)
	GetPublicKey(*http.Client) (*http.Response, error)
	ParseJWT(string, map[string]*json.RawMessage) (*jwt.Token, error)
	GetSub(*jwt.Token) (string, bool)
}
