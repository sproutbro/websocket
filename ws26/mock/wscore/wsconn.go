package wscore

import "github.com/gorilla/websocket"

type WebSocketConn interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}

type RealWSConn struct {
	Conn *websocket.Conn
}

func (r *RealWSConn) ReadMessage() (int, []byte, error) {
	return r.Conn.ReadMessage()
}

func (r *RealWSConn) WriteMessage(mt int, msg []byte) error {
	return r.Conn.WriteMessage(mt, msg)
}

func (r *RealWSConn) Close() error {
	return r.Conn.Close()
}
