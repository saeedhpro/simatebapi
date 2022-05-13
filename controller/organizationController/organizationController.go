package organizationController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"gorm.io/gorm"
	"strconv"
)

type OrganizationControllerInterface interface {
	Get(c *gin.Context)
}

type OrganizationControllerStruct struct {

}

func NewOOrganizationController() OrganizationControllerInterface {
	x := OrganizationControllerStruct{}
	return &x
}

func (o OrganizationControllerStruct) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := organizationRepository.GetOrganizationByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, response)
}