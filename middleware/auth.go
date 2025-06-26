package middleware

import (
	"net/http"
	"strings"
)

var validToken = "secrettoken123"

// WithAuth는 Authorization 헤더를 검사하여 인증되지 않은 요청을 차단하는 미들웨어입니다.
// 올바른 Bearer 토큰이 없을 경우 401 또는 403 응답을 반환합니다.
func WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// 토큰이 없거나 형식이 이상한 경우
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 실제 토큰만 추출해서 비교
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != validToken {
			http.Error(w, "Invalid Token", http.StatusForbidden)
			return
		}

		// 인증 통과 → 다음 단계로 진행
		next.ServeHTTP(w, r)
	}
}
