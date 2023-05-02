package database

import (
	"go_perpustakaan/config"
	"go_perpustakaan/models"
)

func GetPeminjaman() (peminjaman []models.Peminjaman, err error) {
	err = config.DB.Preload("Mahasiswa").Preload("Buku").Find(&peminjaman).Error

	if err != nil {
		return []models.Peminjaman{}, err
	}

	return
}

func GetPeminjamanById(id any) (peminjaman models.Peminjaman, err error) {
	err = config.DB.Table("peminjamen").Where("id = ?", id).Find(&peminjaman).Error

	if err != nil {
		return models.Peminjaman{}, err
	}

	return peminjaman, nil
}

func CreatePeminjaman(peminjaman models.Peminjaman) (models.Peminjaman, error) {
	err := config.DB.Create(&peminjaman).Error

	if err != nil {
		return models.Peminjaman{}, err
	}

	return peminjaman, nil
}

func UpdatePeminjaman(peminjaman models.Peminjaman, id any) (models.Peminjaman, error) {
	err := config.DB.Table("peminjamen").Where("id = ?", id).Updates(&peminjaman).Error

	if err != nil {
		return models.Peminjaman{}, err
	}

	return peminjaman, nil
}

func DeletePeminjaman(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Peminjaman{}).Error

	if err != nil {
		return nil, err
	}

	return "Peminjaman behasil dihapus", nil
}
