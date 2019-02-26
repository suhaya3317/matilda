package infrastructure

import (
	"matilda/backend/interface/controllers"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func RegisterHandlers() {
	r := mux.NewRouter().StrictSlash(true)

	gorillaMuxHandler := NewGorillaMuxHandler()
	movieAPIHandler := NewMovieAPIHandler()
	logHandler := NewLogHandler()

	movieController := controllers.NewMovieController(gorillaMuxHandler, movieAPIHandler, logHandler)

	r.Methods("GET").Path("/api/v1/movies").Queries("page", "{page}").
		Handler(controllers.AppHandler(movieController.GetMovies))
	r.Methods("GET").Path("/api/v1/movies/{movieID}").
		Handler(controllers.AppHandler(movieController.GetMovie))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))
}
