package router_test

import (
	"myonly/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Mock 오브젝트 패턴과
모의객체 패턴이 있는 테스트 코드
httptest 를 이용하여 짭퉁 GET 요청
응답 body, status 를 비교하여 일치하면 통과
응답 값은 router 폴더의 각각의 라우터에 있음
*/
func TestCreateHandler_ReturnsCorrectType(t *testing.T) {
	tests := []struct {
		path       string
		wantOutput string
		statusCode int
	}{
		{"/users", "유저 목록입니다.\n", http.StatusOK},
		{"/posts", "게시글 목록입니다.\n", http.StatusOK},
		{"/unknown", "404 page not found\n", http.StatusNotFound},
	}

	for _, tc := range tests {
		// 핸들러 객체 생성
		h := router.CreateHandler(tc.path)

		// 가짜 응답 기록기 & 요청
		req := httptest.NewRequest("GET", tc.path, nil)
		w := httptest.NewRecorder()

		// Serve 호출
		h.Serve(w, req)

		// 결과 확인
		resp := w.Result()
		body := w.Body.String()

		if status := resp.StatusCode; status != tc.statusCode {
			t.Errorf("status code mismatch: got %d, want %d", status, tc.statusCode)
		}
		if !strings.Contains(body, tc.wantOutput) {
			t.Errorf("response body mismatch: got %q, want %q", body, tc.wantOutput)
		}
	}
}
