package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAdministratorsController(c echo.Context) error {
	admins, err := database.GetAdministrator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	if len(admins) == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Data tidak ada",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all administrator",
		Data:    admins,
	})
}

func GetAdministratorController(c echo.Context) error {
	id := c.Param("id")

	admin, err := database.GetAdministratorById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get administrator",
		Data:    admin,
	})
}

func CreateAdministratorController(c echo.Context) error {
	admin := models.Administrator{}
	c.Bind(&admin)

	if err := c.Validate(admin); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	admin, err := database.CreateAdministrator(admin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create administrator",
		Data:    admin,
	})
}

func UpdateAdministratorController(c echo.Context) error {
	id := c.Param("id")

	admin := models.Administrator{}
	c.Bind(&admin)

	admin, err := database.UpdateAdministrator(admin, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update administrator",
		Data:    admin,
	})
}

func DeleteAdministratorController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteAdministrator(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete administrator",
	})
}
