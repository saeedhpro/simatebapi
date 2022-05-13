package cityController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository/cityRepository"
	"gorm.io/gorm"
	"strconv"
)

type CityControllerInterface interface {
	GetCity(c *gin.Context)
	GetCityList(c *gin.Context)
}

type CityControllerStruct struct {

}

func NewCityController() CityControllerInterface {
	x := CityControllerStruct{}
	return &x
}

func (p CityControllerStruct) GetCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	city, err := cityRepository.GetCityByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, city)
}

func (p CityControllerStruct) GetCityList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	counties, err := cityRepository.GetCityListBy(&models.CityModel{CountyID: uint64(id)})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, counties)
}
