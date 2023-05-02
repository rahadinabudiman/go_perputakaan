package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Mahasiswa struct {
	gorm.Model
	Nama   string `json:"nama" form:"nama"`
	NIM    string `json:"nim" form:"nim"`
	Prodi  string `json:"prodi" form:"prodi"`
	Status string `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
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
