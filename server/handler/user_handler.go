package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IUserHandler interface {
	Login(ctx echo.Context) error
}

type UserHandler struct {
	userInteractor usecase.IUserInteractor	
	authInteractor usecase.IAuthInteractor
}

func NewUserHandler(userInteractor usecase.IUserInteractor, authInteractor usecase.IAuthInteractor) IUserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
		authInteractor: authInteractor,
	}
}

func (h *UserHandler) Login(ctx echo.Context) error {
	return nil
}
