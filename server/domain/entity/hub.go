package entity

// チャットルームごとのクライアントの管理を行う
type Hub struct {
	Clients            map[uint64]map[*Client]bool
	NotifyClients      map[uint64]map[*NotifyClient]bool
	RegisterCh         chan *Client
	RegisterNotifyCh   chan *NotifyClient
	UnregisterCh       chan *Client
	UnregisterNotifyCh chan *NotifyClient
	BroadcastCh        chan *BroadcastCh
	NotifyBroadcastCh  chan *NotifyBroadcastCh
}

func NewHubEntity() *Hub {
	return &Hub{
		Clients:            make(map[uint64]map[*Client]bool),
		NotifyClients:      make(map[uint64]map[*NotifyClient]bool),
		RegisterCh:         make(chan *Client),
		RegisterNotifyCh:   make(chan *NotifyClient),
		UnregisterCh:       make(chan *Client),
		UnregisterNotifyCh: make(chan *NotifyClient),
		BroadcastCh:        make(chan *BroadcastCh),
		NotifyBroadcastCh:  make(chan *NotifyBroadcastCh),
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.RegisterClient(client, client.ChannelID)
		case client := <-h.RegisterNotifyCh:
			h.RegisterNotifyClient(client, client.UserID)
		case client := <-h.UnregisterCh:
			h.UnregisterClient(client, client.ChannelID)
		case client := <-h.UnregisterNotifyCh:
			h.UnregisterNotifyClient(client, client.UserID)
		case message := <-h.BroadcastCh:
			h.BroadcastMessage(message.Message, message.ChannelID)
		case message := <-h.NotifyBroadcastCh:
			h.NotifyBroadcastMessage(message.Message, message.UserIDs)
		}
	}
}

func (h *Hub) RegisterClient(client *Client, channelID uint64) {
	if _, ok := h.Clients[channelID]; !ok {
		h.Clients[channelID] = make(map[*Client]bool)
	}

	h.Clients[channelID][client] = true
}

func (h *Hub) RegisterNotifyClient(client *NotifyClient, userID uint64) {
	if _, ok := h.NotifyClients[userID]; !ok {
		h.NotifyClients[userID] = make(map[*NotifyClient]bool)
	}

	h.NotifyClients[userID][client] = true
}

func (h *Hub) UnregisterClient(client *Client, channelID uint64) {
	if _, ok := h.Clients[channelID]; !ok {
		return
	}

	if _, ok := h.Clients[channelID][client]; ok {
		delete(h.Clients[channelID], client)
		close(client.SendCh)
	}
}

func (h *Hub) UnregisterNotifyClient(client *NotifyClient, userID uint64) {
	if _, ok := h.NotifyClients[userID]; !ok {
		return
	}

	if _, ok := h.NotifyClients[userID][client]; ok {
		delete(h.NotifyClients[userID], client)
		close(client.SendCh)
	}
}

func (h *Hub) BroadcastMessage(message []byte, channelID uint64) {
	for client := range h.Clients[channelID] {
		select {
		case client.SendCh <- message:
		default:
			delete(h.Clients[channelID], client)
			close(client.SendCh)
		}
	}
}

func (h *Hub) NotifyBroadcastMessage(message []byte, userIDs []uint64) {
	for userID := range h.NotifyClients {
		if _, ok := h.NotifyClients[userID]; !ok {
			continue
		}

		for client := range h.NotifyClients[userID] {
			select {
			case client.SendCh <- message:
			default:
				delete(h.NotifyClients[userID], client)
				close(client.SendCh)
			}
		}
	}
}
