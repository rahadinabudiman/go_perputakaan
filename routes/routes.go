package routes

import (
	"go_perpustakaan/controllers"
	m "go_perpustakaan/middlewares"
	"go_perpustakaan/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	m.Log(e)
	cv := &util.CustomValidator{Validators: validator.New()}
	e.Validator = cv

	// All Routes
	e.GET("/cookie", controllers.GetCookieHandler)

	// Mahasiswa Routes
	mahasiswa := e.Group("/mahasiswa")
	mahasiswa.POST("/login", controllers.LoginMahasiswaController)             // Login Mahasiswa ex : {local}/mahasiswa/login
	mahasiswa.POST("", controllers.CreateMahasiswaController)                  // Create Mahasiswa
	mahasiswa.PUT("/:id", controllers.UpdateMahasiswaController, m.IsLoggedIn) // Edit Mahasiswa

	mahasiswa.GET("/buku", controllers.GetBukusController, m.IsLoggedIn)                  // Get All Buku
	mahasiswa.GET("/buku/:id", controllers.GetBukuController, m.IsLoggedIn)               // Get Buku by ID
	mahasiswa.GET("/buku/title/:title", controllers.GetBukuTitleController, m.IsLoggedIn) // Get Buku by ID

	mahasiswa.POST("/pinjam", controllers.CreatePeminjamanController, m.IsLoggedIn, m.JWTValidator) // Pinjam Buku

	mahasiswa.POST("/kembali", controllers.CreatePengembalianController, m.IsLoggedIn, m.JWTValidator) // Kembalikan Buku

	// Administrator Routes
	administrator := e.Group("/administrator")
	// Administrator Data Routes
	administrator.POST("", controllers.CreateAdministratorController)
	administrator.POST("/login", controllers.LoginAdministratorController)
	administrator.GET("", controllers.GetAdministratorsController, m.IsLoggedIn, m.IsAdmin)
	administrator.GET("/:id", controllers.GetAdministratorController, m.IsLoggedIn, m.IsAdmin)
	administrator.PUT("/:id", controllers.UpdateAdministratorController, m.IsLoggedIn, m.IsAdmin)
	administrator.DELETE("/:id", controllers.DeleteAdministratorController, m.IsLoggedIn, m.IsAdmin)

	// Mahasiswa Data Routes
	administrator.GET("/mahasiswa", controllers.GetMahasiswaController, m.IsLoggedIn, m.IsAdmin)
	administrator.PUT("/mahasiswa/:id", controllers.UpdateMahasiswaController, m.IsLoggedIn, m.IsAdmin)
	administrator.GET("/mahasiswa/:id", controllers.GetMahasiswaByIdController, m.IsLoggedIn, m.IsAdmin)
	administrator.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswaController, m.IsLoggedIn, m.IsAdmin)

	// Buku Data Routes
	administrator.GET("/buku", controllers.GetBukusController, m.IsLoggedIn, m.IsAdmin)
	administrator.GET("/buku/:id", controllers.GetBukuController, m.IsLoggedIn, m.IsAdmin)
	administrator.POST("/buku", controllers.CreateBukuController, m.IsLoggedIn, m.IsAdmin)
	administrator.PUT("/buku/:id", controllers.UpdateBukuController, m.IsLoggedIn, m.IsAdmin)
	administrator.DELETE("/buku/:id", controllers.DeleteBukuController, m.IsLoggedIn, m.IsAdmin)

	// Pinjaman Buku
	administrator.GET("/pinjam", controllers.GetPeminjamansController, m.IsLoggedIn, m.IsAdmin)
	administrator.GET("/pinjam/:id", controllers.GetPeminjamanByIdController, m.IsLoggedIn, m.IsAdmin)
	administrator.POST("/pinjam", controllers.CreatePeminjamanAdminController, m.IsLoggedIn, m.IsAdmin, m.JWTValidatorAdmin)
	administrator.PUT("/pinjam/:id", controllers.UpdatePeminjamanController, m.IsLoggedIn, m.IsAdmin)
	administrator.DELETE("/pinjam/:id", controllers.DeletePeminjamanController, m.IsLoggedIn, m.IsAdmin)

	administrator.GET("/kembali", controllers.GetPengembalianController, m.IsLoggedIn, m.IsAdmin)
	administrator.GET("/kembali/:id", controllers.GetPengembalianByIdController, m.IsLoggedIn, m.IsAdmin)
	administrator.POST("/kembali", controllers.CreatePengembalianController, m.IsLoggedIn, m.IsAdmin)
	administrator.PUT("/kembali/:id", controllers.UpdatePengembalianController, m.IsLoggedIn, m.IsAdmin)
	administrator.DELETE("/kembali/:id", controllers.DeletePengembalianController, m.IsLoggedIn, m.IsAdmin)

	return e
}
