package router_test

import (
	"myonly/internal/httpserver/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	mux := router.NewRouter() // 라우터 생성 (핸들러 + 미들웨어)

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "루트 경로 응답 테스트",
			method:     http.MethodGet,
			path:       "/",
			wantStatus: http.StatusOK,
			wantBody:   "Hello from Router!",
		},
		{
			// 전체경로 일치
			name:       "없는 경로 응답 테스트",
			method:     http.MethodGet,
			path:       "/not-found",
			wantStatus: http.StatusOK,
			wantBody:   "Hello from Router!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			// 상태코드 확인
			if rec.Code != tt.wantStatus {
				t.Errorf("원하는 상태코드: %d, 실제: %d", tt.wantStatus, rec.Code)
			}

			// 응답 본문 확인
			if !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf("본문 불일치. 기대값: %q, 실제: %q", tt.wantBody, rec.Body.String())
			}
		})
	}
}

/*
Mock 오브젝트 패턴과
모의객체 패턴이 있는 테스트 코드
httptest 를 이용하여 짭퉁 GET 요청
응답 body, status 를 비교하여 일치하면 통과
응답 값은 router 폴더의 각각의 라우터에 있음
*/
// func TestCreateHandler_ReturnsCorrectType(t *testing.T) {

// 	for _, tc := range tests {
// 		// 핸들러 객체 생성

// 		// Serve 호출

// 	}
// }
