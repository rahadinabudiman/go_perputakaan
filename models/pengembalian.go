package models

import (
	"github.com/jinzhu/gorm"
)

type Pengembalian struct {
	gorm.Model
	MahasiswaID int       `json:"mahasiswa_id" form:"mahasiswa_id"`
	Mahasiswa   Mahasiswa `json:"mahasiswa"`
	BukuID      int       `json:"buku_id" form:"buku_id"`
	Buku        Buku      `json:"buku"`
}
