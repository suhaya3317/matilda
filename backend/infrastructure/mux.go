package infrastructure

import (
	"matilda/backend/interface/gorilla_mux"
	"net/http"

	"github.com/gorilla/mux"
)

type GorillaMuxHandler struct {
	Vars func(*http.Request) map[string]string
}

func NewGorillaMuxHandler() gorilla_mux.GorillaMuxHandler {
	gorillaMuxHandler := new(GorillaMuxHandler)
	gorillaMuxHandler.Vars =
		func(r *http.Request) map[string]string {
			return mux.Vars(r)
		}
	return gorillaMuxHandler
}

func (handler *GorillaMuxHandler) Get(r *http.Request, key string) string {
	return handler.Vars(r)[key]
}
