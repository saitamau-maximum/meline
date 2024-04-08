package gateway_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/saitamau-maximum/meline/adapter/gateway"
	"github.com/saitamau-maximum/meline/adapter/response"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type FailAuthUserNotFound string

const (
	FailAuthUserNotFoundValue FailAuthUserNotFound = "user_not_found"
)

var (
	ErrorResponseUnauthorized        response.ErrorResponse = response.ErrorResponse{Message: "Unauthorized"}
	ErrorResponseInternalServerError response.ErrorResponse = response.ErrorResponse{Message: "Internal Server Error"}
)

func TestAuthGateway_Auth_Success(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Create a new context with a mock user interactor
	ctx := context.Background()
	mockUserInteractor := &mockUserInteractor{}

	// Set the user ID in the claims
	claims := jwt.MapClaims{
		"user_id": float64(1),
		"exp":     float64(time.Now().Add(time.Hour).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.JWT_SECRET))
	cookie := &http.Cookie{
		Name:  config.ACCESS_TOKEN_COOKIE_NAME,
		Value: signedToken,
	}
	req.AddCookie(cookie)

	// Create a new echo context
	c := e.NewContext(req, rec)
	c.SetRequest(c.Request().WithContext(ctx))

	// Create a new AuthGateway instance
	authGateway := gateway.NewAuthGateway(mockUserInteractor)

	// Call the Auth method
	err := authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Authorized", rec.Body.String())
	assert.Equal(t, uint64(1), c.Get("user_id"))
}

func TestAuthGateway_Auth_Unauthorized_NoAccessTokenCookie(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Create a new context
	ctx := context.Background()

	// Create a new echo context
	c := e.NewContext(req, rec)
	c.SetRequest(c.Request().WithContext(ctx))

	// Create a new AuthGateway instance
	authGateway := &gateway.AuthGateway{}

	// Call the Auth method
	authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})(c)

	res := response.ErrorResponse{}

	json.Unmarshal(rec.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, ErrorResponseUnauthorized, res)
	assert.Equal(t, nil, c.Get("user_id"))
}

func TestAuthGateway_Auth_Unauthorized_InvalidToken(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Create a new context
	ctx := context.Background()

	// Set an invalid token in the cookie
	cookie := &http.Cookie{
		Name:  config.ACCESS_TOKEN_COOKIE_NAME,
		Value: "invalid_token",
	}
	req.AddCookie(cookie)

	// Create a new echo context
	c := e.NewContext(req, rec)
	c.SetRequest(c.Request().WithContext(ctx))

	// Create a new AuthGateway instance
	authGateway := &gateway.AuthGateway{}

	// Call the Auth method
	authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})(c)

	res := response.ErrorResponse{}

	json.Unmarshal(rec.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, ErrorResponseUnauthorized, res)
	assert.Equal(t, nil, c.Get("user_id"))
}

func TestAuthGateway_Auth_Unauthorized_ExpiredToken(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Create a new context
	ctx := context.Background()

	// Set an expired token in the cookie
	claims := jwt.MapClaims{
		"user_id": float64(1),
		"exp":     float64(time.Now().Add(-time.Hour).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.JWT_SECRET))
	cookie := &http.Cookie{
		Name:  config.ACCESS_TOKEN_COOKIE_NAME,
		Value: signedToken,
	}
	req.AddCookie(cookie)

	// Create a new echo context
	c := e.NewContext(req, rec)
	c.SetRequest(c.Request().WithContext(ctx))

	// Create a new AuthGateway instance
	authGateway := &gateway.AuthGateway{}

	// Call the Auth method
	authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})(c)

	res := response.ErrorResponse{}

	json.Unmarshal(rec.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, ErrorResponseUnauthorized, res)
	assert.Equal(t, nil, c.Get("user_id"))
}

func TestAuthGateway_Auth_Unauthorized_UserNotFound(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Create a new context with a mock user interactor
	ctx := context.Background()
	mockUserInteractor := &mockUserInteractor{}
	ctx = context.WithValue(ctx, FailAuthUserNotFoundValue, true)

	// Set the user ID in the claims
	claims := jwt.MapClaims{
		"user_id": float64(1),
		"exp":     float64(time.Now().Add(time.Hour).Unix()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.JWT_SECRET))
	cookie := &http.Cookie{
		Name:  config.ACCESS_TOKEN_COOKIE_NAME,
		Value: signedToken,
	}
	req.AddCookie(cookie)

	// Create a new echo context
	c := e.NewContext(req, rec)
	c.SetRequest(c.Request().WithContext(ctx))

	// Create a new AuthGateway instance
	authGateway := gateway.NewAuthGateway(mockUserInteractor)

	// Call the Auth method
	authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})(c)

	res := response.ErrorResponse{}

	json.Unmarshal(rec.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, ErrorResponseInternalServerError, res)
	assert.Equal(t, nil, c.Get("user_id"))
}

type mockUserInteractor struct{}

func (m *mockUserInteractor) GetUserByID(ctx context.Context, id uint64) (*presenter.GetUserByIdResponse, error) {
	return nil, nil
}

func (m *mockUserInteractor) GetUserByGithubIDOrCreate(ctx context.Context, githubID, userName, imageUrl string) (*presenter.GetUserByGithubIdResponse, error) {
	return nil, nil
}

func (m *mockUserInteractor) CreateUser(ctx context.Context, githubID, name, imageURL string) (*presenter.CreateUserResponse, error) {
	return nil, nil
}

func (m *mockUserInteractor) IsUserExists(ctx context.Context, userID uint64) (bool, error) {
	if ctx.Value(FailAuthUserNotFoundValue) != nil {
		return false, errors.New("user not found")
	}

	return true, nil
}
