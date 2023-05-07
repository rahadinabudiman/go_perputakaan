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

	administratorresp := make([]models.AdminResponse, len(admins))
	for i, admin := range admins {
		administratorresp[i] = models.AdminResponse{
			Nama:  admin.Nama,
			Email: admin.Email,
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all admin",
		Data:    administratorresp,
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

	administratorresp := models.AdminResponse{
		Nama:  admin.Nama,
		Email: admin.Email,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create admin",
		Data:    administratorresp,
	})
}

func CreateAdministratorController(c echo.Context) error {
	admin := models.Administrator{}
	c.Bind(&admin)

	if err := c.Validate(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	admins, err := database.CreateAdministrator(admin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	administratorresp := models.AdminResponseCreate{
		Nama:     admins.Nama,
		Email:    admins.Email,
		Password: admins.Password,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create admin",
		Data:    administratorresp,
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
