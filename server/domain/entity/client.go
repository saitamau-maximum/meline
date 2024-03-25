package entity

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Ws *websocket.Conn
	SendCh chan []byte
}

func NewClientEntity(ws *websocket.Conn) *Client {
	return &Client{
		Ws: ws,
		SendCh: make(chan []byte),
	}
}
