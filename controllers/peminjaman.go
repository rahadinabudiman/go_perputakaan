package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPeminjamansController(c echo.Context) error {
	peminjaman, err := database.GetPeminjaman()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all peminjaman",
		Data:    peminjaman,
	})
}

func GetPeminjamanByIdController(c echo.Context) error {
	id := c.Param("id")

	peminjaman, err := database.GetPeminjamanById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get peminjaman",
		Data:    peminjaman,
	})
}

func CreatePeminjamanController(c echo.Context) error {
	peminjaman := models.Peminjaman{}
	c.Bind(&peminjaman)

	if err := c.Validate(peminjaman); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Ubah status mahasiswa
	mahasiswa, err := database.GetMahasiswaById(peminjaman.MahasiswaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	if mahasiswa.Status == "1" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Mahasiswa sudah meminjam buku",
		})
	}

	// Kurangi Stock Buku Saat Berhasil Dipinjam
	buku, err := database.GetBukuById(peminjaman.BukuID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	if buku.Stock < 1 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Buku tidak tersedia",
		})
	}

	mahasiswa.Status = "1"
	if _, err := database.UpdateMahasiswa(mahasiswa, peminjaman.MahasiswaID); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku.Stock--
	if _, err := database.UpdateBukuStock(buku, peminjaman.BukuID); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	peminjaman, err = database.CreatePeminjaman(peminjaman)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create peminjaman",
		Data: map[string]interface{}{
			"peminjaman": peminjaman,
		},
	})
}

func UpdatePeminjamanController(c echo.Context) error {
	id := c.Param("id")

	peminjaman := models.Peminjaman{}
	c.Bind(&peminjaman)

	peminjaman, err := database.UpdatePeminjaman(peminjaman, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update peminjaman",
		Data:    peminjaman,
	})
}

func DeletePeminjamanController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeletePeminjaman(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete peminjaman",
	})
}
