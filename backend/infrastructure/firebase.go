package infrastructure

import (
	"context"
	"matilda/backend/interface/firebase_interface"
	"net/http"
	"os"

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

func (handler *FirebaseHandler) GetCertIDTokenURL(client *http.Client) (*http.Response, error) {
	return client.Get(clientCertURL)
}
