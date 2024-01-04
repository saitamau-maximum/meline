package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/adapter/gateway"
	"github.com/saitamau-maximum/meline/adapter/handler"
	"github.com/saitamau-maximum/meline/adapter/presenter"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/infra/github"
	"github.com/saitamau-maximum/meline/infra/mysql"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	err := config.ValidateAppEnv()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	db, err := mysql.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.RegisterModel((*model.ChannelUsers)(nil), (*model.ChannelToChannels)(nil), (*model.MessageToMessages)(nil), (*model.Channel)(nil), (*model.User)(nil), (*model.Message)(nil))
	defer bunDB.Close()

	apiGroup := e.Group("/api")

	oAuthConf := github.NewGithubOAuthConf()
	oAuthRepository := github.NewOAuthRepository(oAuthConf)
	userRepository := mysql.NewUserRepository(bunDB)
	channelRepository := mysql.NewChannelRepository(bunDB)
	channelUsersRepository := mysql.NewChannelUsersRepository(bunDB)
	channelToChannelsRepository := mysql.NewChannelToChannelsRepository(bunDB)
	messageRepository := mysql.NewMessageRepository(bunDB)
	messageToMessagesRepository := mysql.NewMessageToMessagesRepository(bunDB)
	githubOAuthInteractor := usecase.NewGithubOAuthInteractor(oAuthRepository)
	authInteractor := usecase.NewAuthInteractor()
	channelInteractor := usecase.NewChannelInteractor(channelRepository, channelUsersRepository, userRepository, channelToChannelsRepository, presenter.NewChannelPresenter())
	messageInteractor := usecase.NewMessageInteractor(messageRepository, messageToMessagesRepository)
	userPresenter := presenter.NewUserPresenter()
	userInteractor := usecase.NewUserInteractor(userRepository, userPresenter)
	authGateway := gateway.NewAuthGateway(userInteractor)

	handler.NewOAuthHandler(apiGroup.Group("/auth"), githubOAuthInteractor, authInteractor, userInteractor)
	handler.NewUserHandler(apiGroup.Group("/user", authGateway.Auth), userInteractor)
	handler.NewChannelHandler(apiGroup.Group("/channels", authGateway.Auth), channelInteractor)
	handler.NewMessageHandler(apiGroup.Group("/messages", authGateway.Auth), messageInteractor)

	apiGroup.GET("/", authGateway.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}))

	e.Start(":8000")
}
