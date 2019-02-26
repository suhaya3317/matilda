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

	moviesAPI := new(domain.MoviesAPI)
	err = json.Unmarshal(body, moviesAPI)
	if err != nil {
		return appErrorf(err, "json.Unmarshal()")
	}

	var domainMovies []*domain.Movie
	for i := range moviesAPI.Results {
		domainMovies = append(domainMovies, &domain.Movie{
			ID:         moviesAPI.Results[i].ID,
			Title:      moviesAPI.Results[i].Title,
			PosterPath: "https://image.tmdb.org/t/p/w300_and_h450_bestv2" + moviesAPI.Results[i].PosterPath,
		})
	}

	err = setResponseWriter(w, 200, domainMovies)
	if err != nil {
		return appErrorf(err, "%v", err)
	}
	controller.LogInterceptor.LogInfo(ctx, "success: %v", "GetMovies()")
	return nil
}

func (controller *MovieController) GetMovie(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)
	res, err := controller.MovieAPIInterceptor.GetMovie(ctx, controller.MuxInterceptor.Get(r, "movieID"))
	if err != nil {
		return appErrorf(err, "controller.MovieAPIInterceptor.GetMovie()")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return appErrorf(err, "ioutil.ReadAll()")
	}
	defer res.Body.Close()

	movieAPI := new(domain.MovieAPI)
	err = json.Unmarshal(body, movieAPI)
	if err != nil {
		return appErrorf(err, "json.Unmarshal()")
	}

	var domainMovie domain.Movie
	domainMovie.ID = movieAPI.ID
	domainMovie.Title = movieAPI.Title
	domainMovie.PosterPath = "https://image.tmdb.org/t/p/w300_and_h450_bestv2" + movieAPI.PosterPath

	err = setResponseWriter(w, 200, domainMovie)
	if err != nil {
		return appErrorf(err, "setResponseWriter()")
	}

	controller.LogInterceptor.LogInfo(ctx, "success: %v", "GetMovie()")
	return nil
}

func (controller *MovieController) GetMovieInformation(w http.ResponseWriter, r *http.Request) *appError {
	ctx := appengine.NewContext(r)
	res, err := controller.MovieAPIInterceptor.GetMovieInformation(ctx, controller.MuxInterceptor.Get(r, "movieID"))
	if err != nil {
		return appErrorf(err, "controller.MovieAPIInterceptor.GetMovie()")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return appErrorf(err, "ioutil.ReadAll()")
	}
	defer res.Body.Close()

	movieAPI := new(domain.MovieAPI)
	err = json.Unmarshal(body, movieAPI)
	if err != nil {
		return appErrorf(err, "json.Unmarshal()")
	}

	var domainMovieInfo domain.MovieInformation
	domainMovieInfo.ID = movieAPI.ID
	domainMovieInfo.ReleaseDate = movieAPI.ReleaseDate

	for i := range movieAPI.Credits.Crew {
		if movieAPI.Credits.Crew[i].Job == "Director" {
			domainMovieInfo.Director = movieAPI.Credits.Crew[i].Name
			break
		}
	}

	for i := range movieAPI.Credits.Cast {
		domainMovieInfo.Cast = append(domainMovieInfo.Cast, movieAPI.Credits.Cast[i].Name)
	}

	domainMovieInfo.Detail = movieAPI.Overview

	err = setResponseWriter(w, 200, domainMovieInfo)
	if err != nil {
		return appErrorf(err, "setResponseWriter()")
	}

	controller.LogInterceptor.LogInfo(ctx, "success: %v", "GetMovieInformation()")
	return nil
}
