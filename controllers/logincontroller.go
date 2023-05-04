package controllers

import (
	"fmt"
	"go_perpustakaan/lib/database"
	"go_perpustakaan/middlewares"
	"go_perpustakaan/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginMahasiswaController(c echo.Context) error {
	mahasiswa := models.Mahasiswa{}
	c.Bind(&mahasiswa)

	mahasiswa, err := database.LoginMahasiswa(mahasiswa)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(mahasiswa.ID), mahasiswa.Email, mahasiswa.Role)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	middlewares.CreateCookie(c, token)

	resp := models.MahasiswaResponses{
		ID:    mahasiswa.ID,
		NIM:   mahasiswa.NIM,
		Email: mahasiswa.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success login",
		Data:    resp,
	})
}

func LoginAdministratorController(c echo.Context) error {
	admin := models.Administrator{}
	c.Bind(&admin)

	admin, err := database.LoginAdministrator(admin)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(admin.ID), admin.Email, admin.Role)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	middlewares.CreateCookie(c, token)

	resp := models.AdminResponses{
		ID:    admin.ID,
		Nama:  admin.Nama,
		Email: admin.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success login",
		Data:    resp,
	})
}

func GetCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// handle jika cookie tidak ditemukan
			return c.String(http.StatusUnauthorized, err.Error())
		}
		// handle error lainnya
		return err
	}

	// handle jika cookie ditemukan
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "Cookie ditemukan")
}
