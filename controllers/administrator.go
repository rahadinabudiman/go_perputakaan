package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdministratorsController(c echo.Context) error {
	admins, err := database.GetAdministrator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all administrator",
		Data:    admins,
	})
}
