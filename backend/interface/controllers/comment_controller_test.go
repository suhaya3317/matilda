package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/favclip/testerator"
	"matilda/backend/domain/entity"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockMuxCommentRepository struct {
}

type MockLogCommentRepository struct {
}

func (mockMux *MockMuxCommentRepository) Find(r *http.Request, key string) string {
	return "550"
}

func (mockLog *MockLogCommentRepository) Output(ctx context.Context, format string, args interface{}) {
}

func TestCommentController_CreateComment(t *testing.T) {
	t.Skip()
	c, err := json.Marshal(entity.Comment{CommentText: "test comment!"})
	if err != nil {
		t.Fatalf("Failed to json.Marshal: %v", err)
	}
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("PUT", "/api/v1/movies/550/comments", bytes.NewBuffer(c))
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

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
	t.Skip()
	inst, _, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/api/v1/movies/550/comments", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

	res := httptest.NewRecorder()

	apErr := TargetComment.GetComments(res, req)
	if apErr != nil {
		t.Fatalf("GetComments error: %v", err)
	}

	/*
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Failed to read res.Body: %v", err)
		}
	*/

	wantStatusCode := 200
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}

	/*
		actual := string(body)
		expected := `[{"comment_id":12345,"comment_text":"test comment 1","mine":false,"created_at":"2019-01-01T00:00:00.000000Z","updated_at":"2019-01-01T00:00:00.000000Z","user_id":"test123","name":"test name","icon_path":"test.jpg"},{"comment_id":67890,"comment_text":"test comment 2","mine":false,"created_at":"2019-01-01T00:00:00.000000Z","updated_at":"2019-01-01T00:00:00.000000Z","user_id":"test456","name":"test nickname","icon_path":"test2.jpg"}]`

		err = IsEqualJSON(actual, expected)
		if err != nil {
			t.Fatalf("Failed to IsEqualJson. error: %v, got: %v, want: %v", err, actual, expected)
		}
	*/
}
