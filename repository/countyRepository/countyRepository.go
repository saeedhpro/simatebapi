package countyRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/cityRepository"
)

func GetCountyByID(ID uint64) (*models.CountyModel, error) {
	county := models.CountyModel{ID: ID}
	err := repository.DB.MySQL.Preload("Province").First(&county, &county).Error
	return &county, err
}

func GetCountyListBy(conditions *models.CountyModel) ([]models.CountyModel, error) {
	counties := []models.CountyModel{}
	err := repository.DB.MySQL.Find(&counties, &conditions).Error
	if err != nil {
		return counties, err
	}
	for i := 0; i < len(counties); i++ {
		cities, _ := cityRepository.GetCityListBy(&models.CityModel{CountyID: counties[i].ID})
		counties[i].Cities = cities
	}
	return counties, err
}

func GetSimpleCountyListBy(conditions *models.CountyModel) ([]models.CountyModel, error) {
	counties := []models.CountyModel{}
	err := repository.DB.MySQL.Find(&counties, &conditions).Error
	if err != nil {
		return counties, err
	}
	return counties, err
}
