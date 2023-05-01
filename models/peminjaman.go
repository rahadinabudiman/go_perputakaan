package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Peminjaman struct {
	gorm.Model
	BukuID          uint      `json:"buku_id" form:"buku_id"`
	MahasiswaID     uint      `json:"mahasiswa_id" form:"mahasiswa_id"`
	Tanggal_pinjam  time.Time `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_kembali time.Time `json:"tanggal_kembali" form:"tanggal_kembali"`
}
