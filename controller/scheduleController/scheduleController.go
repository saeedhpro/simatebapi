package scheduleController

import (
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/scheduleRepository"
	"strconv"
)

type ScheduleControllerInterface interface {
	GetOrganizationScheduleList(c *gin.Context)
}

type ScheduleControllerStruct struct {
}

func NewScheduleController() ScheduleControllerInterface {
	x := ScheduleControllerStruct{}
	return &x
}

func (s *ScheduleControllerStruct) GetOrganizationScheduleList(c *gin.Context) {
	staff := token.GetStaffUser(c)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startDate := c.Query("start")
	endDate := c.Query("end")
	filter := models.VipScheduleModel{
		OrganizationID: staff.OrganizationID,
	}
	if page < 1 {
		response, _ := scheduleRepository.GetScheduleListBy(&filter, startDate, endDate)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := scheduleRepository.GetPaginatedScheduleListBy(&filter, startDate, endDate, page, limit)
	c.JSON(200, response)
}
