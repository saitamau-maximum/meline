package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type UserHandler struct {
	userInteractor usecase.IUserInteractor
}

func NewUserHandler(userGroup *echo.Group, userInteractor usecase.IUserInteractor) {
	userHandler := &UserHandler{
		userInteractor: userInteractor,
	}

	userGroup.GET("/me", userHandler.Me)
}

func (h *UserHandler) Me(c echo.Context) error {
	userId := c.Get("user_id").(uint64)
	userResponse, err := h.userInteractor.GetUserByID(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, userResponse)
}
