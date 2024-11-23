package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saitamau-maximum/meline/domain/entity"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingWait       = (pongWait * 7) / 10
	maxMessageSize = 8192
)

type IMessageClientInteractor interface {
	ReadPump(ctx context.Context, client *entity.Client, hub *entity.Hub) error
	WritePump(ctx context.Context, client *entity.Client) error
}

type MessageClientInteractor struct{}

func NewMessageClientInteractor() *MessageClientInteractor {
	return &MessageClientInteractor{}
}

func (c *MessageClientInteractor) ReadPump(ctx context.Context, client *entity.Client, hub *entity.Hub) error {
	defer func() {
		hub.UnregisterCh <- client
		client.Ws.Close()
	}()

	client.Ws.SetReadLimit(maxMessageSize)
	client.Ws.SetReadDeadline(time.Now().Add(pongWait))
	client.Ws.SetPongHandler(func(string) error {
		client.Ws.SetReadDeadline(time.Now().Add(pongWait))
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
					return fmt.Errorf("websocket: %w", err)
				}

				return nil
			}
		}
	}
}

func (c *MessageClientInteractor) WritePump(ctx context.Context, client *entity.Client) error {
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
			if err := client.Ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		}
	}
}
