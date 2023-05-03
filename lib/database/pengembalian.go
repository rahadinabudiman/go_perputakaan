package database

import (
	"go_perpustakaan/config"
	"go_perpustakaan/models"
)

func GetPengembalian() (pengembalian []models.Pengembalian, err error) {
	err = config.DB.Preload("Mahasiswa").Preload("Buku").Find(&pengembalian).Error

	if err != nil {
		return []models.Pengembalian{}, err
	}

	return
}

func GetPengembalianById(id any) (pengembalian models.Pengembalian, err error) {
	err = config.DB.Table("pengembalians").Where("id = ?", id).Find(&pengembalian).Error

	if err != nil {
		return models.Pengembalian{}, err
	}

	return pengembalian, nil
}

func CreatePengembalian(pengembalian models.Pengembalian) (models.Pengembalian, error) {
	err := config.DB.Create(&pengembalian).Error

	if err != nil {
		return models.Pengembalian{}, err
	}

	return pengembalian, nil
}

func UpdatePengembalian(pengembalian models.Pengembalian, id any) (models.Pengembalian, error) {
	err := config.DB.Table("pengembalians").Where("id = ?", id).Updates(&pengembalian).Error

	if err != nil {
		return models.Pengembalian{}, err
	}

	return pengembalian, nil
}

func DeletePengembalian(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Pengembalian{}).Error

	if err != nil {
		return nil, err
	}

	return "Peminjaman behasil dihapus", nil
}
