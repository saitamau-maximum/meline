package handler

import (
	"context"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
	"golang.org/x/sync/errgroup"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketHandler struct {
	clientInteractor usecase.IClientInteractor
	hub              *entity.Hub
}

func NewWebSocketHandler(websocketGroup *echo.Group, clientInteractor usecase.IClientInteractor, hub *entity.Hub) {
	webSocketHandler := &WebSocketHandler{
		clientInteractor: clientInteractor,
		hub:              hub,
	}

	websocketGroup.GET("/:channel_id", webSocketHandler.WebSocket)
}

func (h *WebSocketHandler) WebSocket(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())

	channelId := c.Param("channel_id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		cancel()
		return err
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		cancel()
		return err
	}

	client := entity.NewClientEntity(conn, channelIdUint64)

	h.hub.RegisterClient(client, channelIdUint64)

	var eg errgroup.Group

	eg.Go(func() error {
		return h.clientInteractor.ReadPump(ctx, client, h.hub)
	})
	eg.Go(func() error {
		return h.clientInteractor.WritePump(ctx, client, h.hub)
	})

	if err := eg.Wait(); err != nil {
		cancel()
		return err
	}

	cancel()

	return nil
}
