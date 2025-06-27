package middleware

import (
	"log"
	"net/http"
	"time"
)

// WithLogging: 핸들러를 감싸서 요청 로그를 찍는 미들웨어
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("요청 시작: %s %s", r.Method, r.URL.Path)

		// 실제 핸들러 실행
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("요청 종료: %s %s (%v)", r.Method, r.URL.Path, duration)
	})
}
