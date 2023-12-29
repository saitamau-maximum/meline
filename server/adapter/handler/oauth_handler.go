package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	STATE_LENGTH = 32
)

type OAuthHandler struct {
	githubOAuthInteractor usecase.IGithubOAuthInteractor
	authInteractor        usecase.IAuthInteractor
	userInteractor        usecase.IUserInteractor
}

func NewOAuthHandler(authGroup *echo.Group, githubOAuthInteractor usecase.IGithubOAuthInteractor, authInteractor usecase.IAuthInteractor, userInteractor usecase.IUserInteractor) {
	oAuthHandler := &OAuthHandler{
		githubOAuthInteractor: githubOAuthInteractor,
		authInteractor:        authInteractor,
		userInteractor:        userInteractor,
	}

	authGroup.GET("/login", oAuthHandler.Login)
	authGroup.GET("/callback", oAuthHandler.CallBack)
}

func (h *OAuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	state := h.authInteractor.GenerateState(STATE_LENGTH)

	stateCookie := h.authInteractor.GenerateStateCookie(state, config.IsDev)
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

	var userId uint64

	getUserRes, err := h.userInteractor.GetUserByGithubID(ctx, userRes.OAuthUserID, userRes.Name, userRes.ImageURL)
	if err != nil {
		return err
	}

	userId = getUserRes.ID

	// Set Access Token
	token, err := h.authInteractor.CreateAccessToken(ctx, userId)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	atCookie := h.authInteractor.GenerateAccessTokenCookie(token, config.IsDev)

	c.SetCookie(atCookie)

	return c.Redirect(http.StatusTemporaryRedirect, config.FRONT_OAUTH_SUCCESS_URL)
}
