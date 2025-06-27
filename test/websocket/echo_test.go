package websocket_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

// 테스트 서버용 WebSocket 핸들러 (Echo)
func echoHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 테스트용 CORS 허용
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		t := w.(http.Flusher)
		t.Flush()
		return
	}
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		_ = conn.WriteMessage(mt, msg) // Echo
	}
}

// 테스트 서버용 WebSocket 핸들러 (Echo)
func TestWebSocketEcho(t *testing.T) {
	// 1. 가짜 HTTP 서버 생성
	server := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer server.Close()

	// 2. 실제 WebSocket URL로 변환 (ws:// → http://에서 호스트 추출)
	url := "ws" + server.URL[4:] + "/ws"

	// 3. 클라이언트 연결 (Dial)
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("WebSocket 연결 실패: %v", err)
	}
	defer ws.Close()

	// 4. 테스트 메시지
	msgToSend := "형님 테스트 메시지!"

	// 5. 메시지 전송
	err = ws.WriteMessage(websocket.TextMessage, []byte(msgToSend))
	if err != nil {
		t.Fatalf("메시지 전송 실패: %v", err)
	}

	// 6. 응답 수신
	_, msgReceived, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("메시지 수신 실패: %v", err)
	}

	// 7. 결과 비교
	if string(msgReceived) != msgToSend {
		t.Errorf("받은 메시지 다름. 예상: %s, 실제: %s", msgToSend, msgReceived)
	}
}
