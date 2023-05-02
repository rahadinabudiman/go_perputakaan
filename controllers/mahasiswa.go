package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMahasiswaController(c echo.Context) error {
	mahasiswa, err := database.GetMahasiswa()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, models.Response{
		Message: "success get all mahasiswa",
		Data:    mahasiswa,
	})
}

func GetMahasiswaByIdController(c echo.Context) error {
	id := c.Param("id")

	mahasiswa, err := database.GetMahasiswaById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get mahasiswa",
		Data:    mahasiswa,
	})
}

func CreateMahasiswaController(c echo.Context) error {
	mahasiswa := models.Mahasiswa{}
	c.Bind(&mahasiswa)

	mahasiswa, err := database.CreateMahasiswa(mahasiswa)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create mahasiswa",
		Data:    mahasiswa,
	})
}

func UpdateMahasiswaController(c echo.Context) error {
	id := c.Param("id")

	mahasiswa := models.Mahasiswa{}
	c.Bind(&mahasiswa)

	mahasiwa, err := database.UpdateMahasiswa(mahasiswa, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update mahasiswa",
		Data:    mahasiwa,
	})
}

func DeleteMahasiswaController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteMahasiswa(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete mahasiswa",
	})
}
