package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

type MockDatastoreUserRepository struct {
}

type MockLogUserRepository struct {
}

func (mockDatastore *MockDatastoreUserRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return nil, nil
}

func (mockLog *MockLogUserRepository) Output(ctx context.Context, format string, args interface{}) {
}

func TestUserController_CreateUser(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("PUT", "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("Authorization", "Bearer: "+os.Getenv("TEST_AUTH_TOKEN"))

	res := httptest.NewRecorder()

	apErr := TargetUser.CreateUser(res, req)
	if apErr != nil {
		t.Fatalf("CreateUser error: %v", apErr)
	}

	wantStatusCode := 202
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}
}
