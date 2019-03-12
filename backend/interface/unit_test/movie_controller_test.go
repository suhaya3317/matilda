package unit_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/favclip/testerator"

	"google.golang.org/appengine/aetest"
)

type MockMovieMuxRepository struct {
}

func (mux *MockMovieMuxRepository) Find(r *http.Request, key string) string {
	return "5"
}

func TestMovieController_GetMovies(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies?page=5", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+AuthToken)

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovies(res, req)
	if apErr != nil {
		t.Fatalf("GetMovies error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	wantStatusCode := 200
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}

	var dat []map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		t.Fatalf("Failed to json.Unmarshal: %v", err)
	}

	var actual []string
	for k := range dat[0] {
		actual = append(actual, k)
	}

	expected := []string{"movie_id", "title", "poster_path"}
	ok := Equal(actual, expected)
	if !ok {
		t.Fatalf("Failed to equal. got: %v, want: %v", actual, expected)
	}
}

func TestMovieController_GetMovie(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/5", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+AuthToken)

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovie(res, req)
	if apErr != nil {
		t.Fatalf("GetMovie error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	wantStatusCode := 200
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		t.Fatalf("Failed to json.Unmarshal: %v", err)
	}

	var actual []string
	for k := range dat {
		actual = append(actual, k)
	}

	expected := []string{"movie_id", "title", "poster_path"}
	ok := Equal(actual, expected)
	if !ok {
		t.Fatalf("Failed to equal. got: %v, want: %v", actual, expected)
	}
}

func TestMovieController_GetMovieInformation(t *testing.T) {
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/399579/information", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+AuthToken)

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovieInformation(res, req)
	if apErr != nil {
		t.Fatalf("GetMovieInformation error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	wantStatusCode := 200
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		t.Fatalf("Failed to json.Unmarshal: %v", err)
	}

	var actual []string
	for k := range dat {
		actual = append(actual, k)
	}

	expected := []string{"movie_id", "release_date", "director", "cast", "detail"}
	ok := Equal(actual, expected)
	if !ok {
		t.Fatalf("Failed to equal. got: %v, want: %v", actual, expected)
	}
}
