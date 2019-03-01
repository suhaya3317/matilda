package controllers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine/urlfetch"
)

const (
	clientCertURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
)

func getSub(r *http.Request, ctx context.Context) (string, bool, error) {
	authHeader := r.Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	keys, err := fetchPublicKeys(ctx)
	if err != nil {
		return "", false, err
	}
	parsedToken, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid := token.Header["kid"]
		rsaPublicKey := convertKey(string(*keys[kid.(string)]))
		return rsaPublicKey, nil
	})
	if err != nil {
		return "", false, err
	}

	tokenMap := parsedToken.Claims.(jwt.MapClaims)
	sub, ok := tokenMap["sub"].(string)

	return sub, ok, nil
}

func fetchPublicKeys(ctx context.Context) (map[string]*json.RawMessage, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get(clientCertURL)
	if err != nil {
		return nil, err
	}

	var objmap map[string]*json.RawMessage
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&objmap)

	return objmap, err
}

func convertKey(key string) interface{} {
	certPEM := key
	certPEM = strings.Replace(certPEM, "\\n", "\n", -1)
	certPEM = strings.Replace(certPEM, "\"", "", -1)
	block, _ := pem.Decode([]byte(certPEM))
	cert, _ := x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	return rsaPublicKey
}

func setResponseWriter(w http.ResponseWriter, statusCode int, src interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(src)
}
