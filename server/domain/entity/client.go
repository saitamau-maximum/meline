package entity

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Ws        *websocket.Conn
	SendCh    chan []byte
	ChannelID uint64
}

func NewClientEntity(ws *websocket.Conn, channelID uint64) *Client {
	return &Client{
		Ws:        ws,
		SendCh:    make(chan []byte),
		ChannelID: channelID,
	}
}
