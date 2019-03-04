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

	firebaseHandler := NewFirebaseHandler()
	gorillaMuxHandler := NewGorillaMuxHandler()
	datastoreHandler := NewDatastoreHandler()
	movieAPIHandler := NewMovieAPIHandler()
	logHandler := NewLogHandler()

	firebaseController := controllers.NewFirebaseController(firebaseHandler)
	userController := controllers.NewUserController(datastoreHandler, logHandler, firebaseHandler)
	movieController := controllers.NewMovieController(gorillaMuxHandler, movieAPIHandler, logHandler)

	r.Methods("PUT").Path("/api/v1/users").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(userController.CreateUser)))
	r.Methods("GET").Path("/api/v1/movies").Queries("page", "{page}").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovies)))
	r.Methods("GET").Path("/api/v1/movies/{movieID}").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovie)))
	r.Methods("GET").Path("/api/v1/movies/{movieID}/information").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovieInformation)))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))
}
