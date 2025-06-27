// File: handler.go
// Author: 이정균
// Date: 2025-06-28
//
//	모든 WebSocket 접속은 이곳을 거쳐서 연결됩니다.
//
// Description:
//   - websocket.Upgrader - Gorilla WebSocket에서 HTTP → WS 변환 담당
//   - CheckOrigin - CORS 무시 설정 (실무에서는 도메인 체크 필요)
//   - ServeWS - WebSocket 연결을 처리하는 메인 핸들러
//   - hub.Register(client) 연결된 클라이언트를 Hub에 등록
//   - go client.ReadPump() 수신 루프
//   - go client.WritePump() 송신 루프
//
// Related files:
//   - internal/logger
//
// Test Strategy:
//   - 업그레이드 성공 : 클라이언트가 WebSocket으로 정상 연결되는지
//   - 업그레이드 실패 : 업그레이드 불가 상황에서 400 에러 응답
//   - Hub에 등록되는지 : hub.Register()가 호출되는지 확인
//
// Last Modified: 2025-06-28 by JackieChan
package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HubInterface interface {
	Register(client *Client)
}

func ServeWS(hub HubInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "webSocket 업그레이드 실패", http.StatusBadRequest)
			return
		}

		client := NewClient(conn, hub)
		hub.Register(client)

		go client.ReadPump()
		go client.WritePump()
	}
}
