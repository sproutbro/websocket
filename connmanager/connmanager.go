package connmanager

import (
	"log"
	"webgame/socket/global"

	"github.com/gorilla/websocket"
)

// 테스트 코드
// 인터페이스 갈아끼우기 연습
// 멥이랑 타입 익히기
var connMap = make(map[*websocket.Conn]bool)

func Register(conn *websocket.Conn) {
	global.Mutex.Lock()
	connMap[conn] = true
	global.Mutex.Unlock()
}

func Unregister(conn *websocket.Conn) {
	global.Mutex.Lock()
	delete(connMap, conn)
	global.Mutex.Unlock()
}

func Broadcast(message []byte) {
	global.Mutex.Lock()
	for conn := range connMap {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("connmanager Broadcast 에러")
			Unregister(conn)
		}
	}
	global.Mutex.Unlock()
}

func ClientsLen() int {
	return len(connMap)
}
