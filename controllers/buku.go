package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBukusController(c echo.Context) error {
	bukus, err := database.GetBuku()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	if len(bukus) == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Data tidak ada",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all buku",
		Data:    bukus,
	})
}

func GetBukuController(c echo.Context) error {
	id := c.Param("id")

	buku, err := database.GetBukuById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get buku",
		Data:    buku,
	})
}

func CreateBukuController(c echo.Context) error {
	buku := models.Buku{}
	c.Bind(&buku)

	if err := c.Validate(buku); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku, err := database.CreateBuku(buku)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create buku",
		Data:    buku,
	})
}

func UpdateBukuController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku := models.Buku{}
	c.Bind(&buku)

	buku, err = database.UpdateBuku(buku, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update buku",
		Data:    buku,
	})
}

func DeleteBukuController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteBuku(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete buku",
	})
}
