package middleware

import (
	"net/http"
)

// CORS 허용 Origin 설정
var allowedOrigins = map[string]bool{
	"http://localhost:3000": true,
	"https://myapp.com":     true,
}

// 상황마다 바뀌는 부분은 이 allowedOrigins
// 필요시 env 파일 등으로 뺄 수 있음
func WithCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		}

		// Preflight 요청 처리
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// CORS 체크 끝났으니 다음 단계로 넘김
		next.ServeHTTP(w, r)
	}
}
