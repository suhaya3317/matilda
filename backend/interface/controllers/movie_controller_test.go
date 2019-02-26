package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"matilda/backend/domain"
	"matilda/backend/usecase"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"google.golang.org/appengine/aetest"
)

var target MovieController

func TestMain(m *testing.M) {
	target = MovieController{
		MuxInterceptor: usecase.MovieMuxInterceptor{
			MovieMuxRepository: &MockMovieMuxRepository{},
		},
		MovieAPIInterceptor: usecase.MovieAPIInterceptor{
			MovieAPIRepository: &MockMovieAPIRepository{},
		},
		LogInterceptor: usecase.LogInterceptor{
			LogRepository: &MockLogRepository{},
		},
	}
	code := m.Run()
	os.Exit(code)
}

type MockMovieMuxRepository struct {
}

type MockMovieAPIRepository struct {
}

type MockLogRepository struct {
}

func (mux *MockMovieMuxRepository) Find(r *http.Request, key string) string {
	return "2"
}

func (movie *MockMovieAPIRepository) FindAll(ctx context.Context, page string) (*http.Response, error) {
	var moviesAPI domain.MoviesAPI
	moviesAPI = domain.MoviesAPI{
		Page:         1,
		TotalResults: 19801,
		TotalPages:   991,
		Results: []domain.Results{
			{
				VoteCount:        650,
				ID:               399579,
				Video:            false,
				VoteAverage:      6.7,
				Title:            "Alita: Battle Angel",
				Popularity:       362.08,
				PosterPath:       "/xRWht48C2V8XNfzvPehyClOvDni.jpg",
				OriginalLanguage: "en",
				OriginalTitle:    "Alita: Battle Angel",
				GenreIDs:         []int{28, 878, 53},
				BackdropPath:     "/aQXTw3wIWuFMy0beXRiZ1xVKtcf.jpg",
				Adult:            false,
				Overview:         "When Alita awakens with no memory of who she is in a future world she does not recognize, she is taken in by Ido, a compassionate doctor who realizes that somewhere in this abandoned cyborg shell is the heart and soul of a young woman with an extraordinary past.",
				ReleaseDate:      "2019-01-31",
			},
		},
	}
	body, err := json.Marshal(moviesAPI)
	if err != nil {
		return nil, err
	}

	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}
	return res, nil
}

func (movie *MockMovieAPIRepository) Find(cxt context.Context, id string) (*http.Response, error) {
	var movieAPI domain.MovieAPI
	movieAPI = domain.MovieAPI{
		Adult:        false,
		BackdropPath: "/aQXTw3wIWuFMy0beXRiZ1xVKtcf.jpg",
		BelongsToCollection: domain.BelongsToCollection{
			ID:           0,
			Name:         "",
			PosterPath:   "",
			BackdropPath: "",
		},
		Budget: 170000000,
		Genres: []domain.Genres{
			{
				ID:   28,
				Name: "Action",
			},
			{
				ID:   878,
				Name: "Science Fiction",
			},
			{
				ID:   53,
				Name: "Thriller",
			},
		},
		Homepage:         "https://www.foxmovies.com/movies/alita-battle-angel",
		ID:               399579,
		ImdbId:           "tt0437086",
		OriginalLanguage: "en",
		OriginalTitle:    "Alita: Battle Angel",
		Overview:         "When Alita awakens with no memory of who she is in a future world she does not recognize, she is taken in by Ido, a compassionate doctor who realizes that somewhere in this abandoned cyborg shell is the heart and soul of a young woman with an extraordinary past.",
		Popularity:       492.344,
		PosterPath:       "/xRWht48C2V8XNfzvPehyClOvDni.jpg",
		ProductionCompanies: []domain.ProductionCompanies{
			{
				ID:            10807,
				LogoPath:      "/j0BcMaJKIiDDYHq9lriTcM0Npka.png",
				Name:          "Troublemaker Studios",
				OriginCountry: "US",
			},
			{
				ID:            574,
				LogoPath:      "/iB6GjNVHs5hOqcEYt2rcjBqIjki.png",
				Name:          "Lightstorm Entertainment",
				OriginCountry: "US",
			},
		},
		ProductionCountries: []domain.ProductionCountries{
			{
				Iso_3166_1: "AR",
				Name:       "Argentina",
			},
		},
		ReleaseDate: "2019-01-31",
		Revenue:     263358391,
		Runtime:     122,
		SpokenLanguages: []domain.SpokenLanguages{
			{
				Iso_639_1: "en",
				Name:      "English",
			},
		},
		Status:      "Released",
		Tagline:     "An angel falls. A warrior rises.",
		Title:       "Alita: Battle Angel",
		Video:       false,
		VoteAverage: 6.9,
		VoteCount:   891,
	}

	body, err := json.Marshal(movieAPI)
	if err != nil {
		return nil, err
	}

	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}
	return res, nil
}

func (log *MockLogRepository) Output(ctx context.Context, format string, args interface{}) {
}

func TestMovieController_GetMovies(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies?page=2", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	res := httptest.NewRecorder()

	apErr := target.GetMovies(res, req)
	if apErr != nil {
		t.Fatalf("GetMovies error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	if res.Code != 200 {
		t.Fatalf("StatusCode: %v, Response.Body: %v", res.Code, body)
	}

	actual := string(body)
	expected := `[{"id":399579,"title":"Alita: Battle Angel","poster_path":"https://image.tmdb.org/t/p/w300_and_h450_bestv2/xRWht48C2V8XNfzvPehyClOvDni.jpg"}]`

	err = IsEqualJSON(actual, expected)
	if err != nil {
		t.Fatalf("Failed to IsEqualJson: %v", err)
	}
}

func TestMovieController_GetMovie(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/399579", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	res := httptest.NewRecorder()

	apErr := target.GetMovie(res, req)
	if apErr != nil {
		t.Fatalf("GetMovie error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	if res.Code != 200 {
		t.Fatalf("StatusCode: %v, Response.Body: %v", res.Code, body)
	}

	actual := string(body)
	expected := `{"id":399579,"title":"Alita: Battle Angel","poster_path":"https://image.tmdb.org/t/p/w300_and_h450_bestv2/xRWht48C2V8XNfzvPehyClOvDni.jpg"}`

	err = IsEqualJSON(actual, expected)
	if err != nil {
		t.Fatalf("Failed to IsEqualJson: %v", err)
	}
}

func IsEqualJSON(a, b string) error {
	err, ok := DeepEqualJSON(a, b)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	err = errors.New("not Equal")
	return err
}

func DeepEqualJSON(j1, j2 string) (error, bool) {
	var err error

	var d1 interface{}
	err = json.Unmarshal([]byte(j1), &d1)

	if err != nil {
		return err, false
	}

	var d2 interface{}
	err = json.Unmarshal([]byte(j2), &d2)

	if err != nil {
		return err, false
	}

	if reflect.DeepEqual(d1, d2) {
		return nil, true
	} else {
		return nil, false
	}
}