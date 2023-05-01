package models

import "github.com/jinzhu/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama   string `json:"name" form:"name"`
	NIM    string `json:"nim" form:"nim"`
	Prodi  string `json:"prodi" form:"prodi"`
	Status int    `json:"status" form:"status"`
}
