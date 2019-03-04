package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func getidToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	return idToken
}

func decodePublicKeys(resp *http.Response) (map[string]*json.RawMessage, error) {
	var objmap map[string]*json.RawMessage
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&objmap)
	err = errors.Wrap(err, "decoder.Decode()")

	return objmap, err
}

func setResponseWriter(w http.ResponseWriter, statusCode int, src interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(src)
}
