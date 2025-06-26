package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
upgradeSocket은 HTTP 요청을 WebSocket 연결로 업그레이드합니다.
성공 시 연결된 클라이언트에 초기 메시지를 전송하며, 이후 메시지 루프를 시작할 수 있습니다.
*/
func upgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("업그레이드 실패: %v", err)
		http.Error(w, "Could not upgrade", http.StatusInternalServerError)
		return
	}

	// 연결 성공 메시지 전송
	log.Println("클라이언트 연결됨!")
	conn.WriteMessage(websocket.TextMessage, []byte("WebSocket 연결 성공!"))
}
