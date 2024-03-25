package entity

type Hub struct {
	Clients map[*Client]bool
	RegisterCh chan *Client
	UnregisterCh chan *Client
	BroadcastCh chan []byte
}

func NewHubEntity() *Hub {
	return &Hub{
		Clients: make(map[*Client]bool),
		RegisterCh: make(chan *Client),
		UnregisterCh: make(chan *Client),
		BroadcastCh: make(chan []byte),
	}
}
