package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/adapter/request"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
	"golang.org/x/sync/errgroup"
)

type MessageHandler struct {
	messageInteractor usecase.IMessageInteractor
	hubInteractor     usecase.IHubInteractor
	clientInteractor  usecase.IClientInteractor
	hub               *entity.Hub
}

func NewMessageHandler(messageGroup *echo.Group, messageInteractor usecase.IMessageInteractor, hubInteractor usecase.IHubInteractor, clientInteractor usecase.IClientInteractor, hub *entity.Hub) {
	messageHandler := &MessageHandler{
		messageInteractor: messageInteractor,
		hubInteractor:     hubInteractor,
		clientInteractor:  clientInteractor,
		hub:               hub,
	}

	messageGroup.GET("", messageHandler.GetByChannelID)
	messageGroup.POST("", messageHandler.Create)
	messageGroup.PUT("/:id", messageHandler.Update)
	messageGroup.POST("/:id/reply", messageHandler.CreateReply)
	messageGroup.DELETE("/:id", messageHandler.Delete)
	messageGroup.GET("/ws/:id", messageHandler.WebSocket)
}

func (h *MessageHandler) GetByChannelID(c echo.Context) error {
	channelId := c.Param("channel_id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	messages, err := h.messageInteractor.GetMessagesByChannelID(c.Request().Context(), channelIdUint64)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) Create(c echo.Context) error {
	channelId := c.Param("channel_id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	createMessageRequest := &request.CreateMessageRequest{}
	if err := c.Bind(createMessageRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := createMessageRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	userId := c.Get("user_id").(uint64)

	if err := h.messageInteractor.Create(c.Request().Context(), userId, channelIdUint64, createMessageRequest.Content); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *MessageHandler) CreateReply(c echo.Context) error {
	channelId := c.Param("channel_id")
	replyToId := c.Param("id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	createMessageRequest := &request.CreateMessageRequest{}
	if err := c.Bind(createMessageRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := createMessageRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	userId := c.Get("user_id").(uint64)

	if err := h.messageInteractor.CreateReply(c.Request().Context(), userId, channelIdUint64, replyToId, createMessageRequest.Content); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *MessageHandler) Update(c echo.Context) error {
	id := c.Param("id")

	updateMessageRequest := &request.UpdateMessageRequest{}
	if err := c.Bind(updateMessageRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := updateMessageRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if err := h.messageInteractor.Update(c.Request().Context(), id, updateMessageRequest.Content); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *MessageHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.messageInteractor.Delete(c.Request().Context(), id); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *MessageHandler) WebSocket(c echo.Context) error {
	channelId := c.Param("id")

	channelIdUint64, err := strconv.ParseUint(channelId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	ws := websocket.Upgrader{}

	conn, err := ws.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	client := entity.NewClientEntity(conn, channelIdUint64)

	h.hubInteractor.RegisterClient(client, channelIdUint64)

	var eg errgroup.Group

	eg.Go(func() error {
		return h.clientInteractor.ReadLoop(c.Request().Context(), h.hub.BroadcastCh, h.hub.ChannelIDCh, client)
	})
	eg.Go(func() error {
		return h.clientInteractor.WriteLoop(c.Request().Context(), client)
	})

	if err := eg.Wait(); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
