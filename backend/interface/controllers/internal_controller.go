package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"matilda/backend/interface/firebase_interface"
	"matilda/backend/usecase"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	errors2 "github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

var Common *InternalController

type InternalController struct {
	FirebaseInternalInterceptor usecase.FirebaseInternalInterceptor
}

func NewInternalController(firebaseHandler firebase_interface.FirebaseHandler) *InternalController {
	return &InternalController{
		FirebaseInternalInterceptor: usecase.FirebaseInternalInterceptor{
			FirebaseInternalRepository: &firebase_interface.FirebaseInternalRepository{
				FirebaseHandler: firebaseHandler,
			},
		},
	}
}

func getUserID(r *http.Request, ctx context.Context) (string, error) {
	idToken := getIDToken(r)

	client := urlfetch.Client(ctx)
	resp, err := Common.FirebaseInternalInterceptor.GetPublicKey(client)
	if err != nil {
		err = errors2.Wrap(err, "controller.FirebaseInternalInterceptor.GetPublicKey()")
		return "", err
	}

	keys, err := decodePublicKeys(resp)
	if err != nil {
		err = errors2.Wrap(err, "decodePublicKeys()")
		return "", err
	}

	parsedToken, err := Common.FirebaseInternalInterceptor.ParseJWT(idToken, keys)
	if err != nil {
		err = errors2.Wrap(err, "controller.FirebaseInternalInterceptor.ParseJWT()")
		return "", err
	}

	sub, ok := Common.FirebaseInternalInterceptor.GetSub(parsedToken)
	if ok == false {
		err := errors.New("controller.FirebaseInternalInterceptor.GetSub(): could not get sub")
		return "", err
	}
	return sub, nil
}

func getIDToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	return idToken
}

func decodePublicKeys(resp *http.Response) (map[string]*json.RawMessage, error) {
	var objmap map[string]*json.RawMessage
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&objmap)
	err = errors.Wrap(err, "decoder.Decode()")

	return objmap, err
}

func mappingJsonToStruct(r *http.Request, src interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, &src); err != nil {
		return err
	}
	return nil
}

func setResponseWriter(w http.ResponseWriter, statusCode int, src interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(src)
}
