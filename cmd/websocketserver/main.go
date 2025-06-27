package main

import (
	"log"
	ws "myonly/internal/websocket"
	"net/http"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", ws.ServeWS(hub))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("✅ WebSocket 서버 시작: ws://localhost:8080/ws")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("서버 종료: ", err)
	}
}
