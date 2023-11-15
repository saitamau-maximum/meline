package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	b = 32
)

type OAuthHandler struct {
	githubOAuthInteractor usecase.IGithubOAuthInteractor
	userInteractor usecase.IUserInteractor
}

func NewOAuthHandler(authGroup *echo.Group, githubOAuthInteractor usecase.IGithubOAuthInteractor, userInteractor usecase.IUserInteractor) {
	oAuthHandler := &OAuthHandler{
		githubOAuthInteractor: githubOAuthInteractor,
		userInteractor: userInteractor,
	}

	authGroup.GET("/login", oAuthHandler.Login)
	authGroup.GET("/callback", oAuthHandler.CallBack)
}

func (h *OAuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	// Get Github OAuth URL
	state := h.githubOAuthInteractor.GenerateState(b)
	url := h.githubOAuthInteractor.GetGithubOAuthURL(ctx, state)

	return c.Redirect(http.StatusMovedPermanently, url)
}

func (h *OAuthHandler) CallBack(c echo.Context) error {
	ctx := context.Background()

	code := c.QueryParam("code")
	gitToken, err := h.githubOAuthInteractor.GetGithubOAuthToken(ctx, code)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	userRes, err := h.githubOAuthInteractor.GetGithubUser(ctx, gitToken)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.userInteractor.GetUserByGithubID(ctx, userRes.OAuthUserID)
	if err != nil {
		if (err == sql.ErrNoRows) {
			user, err = h.userInteractor.CreateUser(ctx, userRes.OAuthUserID, userRes.Name, userRes.ImageURL)
			if err != nil {
				log.Default().Println(err)
				return c.JSON(http.StatusInternalServerError, err)
			}
		} else {
			log.Default().Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Set Access Token
	token, err := h.githubOAuthInteractor.CreateAccessToken(ctx, user)
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
