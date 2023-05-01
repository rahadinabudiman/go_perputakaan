package database

import (
	"go_perpustakaan/config"
	"go_perpustakaan/models"
)

func GetAdministrator() (admin []models.Administrator, err error) {
	err = config.DB.Find(&admin).Error

	if err != nil {
		return []models.Administrator{}, err
	}

	return
}

func CreateAdministrator(admin models.Administrator) (models.Administrator, error) {
	err := config.DB.Create(&admin).Error

	if err != nil {
		return models.Administrator{}, err
	}
	return admin, nil
}

func UpdateAdministrator(admin models.Administrator, id any) (models.Administrator, error) {
	err := config.DB.Table("administrators").Where("id = ?", id).Updates(&admin).Error

	if err != nil {
		return models.Administrator{}, err
	}
	return admin, nil
}

func GetAdministratorById(id any) (models.Administrator, error) {
	var admin models.Administrator

	err := config.DB.Where("id = ?", id).First(&admin).Error

	if err != nil {
		return models.Administrator{}, err
	}

	return admin, nil
}

func DeleteAdministrator(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Administrator{}).Error

	if err != nil {
		return nil, err
	}

	return "Administrator behasil dihapus", nil
}
