package models

import (
	"github.com/jinzhu/gorm"
)

type Pengembalian struct {
	gorm.Model
	NIM   int    `json:"nim" form:"nim" validate:"required"`
	Judul string `json:"judul" form:"judul" validate:"required"`
}

// For Response Peminjaman
type PengembalianResponse struct {
	NIM   int    `json:"nim" form:"nim"`
	Judul string `json:"judul" form:"judul"`
}
