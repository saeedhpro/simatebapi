package appointmentController

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/appointmentRepository"
	"github.com/saeedhpro/apisimateb/repository/caseTypeRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type AppointmentControllerInterface interface {
	CreateAppointment(c *gin.Context)
	GetAppointment(c *gin.Context)
	GetAppointmentResults(c *gin.Context)
	GetUserAppointmentList(c *gin.Context)
	GetUserResultedAppointmentList(c *gin.Context)
	GetOrganizationAppointmentList(c *gin.Context)
	FilterOrganizationAppointment(c *gin.Context)
	GetQueList(c *gin.Context)
	AcceptAppointment(c *gin.Context)
	UpdateAppointment(c *gin.Context)
	CancelAppointment(c *gin.Context)
	ReserveAppointment(c *gin.Context)
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
	staff := token.GetStaffUser(c)
	staffOrg, err := organizationRepository.GetOrganizationByID(staff.OrganizationID)
	if err != nil {
		c.JSON(500, "staff org")
		return
	}
	filter := models.AppointmentModel{}
	if staffOrg.IsPhotography() {
		filter.PhotographyID = staffOrg.ID
		filter.Photography = staffOrg
		filter.Status = 2
	} else if staffOrg.IsLaboratory() {
		filter.LaboratoryID = staffOrg.ID
		filter.Laboratory = staffOrg
		filter.Status = 2
	} else if staffOrg.IsRadiology() {
		filter.RadiologyID = staffOrg.ID
		filter.Radiology = staffOrg
		filter.Status = 2
	} else {
		filter.OrganizationID = staffOrg.ID
		filter.Organization = staffOrg
	}
	isDoctor := staffOrg.IsDoctor()
	if page < 1 {
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, start, end, isDoctor, false)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, start, end, isDoctor, false, page, limit)
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) GetAppointmentResults(c *gin.Context) {
	results := []string{}
	id, _ := strconv.Atoi(c.Param("id"))
	t := c.Query("type")
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	staff := token.GetStaffUser(c)
	staffOrg, err := organizationRepository.GetOrganizationByID(staff.OrganizationID)
	if err != nil {
		c.JSON(500, "Staff Organization not found")
		return
	}
	rnd := ""
	if appointment.LResultAt == nil &&
		appointment.RResultAt == nil &&
		appointment.PResultAt == nil {
		c.JSON(200, results)
		return
	}
	if staffOrg.IsPhotography() {
		rnd = appointment.PRndImg
	} else if staffOrg.IsLaboratory() {
		rnd = appointment.LRndImg
	} else if staffOrg.IsRadiology() {
		rnd = appointment.RRndImg
	} else {
		if t == "photography" {
			rnd = appointment.PRndImg
		} else if t == "radiology" {
			rnd = appointment.RRndImg
		} else if t == "laboratory" {
			rnd = appointment.LRndImg
		}
	}
	route := fmt.Sprintf("img/result/%d/%s", appointment.ID, rnd)
	files, err := ioutil.ReadDir(fmt.Sprintf("./res/%s", route))
	if err != nil {
		fmt.Println("read files", err.Error())
		c.JSON(200, results)
		return
	}
	for _, f := range files {
		results = append(results, fmt.Sprintf("http://%s/%s/%s", c.Request.Host, route, f.Name()))
	}
	c.JSON(200, results)
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
	response, _ := appointmentRepository.FilterOrganizationAppointment(organizationID, statues, q, start, end, false, page, limit)
	c.JSON(200, response)
}

