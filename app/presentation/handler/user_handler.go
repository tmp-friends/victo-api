package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmp-friends/victo-api/app/usecase"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) GetMyInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		session, err := c.Cookie("session")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		u, err := uh.usecase.GetMyInfo(ctx, session)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return c.JSON(http.StatusOK, u)
	}
}

type LoginPost struct {
	IdToken string `json:"idToken"`
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := new(LoginPost)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		cookieValue, expiresIn, err := uh.usecase.Login(ctx, p.IdToken)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		cookie := uh.generateCookie(cookieValue, expiresIn)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "write a cookie")
	}
}

func (uh *userHandler) generateCookie(cookieValue string, expiresIn int) *http.Cookie {
	cookie := new(http.Cookie)

	cookie.Name = "session"
	cookie.Value = cookieValue
	cookie.MaxAge = expiresIn
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true

	return cookie
}

func (uh *userHandler) FindFollowingHashtags() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		p := c.Param("id")

		hashtags, err := uh.usecase.FindFollowingHashtags(ctx, p)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return c.JSON(http.StatusOK, hashtags)
	}
}
