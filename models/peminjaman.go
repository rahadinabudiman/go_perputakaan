package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Peminjaman struct {
	gorm.Model
	MahasiswaID     int       `json:"mahasiswa_id" form:"mahasiswa_id"`
	Mahasiswa       Mahasiswa `json:"mahasiswa"`
	BukuID          int       `json:"buku_id" form:"buku_id"`
	Buku            Buku      `json:"buku"`
	Tanggal_pinjam  time.Time `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_kembali time.Time `json:"tanggal_kembali" form:"tanggal_kembali"`
	Status          string    `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
}
