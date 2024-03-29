package entity

// チャットルームごとのクライアントの管理を行う
type Hub struct {
	Clients      map[uint64]map[*Client]bool
	RegisterCh   chan *Client
	UnregisterCh chan *Client
	BroadcastCh  chan []byte
	ChannelIDCh  chan uint64
}

func NewHubEntity() *Hub {
	return &Hub{
		Clients:      make(map[uint64]map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnregisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
		ChannelIDCh:  make(chan uint64),
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.RegisterClient(client, client.ChannelID)
		case client := <-h.UnregisterCh:
			h.UnregisterClient(client, client.ChannelID)
		case message := <-h.BroadcastCh:
			h.BroadcastMessage(message)
		}
	}
}

func (h *Hub) RegisterClient(client *Client, channelID uint64) {
	if _, ok := h.Clients[channelID]; !ok {
		h.Clients[channelID] = make(map[*Client]bool)
	}

	h.Clients[channelID][client] = true
}

func (h *Hub) UnregisterClient(client *Client, channelID uint64) {
	if _, ok := h.Clients[channelID][client]; ok {
		delete(h.Clients[channelID], client)
		close(client.SendCh)
	}
}

func (h *Hub) BroadcastMessage(message []byte) {
	channelID := <-h.ChannelIDCh

	for client := range h.Clients[channelID] {
		client.SendCh <- message
	}
}
