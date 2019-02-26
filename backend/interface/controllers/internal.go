package controllers

import (
	"encoding/json"
	"net/http"
)

func setResponseWriter(w http.ResponseWriter, statusCode int, src interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(src)
}
