package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 업그레이드 도구: HTTP 요청 → WebSocket 연결로 변경
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // CORS 테스트용. 실무에선 꼭 검증해야 함!
	},
}

// 웹소켓 요청을 처리하는 핸들러
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// HTTP → WebSocket 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 업그레이드 실패:", err)
		return
	}
	defer conn.Close()

	log.Println("클라이언트 접속됨")

	for {
		// 메시지 수신
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("메시지 읽기 오류:", err)
			break
		}

		log.Printf("받은 메시지: %s", message)

		// Echo: 받은 메시지를 그대로 다시 보냄
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("메시지 쓰기 오류:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", echoHandler)

	port := 8081
	fmt.Printf("WebSocket 서버 시작: ws://localhost:%d/ws\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
