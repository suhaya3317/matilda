package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/favclip/testerator"
)

type MockMovieMuxRepository struct {
}

func (mux *MockMovieMuxRepository) Find(r *http.Request, key string) string {
	return "1"
}

func TestMovieController_GetMovies(t *testing.T) {
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies?page=1", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovies(res, req)
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

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		t.Fatalf("Failed to json.Unmarshal: %v", err)
	}
	var actual []string
	for k := range dat {
		actual = append(actual, k)
	}

	expected := []string{"comment_id", "title", "poster_path"}

	ok := reflect.DeepEqual(actual, expected)
	if !ok {
		t.Fatalf("Failed to reflect.DeepEqual. got: %v, want: %v", actual, expected)
	}
}

func TestMovieController_GetMovie(t *testing.T) {
	t.Skip()
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/399579", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovie(res, req)
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
	expected := `{"movie_id":399579,"title":"Alita: Battle Angel","poster_path":"https://image.tmdb.org/t/p/w300_and_h450_bestv2/xRWht48C2V8XNfzvPehyClOvDni.jpg"}`

	err = IsEqualJSON(actual, expected)
	if err != nil {
		t.Fatalf("Failed to IsEqualJson: %v", err)
	}
}

func TestMovieController_GetMovieInformation(t *testing.T) {
	t.Skip()
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/399579/information", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

	res := httptest.NewRecorder()

	apErr := TargetMovie.GetMovieInformation(res, req)
	if apErr != nil {
		t.Fatalf("GetMovieInformation error: %v", apErr)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	if res.Code != 200 {
		t.Fatalf("StatusCode: %v, Response.Body: %v", res.Code, body)
	}

	actual := string(body)
	expected := `{"movie_id":399579,"release_date":"2019-01-31","director":"Robert Rodriguez","cast":["Rosa Salazar"],"detail":"When Alita awakens with no memory of who she is in a future world she does not recognize, she is taken in by Ido, a compassionate doctor who realizes that somewhere in this abandoned cyborg shell is the heart and soul of a young woman with an extraordinary past."}`

	err = IsEqualJSON(actual, expected)
	if err != nil {
		t.Fatalf("Failed to IsEqualJson: %v", err)
	}
}
