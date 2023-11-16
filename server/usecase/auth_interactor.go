package usecase

import (
	crand "crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/saitamau-maximum/meline/config"
)

type IAuthInteractor interface {
	GenerateState(stateLength int) string
	GenerateStateCookie(state string, isDev bool) *http.Cookie
	GenerateAccessTokenCookie(token string, isDev bool) *http.Cookie
}

type AuthInteractor struct {
}

func NewAuthInteractor() IAuthInteractor {
	return &AuthInteractor{}
}

func (i *AuthInteractor) GenerateState(stateLength int) string {
	k := make([]byte, stateLength)
	if _, err := crand.Read(k); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", k)
}

func (i *AuthInteractor) GenerateStateCookie(state string, isDev bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = config.OAUTH_STATE_COOKIE_NAME
	cookie.Value = state
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = !isDev
	cookie.Expires = time.Now().Add(5 * time.Minute)

	return cookie
}

func (i *AuthInteractor) GenerateAccessTokenCookie(token string, isDev bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = config.ACCESS_TOKEN_COOKIE_NAME
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = !isDev
	cookie.Expires = time.Now().Add(config.ACCESS_TOKEN_EXPIRE)

	return cookie
}
