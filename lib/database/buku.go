package database

import (
	"go_perpustakaan/config"
	"go_perpustakaan/models"
)

func GetBuku() (buku []models.Buku, err error) {
	err = config.DB.Find(&buku).Error

	if err != nil {
		return []models.Buku{}, err
	}

	return
}

func CreateBuku(buku models.Buku) (models.Buku, error) {
	err := config.DB.Create(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func UpdateBuku(buku models.Buku, id int) (models.Buku, error) {
	err := config.DB.Table("bukus").Where("id = ?", id).Updates(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func UpdateBukuStock(buku models.Buku, id int) (models.Buku, error) {
	err := config.DB.Table("bukus").Where("id = ?", id).Save(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func UpdateBukuStockTitle(buku models.Buku, judul string) (models.Buku, error) {
	err := config.DB.Table("bukus").Where("judul = ?", judul).Save(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func GetBukuById(id any) (models.Buku, error) {
	var buku models.Buku

	err := config.DB.Where("id = ?", id).First(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func GetBukuByJudul(judul string) (models.Buku, error) {
	var buku models.Buku

	err := config.DB.Where("judul = ?", judul).First(&buku).Error

	if err != nil {
		return models.Buku{}, err
	}
	return buku, nil
}

func GetBukuAllJudul(judul string) ([]models.Buku, error) {
	var buku []models.Buku

	err := config.DB.Where("judul LIKE ?", "%"+judul+"%").Find(&buku).Error

	if err != nil {
		return nil, err
	}
	return buku, nil
}

func DeleteBuku(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Buku{}).Error

	if err != nil {
		return nil, err
	}

	return "Buku behasil dihapus", nil
}
