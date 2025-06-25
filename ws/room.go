package ws

import "github.com/gorilla/websocket"

var (
	roomMap = make(map[string]map[string]*websocket.Conn)
)
