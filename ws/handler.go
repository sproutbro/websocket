package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 테스트 용도: 모든 도메인 허용
	},
}

// Handler 는 웹소켓 연결을 수락하고, 클라이언트가 보낸 JSON 메시지를
// 파싱하여 특정 사용자에게 전송합니다.
func Handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	room := r.URL.Query().Get("room")

	if id == "" || room == "" {
		http.Error(w, "id 쿼리 파라미터 필요", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("업그레이드 실패:", err)
		return
	}

	log.Printf("ID %s 연결됨 (Room: %s)", id, room)
	// conn.RemoteAddr()
	// AddConnWithID(id, conn)
	AddConnWithRoom(room, id, conn)

	go func() {
		defer func() {
			RemoveConnByID(id)
			conn.Close()
		}()

		for {
			msgType, raw, err := conn.ReadMessage()
			if err != nil {
				log.Println("읽기 실패:", err)
				break
			}

			log.Printf("[%s] - [%s] %s", id, room, raw)
			BroadcastToRoom(room, raw)

			var msg Message
			if err := json.Unmarshal(raw, &msg); err != nil {
				log.Println("메시지 파싱 실패:", err)
				continue
			}

			log.Printf("[%s] - [%s]: %s", id, msg.To, msg.Msg)

			// 지정한 사용자에게 메시지 전송
			if err := SendTo(msg.To, []byte(msg.Msg)); err != nil {
				log.Println("전송 실패:", err)
			}

			// 모든 연결된 클라이언트에게 메시지 전송
			Broadcast(raw)

			// 본인에게 다시
			err = conn.WriteMessage(msgType, raw)
			if err != nil {
				log.Println("쓰기 실패:", err)
				break
			}
		}
	}()
}
