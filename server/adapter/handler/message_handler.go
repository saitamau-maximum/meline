package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/adapter/request"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
)

type MessageHandler struct {
	messageInteractor usecase.IMessageInteractor
	hub               *entity.Hub
}

func NewMessageHandler(messageGroup *echo.Group, messageInteractor usecase.IMessageInteractor, hub *entity.Hub) {
	messageHandler := &MessageHandler{
		messageInteractor: messageInteractor,
		hub:               hub,
	}

	messageGroup.GET("", messageHandler.GetByChannelID)
	messageGroup.POST("", messageHandler.Create)
	messageGroup.POST("/:id/reply", messageHandler.CreateReply)
	messageGroup.PUT("/:id", messageHandler.Update)
	messageGroup.DELETE("/:id", messageHandler.Delete)
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

	res, err := h.messageInteractor.Create(c.Request().Context(), userId, channelIdUint64, createMessageRequest.Content)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	h.hub.BroadcastCh <- entity.NewBroadcastChEntity(jsonRes, channelIdUint64)

	return c.JSON(http.StatusCreated, res)
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

	res, err := h.messageInteractor.CreateReply(c.Request().Context(), userId, channelIdUint64, replyToId, createMessageRequest.Content)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	h.hub.BroadcastCh <- entity.NewBroadcastChEntity(jsonRes, channelIdUint64)

	return c.JSON(http.StatusCreated, res)
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
