package controllers

import (
	"encoding/json"
	"io/ioutil"
	"matilda/backend/domain"
	"matilda/backend/interface/gorilla_mux"
	"matilda/backend/interface/logging"
	"matilda/backend/interface/movie"
	"matilda/backend/usecase"
	"net/http"

	"google.golang.org/appengine"
)

type MovieController struct {
	MuxInterceptor      usecase.MovieMuxInterceptor
	MovieAPIInterceptor usecase.MovieAPIInterceptor
	LogInterceptor      usecase.LogInterceptor
}

func NewMovieController(gorillaMuxHandler gorilla_mux.GorillaMuxHandler, movieAPIHandler movie.MovieAPIHandler, logHandler logging.LogHandler) *MovieController {
	return &MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &gorilla_mux.MovieMuxRepository{
				GorillaMuxHandler: gorillaMuxHandler,
			},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &movie.MovieAPIRepository{
				MovieAPIHandler: movieAPIHandler,
			},
		},
		LogInterceptor: usecase.LogInterceptor{
			LogRepository: &logging.LogRepository{
				LogHandler: logHandler,
			},
		},
	}
}

func (controller *MovieController) GetMovies(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)
	res, err := controller.MovieAPIInterceptor.GetPopularMovies(ctx, controller.MuxInterceptor.Get(r, "page"))
	if err != nil {
		return appErrorf(err, "controller.MovieAPIInterceptor.GetPopularMovies()")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return appErrorf(err, "ioutil.ReadAll()")
	}
	defer res.Body.Close()

	movies := new(domain.Movies)
	err = json.Unmarshal(body, movies)
	if err != nil {
		return appErrorf(err, "json.Unmarshal()")
	}

	var matildaMovies []*domain.MatildaMovie
	for i := range movies.Results {
		matildaMovies = append(matildaMovies, &domain.MatildaMovie{
			ID:         movies.Results[i].ID,
			Title:      movies.Results[i].Title,
			PosterPath: "https://image.tmdb.org/t/p/w300_and_h450_bestv2" + movies.Results[i].PosterPath,
		})
	}

	err = setResponseWriter(w, 200, matildaMovies)
	if err != nil {
		return appErrorf(err, "%v", err)
	}
	controller.LogInterceptor.LogInfo(ctx, "success: %v", "GetMovies()")
	return nil
}
