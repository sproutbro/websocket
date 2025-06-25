package ws

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// roomMap 은 채팅방 이름별로 클라이언트 ID와 연결을 저장합니다.
// 예: roomMap["room1"]["alice"] = conn
var (
	connMap = make(map[string]*websocket.Conn)
	roomMap = make(map[string]map[string]*websocket.Conn)
	mutex   = &sync.Mutex{}
)

// AddConn 은 새로 연결된 클라이언트를 목록에 추가합니다.
func AddConn(conn *websocket.Conn) {
	mutex.Lock()
	// connMap[conn] = true
	// mutex.Unlock()
}

// AddConnWithID 는 특정 ID와 함께 클라이언트를 등록합니다.
func AddConnWithID(id string, conn *websocket.Conn) {
	mutex.Lock()
	connMap[id] = conn
	mutex.Unlock()
}

// RemoveConn 은 끊어진 클라이언트를 목록에서 제거합니다.
func RemoveConn(conn *websocket.Conn) {
	mutex.Lock()
	// delete(connMap, conn)
	// mutex.Unlock()
}

func RemoveConnByID(id string) {
	mutex.Lock()
	delete(connMap, id)
	mutex.Unlock()
}

// CountClients 는 현재 연결된 클라이언트 수를 반환합니다.
func CountClients() int {
	mutex.Lock()
	defer mutex.Unlock()
	return len(connMap)
}

// GetAllClients 는 현재 연결된 모든 클라이언트를 반환합니다.
func GetAllClients() []string {
	mutex.Lock()
	defer mutex.Unlock()

	conns := make([]string, 0, len(connMap))
	for id, conn := range connMap {
		conns = append(conns, id)
		if conn == nil {
			log.Println("GetAllClients 오류 ")
		}
	}
	return conns
}

// Broadcast 는 등록된 모든 클라이언트에게 메시지를 보냅니다.
func Broadcast(message []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for id, conn := range connMap {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("메시지 전송 실패 (%v): %v", conn.RemoteAddr(), err)
			conn.Close()
			delete(connMap, id)
		}
	}
}

// SendTo 는 지정한 ID를 가진 클라이언트에게만 메시지를 보냅니다.
func SendTo(id string, message []byte) error {
	mutex.Lock()
	conn, ok := connMap[id]
	mutex.Unlock()

	if !ok {
		return fmt.Errorf("ID %s 에 해당하는 연결 없음", id)
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}

// AddConnWithRoom 은 특정 채팅방에 속한 클라이언트를 등록합니다.
func AddConnWithRoom(room, id string, conn *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := roomMap[room]; !exists {
		roomMap[room] = make(map[string]*websocket.Conn)
	}
	roomMap[room][id] = conn
}

// BroadcastToRoom 은 지정한 방에 속한 모든 사용자에게 메시지를 전송합니다.
func BroadcastToRoom(room string, message []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for id, conn := range roomMap[room] {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("[%s] 메시지 전송 실패: %v", id, err)
			conn.Close()
			delete(roomMap[room], id)
		}
	}
}
