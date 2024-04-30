package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saitamau-maximum/meline/domain/entity"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingWait       = (pongWait * 7) / 10
	maxMessageSize = 512
)

type IClientInteractor interface {
	ReadLoop(ctx context.Context, client *entity.Client, hub *entity.Hub) error
	WriteLoop(ctx context.Context, client *entity.Client, hub *entity.Hub) error
}

type ClientInteractor struct{}

func NewClientInteractor() *ClientInteractor {
	return &ClientInteractor{}
}

func (c *ClientInteractor) ReadLoop(ctx context.Context, client *entity.Client, hub *entity.Hub) error {
	defer func() {
		hub.UnregisterCh <- client
		client.Ws.Close()
	}()

	client.Ws.SetReadLimit(maxMessageSize)
	client.Ws.SetReadDeadline(time.Now().Add(pongWait))
	client.Ws.SetPongHandler(func(string) error {
		client.Ws.SetReadDeadline(time.Now().Add(pongWait))
		log.Println("pong")
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, _, err := client.Ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("websocket: %v", err)
				}

				return nil
			}
		}
	}
}

func (c *ClientInteractor) WriteLoop(ctx context.Context, client *entity.Client, hub *entity.Hub) error {
	ticker := time.NewTicker(pingWait)
	defer func() {
		ticker.Stop()
		client.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-client.SendCh:
			client.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				return nil
			}

			w, err := client.Ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return fmt.Errorf("websocket: %w", err)
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		case <-ticker.C:
			client.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			log.Println("ping")
			if err := client.Ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		}
	}
}
