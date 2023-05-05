package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Mahasiswa struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	NIM      int    `json:"nim" form:"nim"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Prodi    string `json:"prodi" form:"prodi"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'Mahasiswa');default:'Mahasiswa'; not-null"`
	Status   string `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
}

// For Response
type MahasiswaResponse struct {
	gorm.Model
	NIM   int    `json:"nim" form:"nim"`
	Email string `json:"email" form:"email"`
	Prodi string `json:"prodi" form:"prodi"`
}

// For JWT Only
type MahasiswaResponses struct {
	ID    uint   `json:"id" form:"id"`
	NIM   int    `json:"nim" form:"nim"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

type CustomValidator struct {
	Validators *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validators.Struct(i)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
