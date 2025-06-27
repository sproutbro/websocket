package websocket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	hub  HubInterface
	send chan []byte
}

func NewClient(conn *websocket.Conn, hub HubInterface) *Client {
	return &Client{
		conn: conn,
		hub:  hub,
		send: make(chan []byte, 256),
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.(*Hub).unregister <- c
		c.conn.Close()
	}()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		msg, err := ParseMessage(data)
		if err != nil {
			continue // 잘못된 메시지는 무시
		}

		// 메시지를 JSON으로 다시 변환하여 broadcast
		encoded, _ := json.Marshal(msg)
		c.hub.(*Hub).broadcast <- encoded
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()

	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
