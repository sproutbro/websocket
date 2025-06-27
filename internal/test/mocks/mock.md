```go
router := router.NewRouter()

// 가짜요청해더
httptest.NewRequest("GET", "/", nil)

// 기억레코더
httptest.NewRecorder()

// 결과 확인
resp := w.Result()
body := w.Body.String()

// 요청 보내기
mux.ServeHTTP(rec, req)

// 비교코드
if status := resp.StatusCode; status != tc.statusCode {
	t.Errorf("status code mismatch: got %d, want %d", status, tc.statusCode)
}
if !strings.Contains(body, tc.wantOutput) {
	t.Errorf("response body mismatch: got %q, want %q", body, tc.wantOutput)
}
```

```go
	tests := []struct {
		path       string
		wantOutput string
		statusCode int
	}{
		{"/users", "유저 목록입니다.\n", http.StatusOK},
		{"/posts", "게시글 목록입니다.\n", http.StatusOK},
		{"/unknown", "404 page not found\n", http.StatusNotFound},
	}
	if tests != nil {
		// log.Fatal()
	}
```

```go
route := []struct {
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
		name:       "없는 경로 응답 테스트",
		method:     http.MethodGet,
		path:       "/not-found",
		wantStatus: http.StatusNotFound,
		wantBody:   "404 page not found",
	},
}
```