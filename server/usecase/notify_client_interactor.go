package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saitamau-maximum/meline/domain/entity"
)

type INotifyClientInteractor interface {
	ReadPump(ctx context.Context, client *entity.NotifyClient, hub *entity.Hub) error
	WritePump(ctx context.Context, client *entity.NotifyClient, hub *entity.Hub) error
}

type NotifyClientInteractor struct{}

func NewNotifyClientInteractor() *NotifyClientInteractor {
	return &NotifyClientInteractor{}
}

func (c *NotifyClientInteractor) ReadPump(ctx context.Context, client *entity.NotifyClient, hub *entity.Hub) error {
	defer func() {
		hub.UnregisterNotifyCh <- client
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

func (c *NotifyClientInteractor) WritePump(ctx context.Context, client *entity.NotifyClient, hub *entity.Hub) error {
	ticker := time.NewTicker(pingWait)

	defer func() {
		client.Ws.Close()
		ticker.Stop()
	}()

	for {
		select {
		case message, ok := <-client.SendCh:
			client.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				client.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				return nil
			}

			if err := client.Ws.WriteMessage(websocket.TextMessage, message); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			client.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		}
	}
}
