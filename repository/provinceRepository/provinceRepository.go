package provinceRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetProvinceByID(ID uint64) (*models.ProvinceModel, error) {
	province := models.ProvinceModel{ID: ID}
	err := repository.DB.MySQL.First(&province, &province).Error
	return &province, err
}


func GetProvinceListBy(conditions *models.ProvinceModel) ([]models.ProvinceModel, error) {
	provinces := []models.ProvinceModel{}
	err := repository.DB.MySQL.Find(&provinces, &conditions).Error
	if err != nil {
		return provinces, err

	}
	return provinces, err
}
