package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IAuthHandler interface {
	Login(c echo.Context) error
	Callback(c echo.Context) error
}

type AuthHandler struct {
	interactor usecase.IAuthInteractor
}

func NewAuthHandler(interactor usecase.IAuthInteractor) IAuthHandler {
	return &AuthHandler{
		interactor: interactor,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	state := c.Request().URL.Query().Get("state")

	redirectUrl := h.interactor.GetGithubOAuthURL(c.Request().Context(), state)

	return c.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func (h *AuthHandler) Callback(c echo.Context) error {
	query := c.Request().URL.Query()
	code := query.Get("code")

	accessToken, err := h.interactor.GetGithubOAuthToken(context.Background(), code)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name: "access_token",
		Value: accessToken,
	}

	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/logged_in")
}
