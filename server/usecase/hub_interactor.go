package usecase

import "github.com/saitamau-maximum/meline/domain/entity"

type IHubInteractor interface {
	RunLoop()
	RegisterClient(client *entity.Client ,channelID uint64)
	UnregisterClient(client *entity.Client ,channelID uint64)
	BroadcastMessage(message []byte)
}

type HubInteractor struct {
	Hub *entity.Hub
}

func NewHubInteractor(h *entity.Hub) *HubInteractor {
	return &HubInteractor{
		Hub: h,
	}
}

func (h *HubInteractor) RunLoop() {
	for {
		select {
		case client := <-h.Hub.RegisterCh:
			h.RegisterClient(client, client.ChannelID)
		case client := <-h.Hub.UnregisterCh:
			h.UnregisterClient(client, client.ChannelID)
		case message := <-h.Hub.BroadcastCh:
			h.BroadcastMessage(message)
		}
	}
}

func (h *HubInteractor) RegisterClient(client *entity.Client, channelID uint64) {
	h.Hub.Clients[channelID][client] = true
}

func (h *HubInteractor) UnregisterClient(client *entity.Client, channelID uint64) {
	if _, ok := h.Hub.Clients[channelID][client]; ok {
		delete(h.Hub.Clients[channelID], client)
		close(client.SendCh)
	}
}

func (h *HubInteractor) BroadcastMessage(message []byte) {
	channelID := <-h.Hub.ChannelIDCh

	for client := range h.Hub.Clients[channelID] {
		client.SendCh <- message
	}
}
