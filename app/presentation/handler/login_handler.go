package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmp-friends/victo-api/app/usecase"
)

type loginHandler struct {
	usecase usecase.ILoginUsecase
}

func NewLoginHandler(lu usecase.ILoginUsecase) *loginHandler {
	return &loginHandler{
		usecase: lu,
	}
}

type Post struct {
	IdToken string `json:"idToken"`
}

func (lu *loginHandler) CreateSessionCookie() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := new(Post)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		cookieValue, expiresIn, err := lu.usecase.CreateSessionCookie(ctx, p.IdToken)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.Value = cookieValue
		cookie.MaxAge = int(expiresIn.Seconds())
		cookie.HttpOnly = true
		// TODO: Secureもtrueにする

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, "write a cookie")
	}
}
