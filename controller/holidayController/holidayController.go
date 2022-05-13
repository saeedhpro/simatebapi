package holidayController

import (
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/holidayRepository"
)

type HolidayControllerInterface interface {
	GetOrganizationHolidaysList(c *gin.Context)
}

type HolidayControllerStruct struct {
}

func NewHolidayController() HolidayControllerInterface {
	x := HolidayControllerStruct{}
	return &x
}

func (h HolidayControllerStruct) GetOrganizationHolidaysList(c *gin.Context) {
	organizationID := token.GetStaffUser(c).OrganizationID
	startDate := c.Query("start")
	endDate := c.Query("end")
	holidays, _ := holidayRepository.GetHolidayListBy(&models.HolidayModel{OrganizationID: &organizationID}, "", startDate, endDate)
	c.JSON(200, holidays)
}
