package appointmentController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/appointmentRepository"
	"github.com/saeedhpro/apisimateb/repository/caseTypeRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type AppointmentControllerInterface interface {
	CreateAppointment(c *gin.Context)
	GetAppointment(c *gin.Context)
	GetUserAppointmentList(c *gin.Context)
	GetOrganizationAppointmentList(c *gin.Context)
	FilterOrganizationAppointment(c *gin.Context)
	GetQueList(c *gin.Context)
}

type AppointmentControllerStruct struct {
}

func NewAppointmentController() AppointmentControllerInterface {
	x := AppointmentControllerStruct{}
	return &x
}

func (u *AppointmentControllerStruct) CreateAppointment(c *gin.Context) {
	var request requests.AppointmentCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	staff := token.GetStaffUser(c)
	appointment, err := appointmentRepository.CreateAppointment(&request, staff.UserID, staff.OrganizationID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, &appointment)
	return
}

func (u *AppointmentControllerStruct) GetAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, appointment)
	return
}

func (u *AppointmentControllerStruct) GetOrganizationAppointmentList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	start := c.Query("start")
	end := c.Query("end")
	organizationID := token.GetStaffUser(c).OrganizationID
	filter := models.AppointmentModel{OrganizationID: organizationID}
	if page < 1 {
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, start, end)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, start, end, page, limit)
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) FilterOrganizationAppointment(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	start := c.Query("start")
	end := c.Query("end")
	status := c.Query("status")
	statues := []string{}
	if status != "" {
		statues = strings.Split(status, ",")
	}
	organizationID := token.GetStaffUser(c).OrganizationID
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.FilterOrganizationAppointment(organizationID, statues, q, start, end, page, limit)
	c.JSON(200, response)
}

func (u *AppointmentControllerStruct) GetQueList(c *gin.Context) {
	startDate := c.Query("start")
	endDate := c.Query("end")
	var que models.QueStruct
	organizationID := token.GetStaffUser(c).OrganizationID
	ques, _ := appointmentRepository.GetAppointmentListBetweenDates(&organizationID, startDate, endDate)
	que.Ques = ques
	organization, err := organizationRepository.GetOrganizationByID(organizationID)
	if err == nil {
		que.WorkHour = models.WorkHour{Start: organization.WorkHourStart, End: organization.WorkHourEnd}
	} else {
		que.WorkHour = models.WorkHour{Start: "16:00:00", End: "21:00:00"}
	}
	limits, _ := caseTypeRepository.GetCaseTypeListBy(&models.CaseType{OrganizationID: organizationID})
	que.Limits = limits
	c.JSON(200, que)
	return
}

func (u *AppointmentControllerStruct) GetUserAppointmentList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userID, _ := strconv.Atoi(c.Param("id"))
	staff := token.GetStaffUser(c)
	organization, err := organizationRepository.GetOrganizationByID(staff.OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	filter := models.AppointmentModel{
		UserID: uint64(userID),
	}
	switch organization.ProfessionID {
	case 1:
		filter.PhotographyID = organization.ID
		break
	case 2:
		filter.LaboratoryID = organization.ID
		break
	case 3:
		filter.RadiologyID = organization.ID
		break
	default:
		filter.OrganizationID = organization.ID
		break
	}
	if page < 1 {
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, "", "")
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, "", "", page, limit)
	c.JSON(200, response)
	return
}
