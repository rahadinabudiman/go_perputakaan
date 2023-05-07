package middlewares

import (
	"go_perpustakaan/constants"
	"go_perpustakaan/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateToken(id int, email, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}

func CreateTokenMahasiswa(id, nim int, email, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["nim"] = nim
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_JWT),
	TokenLookup:   "cookie:JWTCookie",
	ErrorHandler: func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	},
},
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or missing jwt token")
		}
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt claims")
		}
		if role, ok := claims["role"].(string); !ok || role != "Admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "user is not an admin")
		}
		return next(c)
	}
}

func JWTValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("JWTCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET_JWT), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		nim, ok := claims["nim"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		c.Set("nim", int(nim))
		return next(c)
	}
}
