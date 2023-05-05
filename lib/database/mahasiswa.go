package database

import (
	"go_perpustakaan/config"
	"go_perpustakaan/models"
)

func GetMahasiswa() (mahasiswa []models.Mahasiswa, err error) {
	err = config.DB.Find(&mahasiswa).Error

	if err != nil {
		return []models.Mahasiswa{}, err
	}
	return
}

func GetMahasiswaById(id any) (mahasiswa models.Mahasiswa, err error) {
	err = config.DB.Table("mahasiswas").Where("id = ?", id).Find(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func GetMahasiswaByNIM(nim any) (mahasiswa models.Mahasiswa, err error) {
	err = config.DB.Table("mahasiswas").Where("nim = ?", nim).Find(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func CreateMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error) {
	err := config.DB.Create(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func UpdateMahasiswa(mahasiswa models.Mahasiswa, id int) (models.Mahasiswa, error) {
	err := config.DB.Table("mahasiswas").Where("id = ?", id).Updates(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func UpdateMahasiswaByNIM(mahasiswa models.Mahasiswa, nim any) (models.Mahasiswa, error) {
	err := config.DB.Table("mahasiswas").Where("nim = ?", nim).Updates(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func DeleteMahasiswa(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Mahasiswa{}).Error

	if err != nil {
		return nil, err
	}

	return "Mahasiswa behasil dihapus", nil
}

func LoginMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error) {
	err := config.DB.Where("nim = ? AND password = ?", mahasiswa.NIM, mahasiswa.Password).First(&mahasiswa).Error

	if err != nil {
		return models.Mahasiswa{}, nil
	}

	return mahasiswa, nil
}
