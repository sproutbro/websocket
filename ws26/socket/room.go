package ws26socket

type Room struct {
	ID      string
	Clients map[string]*Client
}

func NewRoom(id string) *Room {
	return &Room{
		ID:      id,
		Clients: make(map[string]*Client),
	}
}

func (r *Room) AddClient(c *Client) {
	r.Clients[c.ID] = c
}

func (r *Room) RemoveClient(id string) {
	delete(r.Clients, id)
}
