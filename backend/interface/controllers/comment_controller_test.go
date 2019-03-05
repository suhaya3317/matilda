package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"matilda/backend/domain/entity"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"google.golang.org/appengine/aetest"

	"google.golang.org/appengine/datastore"
)

type MockMuxCommentRepository struct {
}

type MockDatastoreCommentRepository struct {
}

type MockLogCommentRepository struct {
}

func (mockMux *MockMuxCommentRepository) Find(r *http.Request, key string) string {
	return "550"
}

func (mockDatastore *MockDatastoreCommentRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return nil, nil
}

func (mockDatastore *MockDatastoreCommentRepository) FindKey(r *http.Request, src interface{}) *datastore.Key {
	key := &datastore.Key{}
	return key
}

func (mockLog *MockLogCommentRepository) Output(ctx context.Context, format string, args interface{}) {
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
