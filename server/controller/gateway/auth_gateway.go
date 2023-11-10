package gateway

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/usecase"
)

type IAuthGateway interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthGateway struct {
	userInteractor usecase.IUserInteractor
}

func NewAuthGateway(userInteractor usecase.IUserInteractor) IAuthGateway {
	return &AuthGateway{
		userInteractor: userInteractor,
	}
}

func (h *AuthGateway) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// Get Access Token
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
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
		userId := claims["user_id"].(float64)

		user, err := h.userInteractor.GetUserByID(ctx, uint64(userId))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		c.Set("user", user)

		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			return c.JSON(http.StatusForbidden, err)
		}

		return next(c)
	}
}
