package entity

import (
	"github.com/gorilla/websocket"
)

type NotifyClient struct {
	Ws               *websocket.Conn
	SendCh           chan []byte
	JoinedChannelIDs map[uint64]*NotifyClientJoinedChannel
	UserID           uint64
}

func NewNotifyClientEntity(ws *websocket.Conn, userID uint64, joinedChannelIDs map[uint64]*NotifyClientJoinedChannel) *NotifyClient {
	return &NotifyClient{
		Ws:               ws,
		SendCh:           make(chan []byte),
		JoinedChannelIDs: joinedChannelIDs,
		UserID:           userID,
	}
}
