package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/saitamau-maximum/meline/utils"
)

const (
	b = 32
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
	state := utils.SecureRandomStr(b)
	url := h.authInteractor.GetGithubOAuthURL(ctx, state)

	return c.Redirect(http.StatusMovedPermanently, url)
}

func (h *AuthHandler) CallBack(c echo.Context) error {
	ctx := context.Background()

	code := c.QueryParam("code")
	gitToken, err := h.authInteractor.GetGithubOAuthToken(ctx, code)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	resByte, err := h.authInteractor.GetGithubUser(ctx, gitToken)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resByte, &res); err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Get User
	githubId := res["login"].(string)

	user, err := h.userInteractor.GetUserByGithubID(ctx, githubId)

	if err != nil {
		if (err == sql.ErrNoRows) {
			return c.Redirect(http.StatusMovedPermanently, "/signup")
		}
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Set Access Token
	token, err := h.authInteractor.CreateAccessToken(ctx, user)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "success")
}
