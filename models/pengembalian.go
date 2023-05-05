package models

import (
	"github.com/jinzhu/gorm"
)

type Pengembalian struct {
	gorm.Model
	NIM       int       `json:"nim" form:"nim"`
	Mahasiswa Mahasiswa `json:"mahasiswa"`
	Judul     string    `json:"judul" form:"judul"`
	Buku      Buku      `json:"buku"`
}

// For Response Peminjaman
type PengembalianResponse struct {
	NIM   int    `json:"nim" form:"nim"`
	Judul string `json:"judul" form:"judul"`
}
