| 항목        | 내용                                             |
| ----------- | ------------------------------------------------ |
| 기능        | HTTP → WebSocket 업그레이드 처리                 |
| 책임        | 클라이언트 생성, Hub에 등록, 수신/전송 루프 시작 |
| 테스트 대상 | 업그레이드 성공 여부, 잘못된 요청 처리           |




## 🧱 목표 구조

    | cmd/
    | └── server/
    |     └── main.go              ← 서버 실행
    | 
    | internal/
    | └── websocket/
    |     ├── hub.go               ← Hub (중앙 관리)
    |     ├── client.go            ← Client 구조체
    |     ├── message.go           ← 메시지 구조
    |     └── handler.go           ← HTTP → WebSocket 업그레이드 핸들러
    | 
    | pkg/
    | └── logger/                  ← 기존 Logger 모듈 재사용

## 🔥 최종 목표: 실시간 채팅 기능

| 항목                | 설명                                                  |
| ------------------- | ----------------------------------------------------- |
| 업그레이드          | `http.Handler` → `WebSocket` 연결로 전환              |
| 메시지 브로드캐스트 | 한 명이 메시지 보내면 모든 유저에게 전파              |
| Hub 설계            | `clients`, `broadcast`, `register`, `unregister` 관리 |
| TDD 설계            | 메시지 포맷 검증, 클라이언트 상태 테스트              |

## 🧪 테스트 기반 설계 방식

| 항목         | 테스트 내용                        |
| ------------ | ---------------------------------- |
| 메시지 파싱  | JSON → 메시지 구조체 변환          |
| 브로드캐스트 | 여러 클라이언트에 메시지 전달 확인 |
| 연결 테스트  | 접속/종료 시 Hub 상태 변화 확인    |
| Mock Client  | 가짜 커넥션으로 동작 검증          |

## 🔧 변경/생성 파일 (예정)

| 경로                             | 설명                        |
| -------------------------------- | --------------------------- |
| `internal/websocket/hub.go`      | 브로드캐스트 중심 로직      |
| `internal/websocket/client.go`   | 클라이언트 구조체           |
| `internal/websocket/message.go`  | 메시지 정의 및 파싱         |
| `internal/websocket/handler.go`  | 업그레이드 처리 및 핸들러   |
| `internal/websocket/hub_test.go` | TDD 기반 메시지 전송 테스트 |
