package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 테스트 용도: 모든 도메인 허용
	},
}

// Handler 는 웹소켓 연결을 처리하고,
// 받은 메시지를 모든 연결된 클라이언트에게 Broadcast합니다.
func Handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id 쿼리 파라미터 필요", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("업그레이드 실패:", err)
		return
	}

	log.Printf("ID %s 연결됨 (%s)", id, conn.RemoteAddr())
	AddConnWithID(id, conn)

	go func() {
		defer func() {
			RemoveConnByID(id) // 👈 클라이언트 제거
			conn.Close()
		}()

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("읽기 실패:", err)
				break
			}

			log.Printf("클라이언트로부터 받은 메시지: %s", msg)

			// 받은 메시지를 그대로 다시 보냄 (echo)
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				log.Println("쓰기 실패:", err)
				break
			}

			// 모든 연결된 클라이언트에게 메시지 전송
			Broadcast(msg)
		}
	}()
}
