package gorilla_mux

import "net/http"

type GorillaMuxHandler interface {
	Get(*http.Request, string) string
}