func (u *AppointmentControllerStruct) GetQueList(c *gin.Context) {
	startDate := fmt.Sprintf("%s 00:00:00", c.Query("start"))
	endDate := fmt.Sprintf("%s 23:59:59", c.Query("end"))
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
	isDoctor := false
	if organization.IsPhotography() {
		filter.PhotographyID = organization.ID
		filter.Photography = organization
	} else if organization.IsLaboratory() {
		filter.LaboratoryID = organization.ID
		filter.Laboratory = organization
	} else if organization.IsRadiology() {
		filter.RadiologyID = organization.ID
		filter.Radiology = organization
	} else {
		filter.OrganizationID = organization.ID
		filter.Organization = organization
		isDoctor = true
	}
	if page < 1 {
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, "", "", isDoctor, true)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, "", "", isDoctor, true, page, limit)
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) GetUserResultedAppointmentList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userID, _ := strconv.Atoi(c.Param("id"))
	t := c.Query("type")
	if t == "" {
		t = "photography"
	}
	staff := token.GetStaffUser(c)
	organization, err := organizationRepository.GetOrganizationByID(staff.OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	filter := models.AppointmentModel{
		UserID: uint64(userID),
	}
	if organization.IsDoctor() {
		filter.OrganizationID = organization.ID
	} else {
		if t == "photography" {
			filter.PhotographyID = organization.ID
			filter.Photography = organization
		} else if t == "radiology" {
			filter.RadiologyID = organization.ID
			filter.Radiology = organization
		}
	}
	fmt.Println(filter.Radiology)
	if page < 1 {
		response, _ := appointmentRepository.GetResultedAppointmentListBy(&filter, t)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedResultedAppointmentListBy(&filter, t, page, limit)
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) AcceptAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	organization, err := organizationRepository.GetOrganizationByID(token.GetStaffUser(c).OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	var request requests.AppointmentUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		if err != nil {
			c.JSON(422, "request bind error")
			return
		}
	}
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "appointment not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	if organization.ID != appointment.OrganizationID || !organization.IsDoctor() {
		c.JSON(403, "access denied")
		return
	}
	staff := token.GetStaffUser(c).OrganizationID
	staffOrg, _ := organizationRepository.GetOrganizationByID(staff)
	if staffOrg.IsPhotography() {
		if len(request.Results) > 0 {
			t := time.Now().Format("2006-04-01 11:35:54")
			request.PAdmissionAt = &t
		}
	} else if staffOrg.IsLaboratory() {
		if len(request.Results) > 0 {
			t := time.Now().Format("2006-04-01 11:35:54")
			request.LAdmissionAt = &t
		}
	} else if staffOrg.IsRadiology() {
		if len(request.Results) > 0 {
			t := time.Now().Format("2006-04-01 11:35:54")
			request.RAdmissionAt = &t
		}
	}
	response, err := appointmentRepository.AcceptAppointment(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) UpdateAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	organization, err := organizationRepository.GetOrganizationByID(token.GetStaffUser(c).OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	var request requests.AppointmentUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		if err != nil {
			c.JSON(422, "request bind error")
			return
		}
	}
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "appointment not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	if !appointment.CanUpdate(organization) {
		c.JSON(403, "access denied")
		return
	}
	rand := ""
	staff := token.GetStaffUser(c).OrganizationID
	staffOrg, _ := organizationRepository.GetOrganizationByID(staff)
	if len(request.Results) > 0 {
		if staffOrg.IsPhotography() {
			t := time.Now().Format("2006-01-02 15:04:05")
			request.PResultAt = &t
			if appointment.PRndImg == "" {
				rand := helpers.RandomString(6)
				request.PRndImg = rand
			} else {
				request.PRndImg = appointment.PRndImg
				rand = appointment.PRndImg
			}
		} else if staffOrg.IsLaboratory() {
			t := time.Now().Format("2006-01-02 15:04:05")
			request.LResultAt = &t
			if appointment.LRndImg == "" {
				rand := helpers.RandomString(6)
				request.LRndImg = rand
			} else {
				request.LRndImg = appointment.LRndImg
				rand = appointment.LRndImg
			}
		} else if staffOrg.IsRadiology() {
			t := time.Now().Format("2006-01-02 15:04:05")
			request.RResultAt = &t
			if appointment.RRndImg == "" {
				rand := helpers.RandomString(6)
				request.RRndImg = rand
			} else {
				request.RRndImg = appointment.RRndImg
				rand = appointment.RRndImg
			}
		}
	}
	response, err := appointmentRepository.UpdateAppointment(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	if len(request.Results) > 0 {
		for i := 0; i < len(request.Results); i++ {
			location := fmt.Sprintf("./res/img/result/%d/%s", appointment.ID, rand)
			files, err := ioutil.ReadDir(location)
			if err != nil {
				if os.IsNotExist(err) {
					err = os.MkdirAll(location, os.ModePerm)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
			names := []string{}
			for i := 0; i < len(files); i++ {
				names = append(names, files[i].Name())
			}
			helpers.SaveImageToDisk(location, names, request.Results[i])
		}
	}
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) CancelAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	organization, err := organizationRepository.GetOrganizationByID(token.GetStaffUser(c).OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "appointment not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	if organization.ID != appointment.OrganizationID || !organization.IsDoctor() {
		c.JSON(403, "access denied")
		return
	}
	response, err := appointmentRepository.CancelAppointment(appointment)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) ReserveAppointment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	organization, err := organizationRepository.GetOrganizationByID(token.GetStaffUser(c).OrganizationID)
	if err != nil {
		c.JSON(500, "get staff organization error")
		return
	}
	appointment, err := appointmentRepository.GetAppointmentByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "appointment not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	if organization.ID != appointment.OrganizationID || !organization.IsDoctor() {
		c.JSON(403, "access denied")
		return
	}
	response, err := appointmentRepository.ReserveAppointment(appointment)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, response)
	return
}
