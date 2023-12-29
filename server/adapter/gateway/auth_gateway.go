package gateway

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase"
)

type AuthGateway struct {
	userInteractor usecase.IUserInteractor
}

func NewAuthGateway(userInteractor usecase.IUserInteractor) *AuthGateway {
	return &AuthGateway{
		userInteractor: userInteractor,
	}
}

func (h *AuthGateway) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Access Token
		cookie, err := c.Cookie(config.ACCESS_TOKEN_COOKIE_NAME)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unsupported signing method")
			}

			return []byte(config.JWT_SECRET), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, err)
		}

		userId := claims["user_id"].(float64)

		c.Set("user_id", uint64(userId))

		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			return c.JSON(http.StatusForbidden, err)
		}

		return next(c)
	}
}
