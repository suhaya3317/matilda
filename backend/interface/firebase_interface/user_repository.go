package firebase_interface

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type FirebaseUserRepository struct {
	FirebaseHandler
}

func (repo *FirebaseUserRepository) FindPublicKey(client *http.Client) (*http.Response, error) {
	return repo.GetPublicKey(client)
}

func (repo *FirebaseUserRepository) ParseToken(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return repo.ParseJWT(idToken, keys)
}

func (repo *FirebaseUserRepository) FindSub(parsedToken *jwt.Token) (string, bool) {
	return repo.GetSub(parsedToken)
}
