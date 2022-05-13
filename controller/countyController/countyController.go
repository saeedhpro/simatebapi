package countyController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository/cityRepository"
	"github.com/saeedhpro/apisimateb/repository/countyRepository"
	"gorm.io/gorm"
	"strconv"
)

type CountyControllerInterface interface {
	GetCounty(c *gin.Context)
	GetCountyList(c *gin.Context)
}

type CountyControllerStruct struct {

}

func NewCountyController() CountyControllerInterface {
	x := CountyControllerStruct{}
	return &x
}

func (p CountyControllerStruct) GetCounty(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	county, err := countyRepository.GetCountyByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	conditions := models.CityModel{CountyID: county.ID}
	cities, _ := cityRepository.GetCityListBy(&conditions)
	county.Cities = cities
	c.JSON(200, county)
}

func (p CountyControllerStruct) GetCountyList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	counties, err := countyRepository.GetCountyListBy(&models.CountyModel{ProvinceID: uint64(id)})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	for i := 0; i < len(counties); i++ {
		conditions := models.CityModel{CountyID: counties[i].ID}
		cities, _ := cityRepository.GetCityListBy(&conditions)
		counties[i].Cities = cities
	}
	c.JSON(200, counties)
}
