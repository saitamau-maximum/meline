package handler

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	STATE_LENGTH = 32
)

var (
	isDev = os.Getenv("ENV") == "dev"
)

type OAuthHandler struct {
	githubOAuthInteractor usecase.IGithubOAuthInteractor
	userInteractor        usecase.IUserInteractor
}

func NewOAuthHandler(authGroup *echo.Group, githubOAuthInteractor usecase.IGithubOAuthInteractor, userInteractor usecase.IUserInteractor) {
	oAuthHandler := &OAuthHandler{
		githubOAuthInteractor: githubOAuthInteractor,
		userInteractor:        userInteractor,
	}

	authGroup.GET("/login", oAuthHandler.Login)
	authGroup.GET("/callback", oAuthHandler.CallBack)
}

func (h *OAuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	state := h.githubOAuthInteractor.GenerateState(STATE_LENGTH)

	stateCookie := h.githubOAuthInteractor.GenerateStateCookie(state, isDev)
	c.SetCookie(stateCookie)

	url := h.githubOAuthInteractor.GetGithubOAuthURL(ctx, state)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OAuthHandler) CallBack(c echo.Context) error {
	ctx := c.Request().Context()

	// Check State
	state := c.QueryParam("state")
	stateCookie, err := c.Cookie(config.OAUTH_STATE_COOKIE_NAME)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	if state != stateCookie.Value {
		log.Default().Println(err)
		return c.JSON(http.StatusUnauthorized, err)
	}

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
		if err == sql.ErrNoRows {
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

	atCookie := h.githubOAuthInteractor.GenerateAccessTokenCookie(token, isDev)

	c.SetCookie(atCookie)

	return c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONT_CALLBACK_URL"))
}
