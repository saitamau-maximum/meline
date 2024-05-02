package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/adapter/request"
	"github.com/saitamau-maximum/meline/usecase"
)

type ChannelHandler struct {
	channelInteractor usecase.IChannelInteractor
}

func NewChannelHandler(channelGroup *echo.Group, channelInteractor usecase.IChannelInteractor) {
	channelHandler := &ChannelHandler{
		channelInteractor: channelInteractor,
	}

	channelGroup.GET("", channelHandler.GetAllChannels)
	channelGroup.GET("/:id", channelHandler.GetChannelByID)
	channelGroup.POST("/:id/join", channelHandler.JoinChannel)
	channelGroup.POST("", channelHandler.CreateChannel)
	channelGroup.POST("/:id/create", channelHandler.CreateChildChannel)
	channelGroup.PUT("/:id", channelHandler.UpdateChannel)
	channelGroup.DELETE("/:id", channelHandler.DeleteChannel)
	channelGroup.DELETE("/:id/leave", channelHandler.LeaveChannel)
}

func (h *ChannelHandler) GetAllChannels(c echo.Context) error {
	userId := c.Get("user_id").(uint64)
	channelsResponse, err := h.channelInteractor.GetAllChannels(c.Request().Context(), userId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, channelsResponse)
}

func (h *ChannelHandler) GetChannelByID(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	channelResponse, err := h.channelInteractor.GetChannelByID(c.Request().Context(), channelId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, channelResponse)
}

func (h *ChannelHandler) CreateChannel(c echo.Context) error {
	userId := c.Get("user_id").(uint64)

	req := request.CreateChannelRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := h.channelInteractor.CreateChannel(c.Request().Context(), req.Name, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *ChannelHandler) CreateChildChannel(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userId := c.Get("user_id").(uint64)

	req := request.CreateChannelRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := h.channelInteractor.CreateChildChannel(c.Request().Context(), req.Name, channelId, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *ChannelHandler) UpdateChannel(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	req := request.UpdateChannelRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := h.channelInteractor.UpdateChannel(c.Request().Context(), channelId, req.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *ChannelHandler) DeleteChannel(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := h.channelInteractor.DeleteChannel(c.Request().Context(), channelId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *ChannelHandler) JoinChannel(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userId := c.Get("user_id").(uint64)

	if err := h.channelInteractor.JoinChannel(c.Request().Context(), channelId, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *ChannelHandler) LeaveChannel(c echo.Context) error {
	id := c.Param("id")

	channelId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userId := c.Get("user_id").(uint64)

	if err := h.channelInteractor.LeaveChannel(c.Request().Context(), channelId, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
