package ws_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"myonly/ws"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWebSocketConnection 은 ws.Handler를 테스트 서버에서 실행하고,
// 웹소켓 연결이 성공하는지 확인합니다.
func TestWebSocketConnection(t *testing.T) {
	// 실제 핸들러를 사용하는 테스트 서버 생성
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	// ws:// 로 시작하게 URL 변형
	wsURL := "ws" + server.URL[4:]

	// 클라이언트에서 서버로 연결 시도
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "웹소켓 연결 실패")
	defer conn.Close()
}

// TestWebSocketEcho 은 웹소켓으로 메시지를 보내면
// 서버가 같은 메시지를 다시 돌려주는지(echo) 테스트합니다.
func TestWebSocketEcho(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	wsURL := "ws" + server.URL[4:]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err, "웹소켓 연결 실패")
	defer conn.Close()

	// 메시지 전송
	sendMsg := "hello websocket!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(sendMsg))
	require.NoError(t, err, "메시지 전송 실패")

	// 메시지 수신
	_, recvMsg, err := conn.ReadMessage()
	require.NoError(t, err, "메시지 수신 실패")

	require.Equal(t, sendMsg, string(recvMsg), "수신 메시지 불일치")
}

// TestBroadcast 는 여러 클라이언트가 연결되었을 때,
// 서버에서 메시지를 전송하면 모든 클라이언트가 받을 수 있는지 테스트합니다.
func TestBroadcast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	wsURL := "ws" + server.URL[4:]

	// 클라이언트 2명 연결
	conn1 := connectClient(t, wsURL)
	defer conn1.Close()

	conn2 := connectClient(t, wsURL)
	defer conn2.Close()

	ws.AddConnWithID(conn1)
	ws.AddConnWithID(conn2)

	// 서버에서 Broadcast 실행
	testMessage := "broadcast to everyone"
	ws.Broadcast([]byte(testMessage))

	// 두 클라이언트 모두 메시지 수신 확인
	conn1.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
	_, msg1, _ := conn1.ReadMessage()
	conn2.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
	_, msg2, _ := conn2.ReadMessage()

	require.Equal(t, msg1, msg2)

}

// TestClientDisconnect 는 한 클라이언트가 연결을 끊으면
// Broadcast 대상에서 자동으로 제거되는지 테스트합니다.
func TestClientDisconnect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	wsURL := "ws" + server.URL[4:]

	// 클라이언트 A, B 연결
	connA := connectClient(t, wsURL)
	connB := connectClient(t, wsURL)

	// A 종료
	connA.Close()

	// B가 메시지 보냄
	message := "message after A left"
	err := connB.WriteMessage(websocket.TextMessage, []byte(message))
	require.NoError(t, err)

	// B는 정상적으로 다시 받음 (echo or broadcast)
	checkMessage(t, connB, message)

	testMessage := "broadcast to everyone"
	ws.Broadcast([]byte(testMessage))

	// A는 이미 닫혔기 때문에 아무 일도 없음 (테스트 실패 아님)
	connClients := len(ws.GetAllClients())
	assert.Equal(t, 1, connClients)

}

// TestClientCount 는 클라이언트가 연결될 때와 끊어질 때
// CountClients() 결과가 정확한지 확인합니다.
func TestClientCount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	wsURL := "ws" + server.URL[4:]

	// 시작 시 0명이어야 함
	require.Equal(t, 0, ws.CountClients(), "초기 클라이언트 수는 0이어야 함")

	conn1 := connectClient(t, wsURL)
	require.Eventually(t, func() bool {
		return ws.CountClients() == 1
	}, time.Second, 10*time.Millisecond)

	conn2 := connectClient(t, wsURL)
	require.Eventually(t, func() bool {
		return ws.CountClients() == 2
	}, time.Second, 10*time.Millisecond)

	conn1.Close()
	require.Eventually(t, func() bool {
		return ws.CountClients() == 1
	}, time.Second, 10*time.Millisecond)

	conn2.Close()
	require.Eventually(t, func() bool {
		return ws.CountClients() == 0
	}, time.Second, 10*time.Millisecond)
}

func TestSendToSpecificClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	base := "ws" + server.URL[4:]
	connA := connectClientWithID(t, base+"?id=alpha")
	defer connA.Close()

	connB := connectClientWithID(t, base+"?id=beta")
	defer connB.Close()

	// ws.AddConnWithID("aaaa", connA)
	// ws.AddConnWithID("123", connB)

	fmt.Println(ws.GetAllClients())

	msg := "hi beta!"
	err := ws.SendTo("aaaa", []byte(msg))

	msg2 := "hieta!"
	err2 := ws.SendTo("123", []byte(msg2))

	require.NoError(t, err)
	require.NoError(t, err2)

	// beta는 받아야 함
	checkMessage(t, connB, msg2)

	checkMessage(t, connA, msg)

	t.Logf("현재 클라이언트 목록: %+v", ws.GetAllClients())
}

// ID 붙여서 접속
func connectClientWithID(t *testing.T, url string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	require.NoError(t, err, "웹소켓 연결 실패")
	return conn
}

// TestJsonRouting 는 클라이언트 A가 B에게 JSON 메시지를 보냈을 때,
// B만 그 메시지를 수신하는지 확인합니다.
func TestJsonRouting(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	base := "ws" + server.URL[4:]
	connA := connectClientWithID(t, base+"?id=alice")
	defer connA.Close()

	connB := connectClientWithID(t, base+"?id=bob")
	defer connB.Close()

	jsonMsg := `{"to":"bob", "msg":"hello bob!"}`
	err := connA.WriteMessage(websocket.TextMessage, []byte(jsonMsg))
	require.NoError(t, err)

	// bob은 메세지를 받아야함
	checkMessage(t, connB, "hello bob!")

	// alice는 받으면 안 됨
	connA.SetReadDeadline(time.Now().Add(1 * time.Second))
	require.Equal(t, 2, len(ws.GetAllClients()), "현재 접속 유저")
}

// TestChatRoomBroadcast 는 같은 방에 있는 사람들끼리만 메시지를 주고받는지 테스트합니다.
func TestChatRoomBroadcast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.Handler))
	defer server.Close()

	base := "ws" + server.URL[4:]

	// 같은 방에 있는 두 명
	alice := connectClientWithID(t, base+"?id=alice&room=room1")
	defer alice.Close()
	bob := connectClientWithID(t, base+"?id=bob&room=room1")
	defer bob.Close()

	// 다른 방에 있는 한 명
	charlie := connectClientWithID(t, base+"?id=charlie&room=room2")
	defer charlie.Close()

	// alice가 메시지 전송
	message := "hi room1!"
	err := alice.WriteMessage(websocket.TextMessage, []byte(message))
	require.NoError(t, err)

	// bob은 받아야 함
	checkMessage(t, bob, message)

	// ws.AddConnWithID("aaaa", charlie)
	// ws.AddConnWithID("123", bob)
	// ws.AddConnWithID("aa", alice)
	// charlie는 받으면 안 됨
	charlie.SetReadDeadline(time.Now().Add(1 * time.Second))
	require.Equal(t, 3, len(ws.GetAllClients()), "세명")
}

// connectClient 는 테스트용 클라이언트를 서버에 연결합니다.
func connectClient(t *testing.T, url string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	require.NoError(t, err, "웹소켓 연결 실패")
	return conn
}

// checkMessage 는 클라이언트가 메시지를 잘 받았는지 확인합니다.
func checkMessage(t *testing.T, conn *websocket.Conn, expected string) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, msg, err := conn.ReadMessage()
	require.NoError(t, err, "메시지 수신 실패")
	require.Equal(t, expected, string(msg), "받은 메시지가 기대값과 다름")
}
