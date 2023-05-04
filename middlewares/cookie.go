package middlewares

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func CreateCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "JWTCookie"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
