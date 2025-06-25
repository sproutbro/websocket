package ws

import (
	"log"
	"net/http"
	"webgame/socket/connmanager"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 모든 도메인 허용 (개발용)
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("업그레이드 실패:", err)
		return
	}

	log.Println("새 연결 수락됨:", conn.RemoteAddr())

	connmanager.Register(conn)

	for {

		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("읽기 오류:", err)
		}

		connmanager.Broadcast(msg)

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("쓰기 오류:", err)
			break
		}

	}

	connmanager.Unregister(conn)
}
