package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tmp-friends/victo-api/app/presentation/handler"
)

const (
	apiVersion = "/v1"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// health check
	healthzGroup := e.Group("/healthz")
	{
		healthzGroup.GET("", handler.Check())
	}

	return e
}
