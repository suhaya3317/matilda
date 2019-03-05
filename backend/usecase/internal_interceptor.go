package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type FirebaseInternalInterceptor struct {
	FirebaseInternalRepository FirebaseInternalRepository
}

func (interceptor *FirebaseInternalInterceptor) GetPublicKey(client *http.Client) (*http.Response, error) {
	return interceptor.FirebaseInternalRepository.FindPublicKey(client)
}

func (interceptor *FirebaseInternalInterceptor) ParseJWT(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return interceptor.FirebaseInternalRepository.ParseToken(idToken, keys)
}

func (interceptor *FirebaseInternalInterceptor) GetSub(parsedToken *jwt.Token) (string, bool) {
	return interceptor.FirebaseInternalRepository.FindSub(parsedToken)
}
