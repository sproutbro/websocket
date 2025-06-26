package ws26socket

import "sync"

type Hub struct {
	Rooms map[string]*Room
	mutex sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) GetOrCreateRoom(id string) *Room {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	room, exists := h.Rooms[id]
	if !exists {
		room = NewRoom(id)
		h.Rooms[id] = room
	}
	return room
}
