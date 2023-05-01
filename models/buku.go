package models

import "github.com/jinzhu/gorm"

type Buku struct {
	gorm.Model
	Judul        string `json:"judul" form:"judul"`
	Penulis      string `json:"penulis" form:"penulis"`
	Tahun_terbit string `json:"tahun_terbit" form:"tahun_terbit"`
	ISBN         int    `json:"isbn" form:"isbn"`
}
