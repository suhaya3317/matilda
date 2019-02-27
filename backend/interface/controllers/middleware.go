package controllers

import (
	"fmt"
	"matilda/backend/interface/firebase_interface"
	"matilda/backend/usecase"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type FirebaseController struct {
	MiddlewareInterceptor usecase.MiddlewareInterceptor
}

func NewFirebaseController(firebaseHandler firebase_interface.FirebaseHandler) *FirebaseController {
	return &FirebaseController{
		MiddlewareInterceptor: usecase.MiddlewareInterceptor{
			MiddlewareRepository: &firebase_interface.MiddlewareRepository{
				FirebaseHandler: firebaseHandler,
			},
		},
	}
}

type AppHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if e := ah(w, r); e != nil { // e is *appError, not os.Error.
		log.Errorf(ctx, "Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}

func (controller *FirebaseController) AuthMiddleware(ah AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		statusCode, err := controller.MiddlewareInterceptor.Auth(ctx, idToken)
		if err != nil {
			log.Errorf(ctx, "AuthMiddleware error: %v", err)
			setResponseWriter(w, statusCode, err)
			return
		}
		ah.ServeHTTP(w, r)
	}
}
