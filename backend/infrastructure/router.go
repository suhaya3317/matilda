package infrastructure

import (
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func RegisterHandlers() {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods("GET").Path("/api/v1/example").
		HandlerFunc(Example)
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))
}

func Example(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "example")
}
