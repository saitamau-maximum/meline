package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IAuthHandler interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
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

func (h *AuthHandler) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// Get Access Token
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, err)
		}

		// Get User
		userId := claims["id"].(uint64)

		user, err := h.userInteractor.GetUser(ctx, userId)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		c.Set("user", user)

		exp := claims["exp"].(int64)
		if exp < time.Now().Unix() {
			return c.JSON(http.StatusForbidden, err)
		}

		return next(c)
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
