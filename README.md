# WebSocket 서버 구조 요약

## 디자인 패턴 적용 구조

### 🔸 Singleton - Logger
- 위치: `logger/logger.go`
- 책임: 전역에서 사용되는 로거
- 수정 시:
  - 로그 포맷 변경 → `logger.go` 내부 `Log()` 함수 수정
  - 출력 방식 변경 (파일 등) → `NewLogger()` 내 I/O 변경

### 🔸 Factory Method - HTTP Router
- 위치: `router/factory.go`
- 책임: 요청 경로별 핸들러 매핑
- 수정 시:
  - 새로운 URL 경로 추가 → `CreateHandler(path string)` switch문 수정
  - 기존 핸들러 수정 → 각 핸들러 구현체 수정 (예: `UserHandler`)

### 🔸 Decorator - Middleware
- 위치: `middleware/`
- 책임: 공통 처리 (CORS, Auth, Logging 등)
- 수정 시:
  - 인증 로직 수정 → `middleware/auth.go`
  - CORS 도메인 변경 → `middleware/cors.go` 내부 허용 목록 수정
  - 로깅 방식 변경 → `middleware/logging.go`

### 📌 Mock Object Pattern (모의 객체 패턴)
- 설명
  - 의존성을 제거한 테스트를 위해 실제 객체 대신 가짜(Mock) 객체를 사용합니다.
  - 주로 Logger, DB, 외부 API 등 테스트가 어려운 부분을 대체합니다.
- 위치:
  - `logger.Logger/`
  - `mockLogger.MockLogger/`

### 📌 Mock Object Pattern (모의 객체 패턴)
- 설명
  - 하나의 테스트 함수에서 여러 입력/출력 케이스를 반복해서 테스트하는 기법입니다.
  - Go 커뮤니티에서 권장되는 테스트 구조입니다.
- 위치:
   -  `router_test.go`
- 장점
   - 테스트 케이스가 한눈에 보임
   - 반복 코드 최소화
   - 케이스 추가가 쉬움

```go
tests := []struct {
    name string
    input string
    want int
}{...}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        ...
    })
}
```

### 📌 Middleware Handler → WebSocket 업그레이드 연결 구조
```go
/middleware/handler.go       // WSHandler() 함수
    ↓
/middleware/upgradeSocket()     // 업그레이드 핵심 로직
    ↑
WithAuth            // 인증 체크
    ↑
WithCORS            // CORS 설정
```

### 📍 Middleware Handler → Factory Router 핸들러 연결 구조
```go
/router/factory.go   // CreateHandler(route string) 함수
    ↓
switch case          → "chat", "user", "admin" 등 핸들러 선택
    ↓
/middleware/handler.go       // Handler() 함수
    ↑
WithAuth            // 인증 체크
    ↑
WithCORS            // CORS 설정
```