package models

import "github.com/jinzhu/gorm"

type Buku struct {
	gorm.Model
	Judul        string `json:"judul" form:"judul" validate:"required"`
	Penulis      string `json:"penulis" form:"penulis" validate:"required"`
	Tahun_terbit string `json:"tahun_terbit" form:"tahun_terbit" validate:"required"`
	ISBN         int    `json:"isbn" form:"isbn" validate:"required"`
	Stock        int    `json:"stock" form:"stock" validate:"required"`
}
