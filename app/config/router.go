package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/domain/service"
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

	// health check
	healthzGroup := e.Group("/healthz")
	{
		healthzGroup.GET("", handler.Check())
	}

	mysqlConnector := NewMySQLConnector()
	hashtagRepository := mysql.NewHashtagRepository(mysqlConnector.Conn)
	hashtagService := service.NewHashtagService(hashtagRepository)
	hashtagUsecase := usecase.NewHashtagUsecase(hashtagService)

	// hashtag
	hashtagGroup := e.Group("/hashtags")
	{
		hashtagHandler := handler.NewHashtagHandler(hashtagUsecase)
		hashtagGroup.GET("", hashtagHandler.FindHashtags())
	}

	return e
}
