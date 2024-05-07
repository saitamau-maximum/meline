package entity

import (
	"github.com/gorilla/websocket"
)

type NotifyClient struct {
	Ws     *websocket.Conn
	SendCh chan []byte
	UserID uint64
}

func NewNotifyClientEntity(ws *websocket.Conn, userID uint64) *NotifyClient {
	return &NotifyClient{
		Ws:     ws,
		SendCh: make(chan []byte),
		UserID: userID,
	}
}
