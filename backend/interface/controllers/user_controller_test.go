package controllers

import (
	"context"
	"github.com/favclip/testerator"
	"matilda/backend/domain/entity"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"google.golang.org/appengine/datastore"
)

type MockDatastoreUserRepository struct {
}

type MockLogUserRepository struct {
}

func (mockDatastore *MockDatastoreUserRepository) Store(r *http.Request, src interface{}) (*datastore.Key, error) {
	return nil, nil
}

func (mockDatastore *MockDatastoreUserRepository) FindMulti(r *http.Request, src []*entity.User) error {
	t, _ := time.Parse("2006-01-02", "2019-01-01")
	src = append(src, &entity.User{
		UserID:    "test123",
		Name:      "test name",
		IconPath:  "test.jpg",
		Deleted:   false,
		CreatedAt: t,
		UpdatedAt: t,
	})
	src = append(src, &entity.User{
		UserID:    "test456",
		Name:      "test nickname",
		IconPath:  "test2.jpg",
		Deleted:   false,
		CreatedAt: t,
		UpdatedAt: t,
	})
	return nil
}

func (mockLog *MockLogUserRepository) Output(ctx context.Context, format string, args interface{}) {
}

func TestUserController_CreateUser(t *testing.T) {
	t.Skip()
	inst, _, err := testerator.SpinUp()
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

	apErr := UserHandler.CreateUser(res, req)
	if apErr != nil {
		t.Fatalf("CreateUser error: %v", apErr)
	}

	wantStatusCode := 202
	if res.Code != wantStatusCode {
		t.Fatalf("got statusCode: %v, want statusCode: %v", res.Code, wantStatusCode)
	}
}
