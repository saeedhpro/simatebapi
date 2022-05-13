package cityRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetCityByID(ID uint64) (*models.CityModel, error) {
	city := models.CityModel{ID: ID}
	err := repository.DB.MySQL.First(&city, &city).Error
	return &city, err
}

func GetCityListBy(conditions *models.CityModel) ([]models.CityModel, error) {
	cities := []models.CityModel{}
	err := repository.DB.MySQL.Preload("County").Find(&cities, &conditions).Error
	if err != nil {
		return cities, err
	}
	return cities, err
}