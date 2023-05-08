package controllers

import (
	"go_perpustakaan/lib/database"
	"go_perpustakaan/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func GetPeminjamansController(c echo.Context) error {
	peminjaman, err := database.GetPeminjaman()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	if len(peminjaman) == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Data tidak ada",
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

	// tambahkan deklarasi tanggal dan jam sekarang
	now := time.Now()
	peminjaman.Tanggal_pinjam = now
	peminjaman.Tanggal_kembali = now.AddDate(0, 0, 7) // tambahkan 7 hari dari tanggal sekarang

	if err := c.Validate(peminjaman); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Ambil NIM dari middleware
	nim, ok := c.Get("nim").(int)
	if !ok || nim == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "NIM tidak tersedia",
		})
	}

	if peminjaman.NIM != nim {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "NIM tidak sama dengan yang login",
		})
	}

	// Ubah status mahasiswa
	mahasiswa, err := database.GetMahasiswaByNIM(peminjaman.NIM)
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
	buku, err := database.GetBukuByJudul(peminjaman.Judul)
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
	if _, err := database.UpdateMahasiswaByNIM(mahasiswa, peminjaman.NIM); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku.Stock--
	if _, err := database.UpdateBukuStockTitle(buku, peminjaman.Judul); err != nil {
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

	peminjamanresponse := models.PeminjamanResponse{
		NIM:             peminjaman.NIM,
		Judul:           peminjaman.Judul,
		Tanggal_kembali: peminjaman.Tanggal_kembali,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create peminjaman",
		Data:    peminjamanresponse,
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

func CreatePeminjamanAdminController(c echo.Context) error {
	peminjaman := models.Peminjaman{}
	c.Bind(&peminjaman)

	// tambahkan deklarasi tanggal dan jam sekarang
	now := time.Now()
	peminjaman.Tanggal_pinjam = now
	peminjaman.Tanggal_kembali = now.AddDate(0, 0, 7) // tambahkan 7 hari dari tanggal sekarang

	if err := c.Validate(peminjaman); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Ambil NIM dari middleware
	role, ok := c.Get("role").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Role tidak tersedia",
		})
	}

	if role != "Admin" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Anda bukan Admin",
		})
	}

	// Ubah status mahasiswa
	mahasiswa, err := database.GetMahasiswaByNIM(peminjaman.NIM)
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
	buku, err := database.GetBukuByJudul(peminjaman.Judul)
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
	if _, err := database.UpdateMahasiswaByNIM(mahasiswa, peminjaman.NIM); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	buku.Stock--
	if _, err := database.UpdateBukuStockTitle(buku, peminjaman.Judul); err != nil {
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

	peminjamanresponse := models.PeminjamanResponse{
		NIM:             peminjaman.NIM,
		Judul:           peminjaman.Judul,
		Tanggal_kembali: peminjaman.Tanggal_kembali,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create peminjaman",
		Data:    peminjamanresponse,
	})
}
