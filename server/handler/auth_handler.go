package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IAuthHandler interface {
	Callback(c echo.Context) error
	Auth(c echo.Context) error
}

type AuthHandler struct {
	interactor usecase.IAuthInteractor
}

func NewAuthHandler(interactor usecase.IAuthInteractor) IAuthHandler {
	return &AuthHandler{
		interactor: interactor,
	}
}

func (h *AuthHandler) Callback(c echo.Context) error {
	query := c.Request().URL.Query()
	code := query.Get("code")

	accessToken, err := h.interactor.GetGithubOAuthToken(context.Background(), code)
	if err != nil {
		return err
	}

	return c.JSON(200, accessToken)
}

func (h *AuthHandler) Auth(c echo.Context) error {
	return nil
}
