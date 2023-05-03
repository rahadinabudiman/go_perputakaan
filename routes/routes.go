package routes

import (
	"go_perpustakaan/controllers"
	m "go_perpustakaan/middlewares"
	"go_perpustakaan/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	m.Log(e)
	e.Validator = &models.CustomValidator{Validators: validator.New()}
	e.POST("/mahasiswa", controllers.CreateMahasiswaController)

	administrator := e.Group("/administrator")
	administrator.GET("", controllers.GetAdministratorsController)
	administrator.GET("/:id", controllers.GetAdministratorController)
	administrator.POST("", controllers.CreateAdministratorController)
	administrator.PUT("/:id", controllers.UpdateAdministratorController)
	administrator.DELETE("/:id", controllers.DeleteAdministratorController)

	mahasiswa := e.Group("/mahasiswa")
	mahasiswa.GET("", controllers.GetMahasiswaController)
	mahasiswa.GET("/:id", controllers.GetMahasiswaByIdController)
	mahasiswa.PUT("/:id", controllers.UpdateMahasiswaController)
	mahasiswa.DELETE("/:id", controllers.DeleteMahasiswaController)

	buku := e.Group("/buku")
	buku.GET("", controllers.GetBukusController)
	buku.GET("/:id", controllers.GetBukuController)
	buku.POST("", controllers.CreateBukuController)
	buku.PUT("/:id", controllers.UpdateBukuController)
	buku.DELETE("/:id", controllers.DeleteBukuController)

	peminjaman := e.Group("/pinjam")
	peminjaman.GET("", controllers.GetPeminjamansController)
	peminjaman.GET("/:id", controllers.GetPeminjamanByIdController)
	peminjaman.POST("", controllers.CreatePeminjamanController)
	peminjaman.PUT("/:id", controllers.UpdatePeminjamanController)
	peminjaman.DELETE("/:id", controllers.DeletePeminjamanController)

	pengembalian := e.Group("/pengembalian")
	pengembalian.GET("", controllers.GetPengembalianController)
	pengembalian.GET("/:id", controllers.GetPengembalianByIdController)
	pengembalian.POST("", controllers.CreatePengembalianController)
	pengembalian.PUT("/:id", controllers.UpdatePengembalianController)
	pengembalian.DELETE("/:id", controllers.DeletePengembalianController)

	return e
}
