package provinceController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository/countyRepository"
	"github.com/saeedhpro/apisimateb/repository/provinceRepository"
	"gorm.io/gorm"
	"strconv"
)

type ProvinceControllerInterface interface {
	GetProvince(c *gin.Context)
	GetProvinceList(c *gin.Context)
}

type ProvinceControllerStruct struct {

}

func NewProvinceController() ProvinceControllerInterface {
	x := ProvinceControllerStruct{}
	return &x
}

func (p ProvinceControllerStruct) GetProvince(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	province, err := provinceRepository.GetProvinceByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	conditions := models.CountyModel{ProvinceID: province.ID}
	counties, _ := countyRepository.GetCountyListBy(&conditions)
	province.Counties = counties
	c.JSON(200, province)
}

func (p ProvinceControllerStruct) GetProvinceList(c *gin.Context) {
	provinces, err := provinceRepository.GetProvinceListBy(&models.ProvinceModel{})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	for i := 0; i < len(provinces); i++ {
		conditions := models.CountyModel{ProvinceID: provinces[i].ID}
		counties, _ := countyRepository.GetSimpleCountyListBy(&conditions)
		provinces[i].Counties = counties
	}
	c.JSON(200, provinces)
}
