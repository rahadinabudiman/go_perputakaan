package models

import "github.com/jinzhu/gorm"

type Administrator struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
