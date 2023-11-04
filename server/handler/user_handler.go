package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase"
)

type IUserHandler interface {
	LoggedIn(ctx echo.Context) error
	SignUp(ctx echo.Context) error
}

type UserHandler struct {
	userInteractor usecase.IUserInteractor
	authInteractor usecase.IAuthInteractor
}

func NewUserHandler(userInteractor usecase.IUserInteractor, authInteractor usecase.IAuthInteractor) IUserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
		authInteractor: authInteractor,
	}
}

func (h *UserHandler) LoggedIn(c echo.Context) error {
	token, err := c.Cookie("access_token")
	if err != nil {
		return err
	}

	respByte, err := h.authInteractor.GetGithubUser(c.Request().Context(), token.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var resp map[string]interface{}

	if err := json.Unmarshal(respByte, &resp); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user := entity.User{
		GithubID: resp["login"].(string),
		Name:     resp["name"].(string),
	}

	if _, err := h.userInteractor.GetUserByGithubID(c.Request().Context(), user.GithubID); err != nil {
		userByte, err := json.Marshal(user)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		
		c.SetCookie(&http.Cookie{
			Name:  "user",
			Value: string(userByte),
		})

		return c.Redirect(http.StatusFound, "auth/signup")
	}

	return c.String(http.StatusOK, "Authorized")
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var user entity.User

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := h.userInteractor.CreateUser(c.Request().Context(), user.GithubID, user.Name); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Created")
}
