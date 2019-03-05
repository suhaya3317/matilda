package firebase_interface

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FirebaseCommentRepository struct {
	FirebaseHandler
}

func (repo *FirebaseCommentRepository) FindPublicKey(client *http.Client) (*http.Response, error) {
	return repo.GetPublicKey(client)
}

func (repo *FirebaseCommentRepository) ParseToken(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return repo.ParseJWT(idToken, keys)
}

func (repo *FirebaseCommentRepository) FindSub(parsedToken *jwt.Token) (string, bool) {
	return repo.GetSub(parsedToken)
}
