package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/controller/handler"
	"github.com/saitamau-maximum/meline/controller/gateway"
	"github.com/saitamau-maximum/meline/infra/github"
	"github.com/saitamau-maximum/meline/infra/mysql"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	e := echo.New()

	db, err := mysql.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	apiGroup := e.Group("/api")

	oAuthConf := github.NewGithubOAuthConf()
	oAuthRepository := github.NewOAuthRepository(oAuthConf)
	userRepository := mysql.NewUserRepository(bunDB)
	authInteractor := usecase.NewGithubOAuthInteractor(oAuthRepository)
	userInteractor := usecase.NewUserInteractor(userRepository)
	authGateway := gateway.NewAuthGateway(userInteractor)
	
	authGroup := apiGroup.Group("/auth")
	handler.NewOAuthHandler(authGroup, authInteractor, userInteractor)

	apiGroup.GET("/", authGateway.Auth(func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}))

	e.Start(":8000")
}

