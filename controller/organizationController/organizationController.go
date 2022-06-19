package organizationController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/holidayRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type OrganizationControllerInterface interface {
	Get(c *gin.Context)
	GetOrganizationByType(c *gin.Context)
	GetHolidays(c *gin.Context)
	CreateHoliday(c *gin.Context)
	UpdateHoliday(c *gin.Context)
	DeleteHoliday(c *gin.Context)
}

type OrganizationControllerStruct struct {
}

func NewOOrganizationController() OrganizationControllerInterface {
	x := OrganizationControllerStruct{}
	return &x
}

func (o *OrganizationControllerStruct) Get(c *gin.Context) {
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

func (o *OrganizationControllerStruct) GetOrganizationByType(c *gin.Context) {
	t := c.Query("type")
	response, err := organizationRepository.GetOrganizationByType(t)
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

func (o *OrganizationControllerStruct) GetHolidays(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startDate := c.Query("start")
	endDate := c.Query("end")
	q := c.Query("q")
	filter := models.HolidayModel{
		OrganizationID: &staff.OrganizationID,
	}
	if page < 1 {
		response, _ := holidayRepository.GetHolidayListBy(&filter, q, startDate, endDate)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := holidayRepository.GetPaginatedHolidayListBy(&filter, q, startDate, endDate, page, limit)
	c.JSON(200, response)
}

func (o *OrganizationControllerStruct) CreateHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	var request requests.HolidayCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	hdate, err := time.Parse("2006-01-02", request.Hdate)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	holiday := models.HolidayModel{
		Title:          request.Title,
		Hdate:          hdate,
		OrganizationID: &staff.OrganizationID,
	}
	err = holidayRepository.CreateHoliday(&holiday)
	c.JSON(200, err)
	return
}

func (o *OrganizationControllerStruct) UpdateHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	holiday, _ := holidayRepository.GetHolidayByID(uint64(id))
	if staff.OrganizationID != *holiday.OrganizationID {
		c.JSON(403, "Access Denied!")
		return
	}
	var request requests.HolidayCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	hdate, err := time.Parse("2006-01-02", request.Hdate)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	err = holidayRepository.UpdateHoliday(&models.HolidayModel{
		ID:             holiday.ID,
		Title:          request.Title,
		Hdate:          hdate,
		OrganizationID: &staff.OrganizationID,
	})
	c.JSON(200, err)
	return
}

func (o *OrganizationControllerStruct) DeleteHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	holiday, _ := holidayRepository.GetHolidayByID(uint64(id))
	if staff.OrganizationID != *holiday.OrganizationID {
		c.JSON(403, "Access Denied!")
		return
	}
	err = holidayRepository.DeleteHoliday(holiday)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "")
	return
}
