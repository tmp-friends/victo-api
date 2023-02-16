package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/infra/mysql"
	"github.com/tmp-friends/victo-api/app/presentation/handler"
	"github.com/tmp-friends/victo-api/app/usecase"
)

const (
	apiVersion = "/v1"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	mysqlConnector := NewMySQLConnector()

	// health check
	healthzGroup := e.Group("/healthz")
	{
		healthzGroup.GET("", handler.Check())
	}

	// hashtag
	hashtagQuery := mysql.NewHashtagQuery(mysqlConnector.Conn)
	hashtagUsecase := usecase.NewHashtagUsecase(hashtagQuery)

	hashtagsGroup := e.Group("/hashtags")
	{
		hashtagHandler := handler.NewHashtagHandler(hashtagUsecase)
		hashtagsGroup.GET("", hashtagHandler.FindHashtags())
	}

	// tweet
	tweetQuery := mysql.NewTweetQuery(mysqlConnector.Conn)
	tweetUsecase := usecase.NewTweetUsecase(tweetQuery)

	tweetsGroup := e.Group("/tweets")
	{
		tweetHandler := handler.NewTweetHandler(tweetUsecase)
		tweetsGroup.GET("/:id", tweetHandler.FindTweet())
	}

	return e
}
