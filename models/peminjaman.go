package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Peminjaman struct {
	gorm.Model
	NIM             int       `json:"nim" form:"nim"`
	Mahasiswa       Mahasiswa `json:"mahasiswa"`
	Judul           string    `json:"judul" form:"judul"`
	Buku            Buku      `json:"buku"`
	Tanggal_pinjam  time.Time `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_kembali time.Time `json:"tanggal_kembali" form:"tanggal_kembali"`
	Status          string    `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
}

// For Response Peminjaman

type PeminjamanResponse struct {
	NIM             int       `json:"nim" form:"nim"`
	Judul           string    `json:"judul" form:"judul"`
	Tanggal_kembali time.Time `json:"tanggal_kembali" form:"tanggal_kembali"`
}
