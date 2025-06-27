package middleware

import (
	"net/http"
	"strings"
)

var validToken = "secrettoken123"

func WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != validToken {
			http.Error(w, "Invalid Token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
