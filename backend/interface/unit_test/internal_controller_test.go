package unit_test

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type MockFirebaseInternalRepository struct{}

func (mock *MockFirebaseInternalRepository) FindPublicKey(client *http.Client) (*http.Response, error) {
	clientCertURL := "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
	return client.Get(clientCertURL)
}

func (mock *MockFirebaseInternalRepository) ParseToken(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	token := &jwt.Token{}
	return token, nil
}

func (mock *MockFirebaseInternalRepository) FindSub(parsedToken *jwt.Token) (string, bool) {
	return "test-sub", true
}
