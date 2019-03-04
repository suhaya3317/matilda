package infrastructure

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"matilda/backend/interface/firebase_interface"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/pkg/errors"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/appengine/log"
)

const (
	clientCertURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
)

type FirebaseHandler struct {
	Opt option.ClientOption
	App func(context.Context, option.ClientOption) (*firebase.App, error)
}

func NewFirebaseHandler() firebase_interface.FirebaseHandler {
	firebaseHandler := new(FirebaseHandler)
	firebaseHandler.Opt = option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	firebaseHandler.App =
		func(ctx context.Context, opt option.ClientOption) (*firebase.App, error) {
			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				return nil, err
			}
			return app, nil
		}
	return firebaseHandler
}

func (handler *FirebaseHandler) Auth(ctx context.Context, idToken string) (int, error) {

	app, err := handler.App(ctx, handler.Opt)
	if err != nil {
		err = errors.Wrap(err, "handler.App")
		return 401, err
	}

	auth, err := app.Auth(ctx)
	err = errors.Wrap(err, "app.Auth")
	if err != nil {
		return 401, err
	}

	token, err := auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		err = errors.Wrap(err, "auth.VerifyIDToken")
		return 401, err
	}

	log.Infof(ctx, "Verified ID token: %v\n", token)
	return 200, nil
}

func (handler *FirebaseHandler) GetPublicKey(client *http.Client) (*http.Response, error) {
	return client.Get(clientCertURL)
}

func (handler *FirebaseHandler) ParseJWT(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid := token.Header["kid"]
		rsaPublicKey := convertKey(string(*keys[kid.(string)]))
		return rsaPublicKey, nil
	})
}

func (handler *FirebaseHandler) GetSub(parsedToken *jwt.Token) (string, bool) {
	tokenMap := parsedToken.Claims.(jwt.MapClaims)
	sub, ok := tokenMap["sub"].(string)
	return sub, ok
}

func convertKey(key string) interface{} {
	certPEM := key
	certPEM = strings.Replace(certPEM, "\\n", "\n", -1)
	certPEM = strings.Replace(certPEM, "\"", "", -1)
	block, _ := pem.Decode([]byte(certPEM))
	cert, _ := x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	return rsaPublicKey
}
