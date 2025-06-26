package main

import ws26socket "myonly/ws26/socket"

// main 함수는 웹서버를 시작하고 /ws 엔드포인트에 웹소켓 핸들러를 연결합니다.
func main() {
	ws26socket.Ws26socketMain()
}

/*
func ws() {
	mux := http.NewServeMux()

	// 웹소켓 핸들러 등록
	mux.HandleFunc("/ws", ws.Handler)

	log.Println("서버 시작: http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("서버 실패:", err)
	}
}
*/
