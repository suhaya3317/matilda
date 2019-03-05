package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FirebaseInternalRepository interface {
	FindPublicKey(*http.Client) (*http.Response, error)
	ParseToken(string, map[string]*json.RawMessage) (*jwt.Token, error)
	FindSub(*jwt.Token) (string, bool)
}
