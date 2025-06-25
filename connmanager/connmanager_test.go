package connmanager

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func createTestServer(t *testing.T) string {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			upgrader := websocket.Upgrader{}
			_, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				t.Fatalf("웹소켓 업그레이드 실패: %v", err)
			}
		}))
	return "ws" + s.URL[4:]
}

func TestRegister(t *testing.T) {
	url := createTestServer(t)

	conn := connectToServer(t, url)
	defer conn.Close()

	resetConnMap(t)

	Register(conn)

	assertConnInMap(t, conn)

	Unregister(conn)

	assertConnNotInMap(t, conn)
}

func TestWebSocketReadAndWrite(t *testing.T) {
	server := createEchoServer(t)
	defer server.Close()

	// 클라이언트 연결
	wsURL := "ws" + server.URL[4:]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "웹소켓 연결 실패")
	defer conn.Close()

	// 클라이언트 → 서버로 메시지 전송
	testMessage := "hello websocket"
	err = conn.WriteMessage(websocket.TextMessage, []byte(testMessage))
	require.NoError(t, err, "메시지 전송 실패")

	// 서버가 다시 보내주는 메시지를 수신
	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // 안오면 실패
	msgType, msg, err := conn.ReadMessage()

	t.Log("기대한 메시지: , 받은 메시지: ", testMessage, string(msg))

	require.NoError(t, err, "메시지 수신 실패")
	require.Equal(t, websocket.TextMessage, msgType)
	require.Equal(t, testMessage, string(msg))
}
