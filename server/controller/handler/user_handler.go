package handler

import (
	// "encoding/json"
	// "net/http"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
)

type IUserHandler interface {
	CreateUser(c echo.Context) error
}

type UserHandler struct {
	userInteractor usecase.IUserInteractor
}

func NewUserHandler(userInteractor usecase.IUserInteractor) IUserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	var user *entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.userInteractor.CreateUser(ctx, user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "User created")
}

