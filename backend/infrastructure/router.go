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

	controllers.Common = controllers.NewInternalController(firebaseHandler)
	firebaseController := controllers.NewFirebaseController(firebaseHandler)
	controllers.UserHandler = controllers.NewUserController(datastoreHandler, logHandler)
	movieController := controllers.NewMovieController(gorillaMuxHandler, movieAPIHandler, logHandler)
	commentController := controllers.NewCommentController(gorillaMuxHandler, datastoreHandler, logHandler)

	r.Methods("PUT").Path("/api/v1/users").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(controllers.UserHandler.CreateUser)))
	r.Methods("GET").Path("/api/v1/movies").Queries("page", "{page}").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovies)))
	r.Methods("GET").Path("/api/v1/movies/{movieID}").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovie)))
	r.Methods("GET").Path("/api/v1/movies/{movieID}/information").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(movieController.GetMovieInformation)))
	r.Methods("PUT").Path("/api/v1/movies/{movieID}/comments").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(commentController.CreateComment)))
	r.Methods("GET").Path("/api/v1/movies/{movieID}/comments").Queries("page", "{page}").
		Handler(firebaseController.AuthMiddleware(controllers.AppHandler(commentController.GetComments)))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))
}
