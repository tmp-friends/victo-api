package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmp-friends/victo-api/app/usecase/healthz"
)

func Check() echo.HandlerFunc {
	return func(c echo.Context) error {
		success := healthz.Execute()

		if !success {
			return c.JSON(http.StatusInternalServerError, c.Path())
		}

		return c.JSON(http.StatusOK, c.Path())
	}
}
