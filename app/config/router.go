package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// health check
	e.GET("/healthz", func(c echo.Context) error {
		message := "OK"
		return c.JSON(
			http.StatusOK,
			message,
		)
	})

	return e
}
