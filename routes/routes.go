package routes

import (
	"go_perpustakaan/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	administrator := e.Group("/administrator")
	administrator.GET("", controllers.GetAdministratorsController)
	administrator.GET("/:id", controllers.GetAdministratorController)
	administrator.POST("", controllers.CreateAdministratorController)
	administrator.PUT("/:id", controllers.UpdateAdministratorController)
	administrator.DELETE("/:id", controllers.DeleteAdministratorController)

	return e
}
