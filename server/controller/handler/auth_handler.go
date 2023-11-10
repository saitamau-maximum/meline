package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IAuthHandler interface {
	Login(c echo.Context) error
	CallBack(c echo.Context) error
}

type AuthHandler struct {
	authInteractor usecase.IAuthInteractor
	userInteractor usecase.IUserInteractor
}

func NewAuthHandler( authInteractor usecase.IAuthInteractor, userInteractor usecase.IUserInteractor) IAuthHandler {
	return &AuthHandler{
		authInteractor: authInteractor,
		userInteractor: userInteractor,
	}
}



func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	// Get Github OAuth URL
	state := "state"
	url := h.authInteractor.GetGithubOAuthURL(ctx, state)

	return c.Redirect(http.StatusMovedPermanently, url)
}

func (h *AuthHandler) CallBack(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.QueryParam("code")
	gitToken, err := h.authInteractor.GetGithubOAuthToken(ctx, code)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	resByte, err := h.authInteractor.GetGithubUser(ctx, gitToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resByte, &res); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Get User
	githubId := res["login"].(string)

	user, err := h.userInteractor.GetUserByGithubID(ctx, githubId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if user == nil {
		return c.Redirect(http.StatusMovedPermanently, "/signup")
	}

	// Set Access Token
	token, err := h.authInteractor.CreateAccessToken(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookie := &http.Cookie{
		Name: "access_token",
		Value: token,
	}

	c.SetCookie(cookie)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
