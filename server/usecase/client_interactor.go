package usecase

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/saitamau-maximum/meline/domain/entity"
)

type IClientInteractor interface {
	Disconnect(client *entity.Client)
	ReadLoop(ctx context.Context, broadCast chan<- []byte, channelId chan<- uint64, client *entity.Client) error
	WriteLoop(ctx context.Context, client *entity.Client) error
}

type ClientInteractor struct{}

func NewClientInteractor() *ClientInteractor {
	return &ClientInteractor{}
}

func (c *ClientInteractor) Disconnect(client *entity.Client) {
	client.Ws.Close()
}

func (c *ClientInteractor) ReadLoop(ctx context.Context, broadCast chan<- []byte, channelId chan<- uint64, client *entity.Client) error {
	defer func() {
		c.Disconnect(client)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			{
				_, message, err := client.Ws.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						return fmt.Errorf("websocket: %w", err)
					}
				}

				broadCast <- message
				channelId <- client.ChannelID
			}
		}
	}
}

func (c *ClientInteractor) WriteLoop(ctx context.Context, client *entity.Client) error {
	defer func() {
		c.Disconnect(client)
	}()

	for {
		select {
		case message, ok := <-client.SendCh:
			if !ok {
				return fmt.Errorf("send channel closed")
			}

			w, err := client.Ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return fmt.Errorf("websocket: %w", err)
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return fmt.Errorf("websocket: %w", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
