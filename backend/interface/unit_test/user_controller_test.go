package unit_test

import (
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

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
	req.Header.Set("Authorization", "Bearer "+AuthToken)

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
