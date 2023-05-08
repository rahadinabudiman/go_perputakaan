package models

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Mahasiswa struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama" validate:"required"`
	NIM      int    `json:"nim" form:"nim" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Prodi    string `json:"prodi" form:"prodi" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'Mahasiswa');default:'Mahasiswa'; not-null"`
	Status   string `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
}

type MahasiswaLogin struct {
	NIM      string `json:"nim" form:"nim" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
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
		var sb strings.Builder
		sb.WriteString("Validation error:\n")

		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("- %s\n", err))
		}

		return echo.NewHTTPError(http.StatusBadRequest, sb.String())
	}

	return nil
}
