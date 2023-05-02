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

func UpdateBuku(buku models.Buku, id any) (models.Buku, error) {
	result := config.DB.Model(&models.Buku{}).Where("id = ?", id).Updates(models.Buku{Judul: buku.Judul, Penulis: buku.Penulis, Tahun_terbit: buku.Tahun_terbit, ISBN: buku.ISBN, Stock: buku.Stock})
	if result.Error != nil {
		return models.Buku{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.Buku{}, result.Error
	}

	updatedBuku := models.Buku{}
	err := config.DB.Where("id = ?", id).Find(&updatedBuku).Error
	if err != nil {
		return models.Buku{}, err
	}

	return updatedBuku, nil
}

func GetBukuById(id any) (models.Buku, error) {
	var buku models.Buku

	err := config.DB.Where("id = ?", id).First(&buku).Error

	if err != nil {
		return models.Buku{}, err
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
