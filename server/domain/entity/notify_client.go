package entity

import (
	"github.com/gorilla/websocket"
)

type NotifyClient struct {
	Ws               *websocket.Conn
	SendCh           chan []byte
	JoinedChannelMap map[uint64]bool
	UserID           uint64
}

func NewNotifyClientEntity(ws *websocket.Conn, userID uint64, joinedChannelMap map[uint64]bool) *NotifyClient {
	return &NotifyClient{
		Ws:               ws,
		SendCh:           make(chan []byte),
		JoinedChannelMap: joinedChannelMap,
		UserID:           userID,
	}
}

func (c *NotifyClient) IsJoinedChannel(channelID uint64) bool {
	_, ok := c.JoinedChannelMap[channelID]
	return ok
}
