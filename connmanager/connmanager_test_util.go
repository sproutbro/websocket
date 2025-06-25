package connmanager

import (
	"myonly/global"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func connectToServer(t *testing.T, url string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("웹소켓 연결 실패: %v", err)
	}
	return conn
}

func assertConnInMap(t *testing.T, conn *websocket.Conn) {
	found := false
	global.Mutex.Lock()
	for c := range connMap {
		if c == conn {
			found = true
			break
		}
	}
	global.Mutex.Unlock()
	if !found {
		t.Errorf("connMap에 연결이 없습니다")
	}
}

func assertConnNotInMap(t *testing.T, conn *websocket.Conn) {
	global.Mutex.Lock()
	for c := range connMap {
		if c == conn {
			t.Errorf("connMap에서 연결이 제거되지 않았습니다")
			break
		}
	}
	global.Mutex.Unlock()
}

func resetConnMap(t *testing.T) {
	t.Cleanup(func() {
		global.Mutex.Lock()
		connMap = make(map[*websocket.Conn]bool)
		global.Mutex.Unlock()
	})
}

func createEchoServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err, "웹소켓 업그레이드 실패")

		go func() {
			for {
				msgType, msg, err := conn.ReadMessage()
				if err != nil {
					return // 연결 끊어졌으면 종료
				}
				_ = conn.WriteMessage(msgType, msg)
			}
		}()
	}))
}
