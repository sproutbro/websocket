package ws26socket

import "github.com/gorilla/websocket"

type Client struct {
	ID     string
	Conn   *websocket.Conn
	RoomID string
}

func NewClient(id string, conn *websocket.Conn, roomID string) *Client {
	return &Client{
		ID:     id,
		Conn:   conn,
		RoomID: roomID,
	}
}
