package ws26socket

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var hub = NewHub()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 업그레이드 실패 : ", err)
		return
	}

	id := generateRandomID()
	roomID := "lobby"
	client := NewClient(id, conn, roomID)
	room := hub.GetOrCreateRoom(roomID)
	room.AddClient(client)

	go handleClient(client, room)
}

func handleClient(c *Client, r *Room) {
	defer func() {
		r.RemoveClient(c.ID)
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("메시지 읽기 실패 : ", err)
			break
		}

		for id, client := range r.Clients {
			if id == c.ID {
				continue
			}
			err = client.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("전송 실패 : ", err)
			}
		}
	}
}

func generateRandomID() string {
	b := make([]byte, 4) // 4바이트 = 8글자
	rand.Read(b)
	return hex.EncodeToString(b)
}

func Ws26socketMain() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("서버시작: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
