package firebase_interface

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FirebaseInternalRepository struct {
	FirebaseHandler
}

func (repo *FirebaseInternalRepository) FindPublicKey(client *http.Client) (*http.Response, error) {
	return repo.GetPublicKey(client)
}

func (repo *FirebaseInternalRepository) ParseToken(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return repo.ParseJWT(idToken, keys)
}

func (repo *FirebaseInternalRepository) FindSub(parsedToken *jwt.Token) (string, bool) {
	return repo.GetSub(parsedToken)
}
