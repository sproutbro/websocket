package wstest

import "github.com/gorilla/websocket"

type MockWSConn struct {
	WrittenMessages [][]byte
	Closed          bool
}

func (m *MockWSConn) ReadMessage() (int, []byte, error) {
	return websocket.TextMessage, []byte("mock read"), nil
}

func (m *MockWSConn) WriteMessage(mt int, msg []byte) error {
	m.WrittenMessages = append(m.WrittenMessages, msg)
	return nil
}

func (m *MockWSConn) Close() error {
	m.Closed = true
	return nil
}
