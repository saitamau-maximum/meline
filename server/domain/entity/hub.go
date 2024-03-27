package entity

// 複数のチャットルームに対するクライアントの管理を行う
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
