package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPengembalianController(c echo.Context) error {
	pengembalian, err := database.GetPengembalian()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all pengembalian",
		Data:    pengembalian,
	})
}

func GetPengembalianByIdController(c echo.Context) error {
	id := c.Param("id")

	pengembalian, err := database.GetPeminjamanById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get pengembalian",
		Data:    pengembalian,
	})
}

func CreatePengembalianController(c echo.Context) error {
	pengembalian := models.Pengembalian{}
	c.Bind(&pengembalian)

	if err := c.Validate(pengembalian); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Check Apakah Mahasiswa Tersebut Meminjam Atau Tidak
	peminjaman, err := database.GetPeminjamanBy2Id(pengembalian.MahasiswaID, pengembalian.BukuID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	statuspin, err := database.GetPeminjamanById(peminjaman.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	statuspin.Status = "1"
	if _, err := database.UpdateStatusPeminjaman(statuspin, statuspin.ID); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Ubah status mahasiswa
	mahasiswa, err := database.GetMahasiswaById(pengembalian.MahasiswaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}
	if mahasiswa.Status == "0" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Mahasiswa belum meminjam buku",
		})
	}

	// Kurangi Stock Buku Saat Berhasil Dipinjam
	buku, err := database.GetBukuById(pengembalian.BukuID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	mahasiswa.Status = "0"
	if _, err := database.UpdateMahasiswa(mahasiswa, pengembalian.MahasiswaID); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku.Stock++
	if _, err := database.UpdateBukuStock(buku, pengembalian.BukuID); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	pengembalian, err = database.CreatePengembalian(pengembalian)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create pengembalian",
		Data:    pengembalian,
	})
}

func UpdatePengembalianController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	pengembalian := models.Pengembalian{}
	c.Bind(&pengembalian)

	pengembalian, err = database.UpdatePengembalian(pengembalian, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update pengembalian",
		Data:    pengembalian,
	})
}

func DeletePengembalianController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	_, err = database.DeletePengembalian(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete pengembalian",
	})
}
