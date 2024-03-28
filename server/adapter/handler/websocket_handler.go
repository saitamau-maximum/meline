package handler

import (
	"log"
	"net/http"
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
	channelId := c.Param("channel_id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	client := entity.NewClientEntity(conn, channelIdUint64)

	h.hub.RegisterClient(client, channelIdUint64)

	var eg errgroup.Group

	eg.Go(func() error {
		return h.clientInteractor.WriteLoop(c.Request().Context(), client)
	})

	if err := eg.Wait(); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
