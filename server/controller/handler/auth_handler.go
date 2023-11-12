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

type AuthHandler struct {
	authInteractor usecase.IAuthInteractor
	userInteractor usecase.IUserInteractor
}

func NewAuthHandler(apiGroup *echo.Group, authInteractor usecase.IAuthInteractor, userInteractor usecase.IUserInteractor) {
	authHandler := &AuthHandler{
		authInteractor: authInteractor,
		userInteractor: userInteractor,
	}

	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", authHandler.Login)
	authGroup.GET("/callback", authHandler.CallBack)
}

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	// Get Github OAuth URL
	state := h.authInteractor.GenerateState(b)
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

	res, err := h.authInteractor.GetGithubUser(ctx, gitToken)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	// Get User
	githubId := res["login"].(string)
	githubName := res["name"].(string)

	user, err := h.userInteractor.GetUserByGithubID(ctx, githubId)
	if err != nil {
		if (err == sql.ErrNoRows) {
			githubImageURL := res["avatar_url"].(string)

			user, err = h.userInteractor.CreateUserAndGetByGithubID(ctx, githubId, githubName, githubImageURL)
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
