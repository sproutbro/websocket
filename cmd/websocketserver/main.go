package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := 8081
	fmt.Printf("WebSocket 서버 시작: ws://localhost:%d/ws\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
