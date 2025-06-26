package main

import (
	"myonly/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	/*
		/ws 경로에
		미들웨어 SocketHandler 함수 /middleware/handler.go
		아직구현 안함
	*/
	mux.HandleFunc("/ws", middleware.SocketHandler())

	/*
		HTTP 요청에 응답
		middleware/handler 의 HttpHandler 함수에서
		미들웨어에서 CORS → Auth → router
		라우터의 factory 에서 경로 등록
	*/
	mux.HandleFunc("/", middleware.HttpHandler())

	http.ListenAndServe(":8080", mux)
}
