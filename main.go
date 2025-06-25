package main

import (
	"log"
	"net/http"
	"webgame/socket/ws"
)

func main() {
	mux := http.NewServeMux()

	// 웹소켓 엔드포인트 등록
	mux.HandleFunc("/ws", ws.Handler)

	log.Println("WebSocket 서버 실행 중: ws://localhost:8080/ws")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("서버 실패:", err)
	}
}
