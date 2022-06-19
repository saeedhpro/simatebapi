package appointmentController

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/appointmentRepository"
	"github.com/saeedhpro/apisimateb/repository/caseTypeRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type AppointmentControllerInterface interface {
	CreateAppointment(c *gin.Context)
	GetAppointment(c *gin.Context)
	GetAppointmentResults(c *gin.Context)
	GetUserAppointmentList(c *gin.Context)
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
	if staffOrg.ProfessionID == 1 {
		filter.PhotographyID = staffOrg.ID
		filter.Photography = staffOrg
	} else if staffOrg.ProfessionID == 2 {
		filter.LaboratoryID = staffOrg.ID
		filter.Laboratory = staffOrg
	} else if staffOrg.ProfessionID == 3 {
		filter.RadiologyID = staffOrg.ID
		filter.Radiology = staffOrg
	} else {
		filter.OrganizationID = staffOrg.ID
		filter.Organization = staffOrg
	}
	isDoctor := staffOrg.IsDoctor()
	if page < 1 {
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, start, end, isDoctor)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, start, end, isDoctor, page, limit)
	c.JSON(200, response)
	return
}

func (u *AppointmentControllerStruct) GetAppointmentResults(c *gin.Context) {
	results := []string{}
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
	if staffOrg.ProfessionID == 1 {
		rnd = appointment.PRndImg
	} else if staffOrg.ProfessionID == 2 {
		rnd = appointment.LRndImg
	} else if staffOrg.ProfessionID == 3 {
		rnd = appointment.RRndImg
	} else {
		c.JSON(422, "")
		return
	}
	route := fmt.Sprintf("img/result/%d/%s", appointment.ID, rnd)
	files, err := ioutil.ReadDir(fmt.Sprintf("./res/%s", route))
	if err != nil {
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
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
		response, _ := appointmentRepository.GetAppointmentListBy(&filter, "", "", false)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := appointmentRepository.GetPaginatedAppointmentListBy(&filter, "", "", false, page, limit)
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
	response, err := appointmentRepository.UpdateAppointment(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	rand := ""
	staff := token.GetStaffUser(c).OrganizationID
	staffOrg, _ := organizationRepository.GetOrganizationByID(staff)
	if staffOrg.ProfessionID == 1 {

	}
	for i := 0; i < len(request.Results); i++ {
		saveImageToDisk(fmt.Sprintf("%d/%s/%d", appointment.OrganizationID, rand, i), request.Results[i])
	}
	c.JSON(200, response)
	return
}

func saveImageToDisk(fileNameBase string, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", fmt.Errorf("invalid image")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		log.Println("errpeed")
		return "", err
	}
	//imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	//if err != nil {
	//	log.Println("errpeed 2")
	//	return "", err
	//}
	//
	//if imgCfg.Width != 750 || imgCfg.Height != 685 {
	//	return "", fmt.Errorf("invalid size")
	//}

	fileName := fileNameBase + "." + "jpeg"
	err = ioutil.WriteFile(fmt.Sprintf("./res/img/result/%s", fileName), buff.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error(), "cf")
		return "", fmt.Errorf("cant save file")
	}
	return fileName, err
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
