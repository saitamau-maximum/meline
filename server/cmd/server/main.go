package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/controller/handler"
	"github.com/saitamau-maximum/meline/controller/gateway"
	"github.com/saitamau-maximum/meline/infra/auth"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	e := echo.New()

	db, err := infra.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	oAuthConf := auth.NewGithubOAuthConf()
	authRepository := auth.NewAuthRepository(oAuthConf)
	userRepository := infra.NewUserRepository(bunDB)
	authInteractor := usecase.NewAuthInteractor(authRepository)
	userInteractor := usecase.NewUserInteractor(userRepository)
	authGatetway := gateway.NewAuthGateway(userInteractor)
	authHandler := handler.NewAuthHandler(authInteractor, userInteractor)
	userHandler := handler.NewUserHandler(userInteractor)

	apiGroup := e.Group("/api")
	apiGroup.GET("/", authGatetway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}))
	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", authHandler.Login)
	authGroup.GET("/callback", authHandler.CallBack)
	authGroup.POST("/signup", userHandler.CreateUser)

	e.Start(":8000")
}

