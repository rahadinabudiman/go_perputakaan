package models

import (
	"github.com/jinzhu/gorm"
)

type Administrator struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'Mahasiswa');default:'Admin'; not-null"`
}

type AdminResponse struct {
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
}

type AdminResponseCreate struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// For JWT Only
type AdminResponses struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
