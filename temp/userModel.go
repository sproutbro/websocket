package temp

import "github.com/gorilla/websocket"

type User struct {
	Key     *websocket.Conn
	Id      string
	Connect bool
}
