package unit_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"matilda/backend/domain/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

type MockMuxCommentRepository struct{}

func (mock *MockMuxCommentRepository) Find(r *http.Request, key string) string {
	return "550"
}

func TestCommentController_CreateComment(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	c, err := json.Marshal(entity.Comment{CommentText: "test comment!"})
	if err != nil {
		t.Fatalf("Failed to json.Marshal: %v", err)
	}
	req, err := inst.NewRequest("PUT", "/api/v1/movies/550/comments", bytes.NewBuffer(c))
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	res := httptest.NewRecorder()

	apErr := TargetComment.CreateComment(res, req)
	if apErr != nil {
		t.Fatalf("CreateComment error: %v", apErr)
	}

	wantStatusCode := 202
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}
}

func TestCommentController_GetComments(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	err = CreateSeedData(inst)
	if err != nil {
		t.Fatalf("Failed to create seed data: %v", err)
	}

	req, err := inst.NewRequest("GET", "/api/v1/movies/550/comments", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	res := httptest.NewRecorder()

	apErr := TargetComment.GetComments(res, req)
	if apErr != nil {
		t.Fatalf("GetComments error: %v", apErr)
	}

	wantStatusCode := 200
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read res.Body: %v", err)
	}

	actual := string(body)
	expected := `[{"comment_id":5629499534213120,"comment_text":"test comment!","mine":true,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","user_id":"test-sub","name":"test name","icon_path":"test path"}]`

	err, ok := IsEqualJSON(actual, expected)
	if err != nil {
		t.Fatalf("IsEqualJSON error: %v", err)
	}
	if !ok {
		t.Fatalf("not equal json. actual: %v, expected: %v", actual, expected)
	}
}
