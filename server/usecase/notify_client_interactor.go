package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type INotifyClientInteractor interface {
	ReadPump(ctx context.Context, client *entity.NotifyClient, hub *entity.Hub) error
	WritePump(ctx context.Context, client *entity.NotifyClient) error
	CreateNotifyClient(ctx context.Context, ws *websocket.Conn, userID uint64) (*entity.NotifyClient, error)
}

type NotifyClientInteractor struct {
	ChannelRepository repository.IChannelRepository
}

func NewNotifyClientInteractor(channelRepository repository.IChannelRepository) *NotifyClientInteractor {
	return &NotifyClientInteractor{
		ChannelRepository: channelRepository,
	}
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

func (c *NotifyClientInteractor) WritePump(ctx context.Context, client *entity.NotifyClient) error {
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

func (c *NotifyClientInteractor) CreateNotifyClient(ctx context.Context, ws *websocket.Conn, userID uint64) (*entity.NotifyClient, error) {
	// NOTE: 通知クライアントが参加しているチャンネルを取得
	joinedChannelIDs, err := c.ChannelRepository.FetchJoinedChannelIDs(ctx, userID)
	if err != nil {
		return nil, err
	}

	joinedChannels := make(map[uint64]bool)
	for _, id := range joinedChannelIDs {
		joinedChannels[id] = true
	}

	return entity.NewNotifyClientEntity(ws, userID, joinedChannels), nil
}
