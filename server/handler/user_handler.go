package handler

import (
	// "encoding/json"
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IUserHandler interface {
	GetUser(c echo.Context) error
	GetUserByGithubID(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type UserHandler struct {
	userInteractor usecase.IUserInteractor
}

func NewUserHandler(userInteractor usecase.IUserInteractor) IUserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	return nil
}

func (h *UserHandler) GetUserByGithubID(c echo.Context) error {
	return nil
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	return nil
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	return nil
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	return nil
}
