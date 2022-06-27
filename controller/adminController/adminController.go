package adminController

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/domain/responses"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/groupRepository"
	"github.com/saeedhpro/apisimateb/repository/holidayRepository"
	"github.com/saeedhpro/apisimateb/repository/messageRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"github.com/saeedhpro/apisimateb/repository/professionRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type AdminControllerInterface interface {
	LastOnlineUsers(c *gin.Context)
	LastOnlinePatients(c *gin.Context)
	GetUsers(c *gin.Context)
	GetOrganizations(c *gin.Context)
	CreateOrganization(c *gin.Context)
	UpdateOrganization(c *gin.Context)
	GetOrganizationsByProfession(c *gin.Context)
	GetProfessions(c *gin.Context)
	GetUserGroups(c *gin.Context)
	GetMessages(c *gin.Context)
	DeleteMessages(c *gin.Context)
	GetHolidays(c *gin.Context)
	CreateHoliday(c *gin.Context)
	UpdateHoliday(c *gin.Context)
	DeleteHoliday(c *gin.Context)
}

type AdminControllerStruct struct {
}

func NewAdminController() AdminControllerInterface {
	x := AdminControllerStruct{}
	return &x
}

func (a *AdminControllerStruct) LastOnlineUsers(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	users, _ := userRepository.GetLastOnlineUsers()
	c.JSON(200, users)
}

func (a *AdminControllerStruct) LastOnlinePatients(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	users, _ := userRepository.GetLastOnlinePatients()
	c.JSON(200, users)
}

func (a *AdminControllerStruct) GetUsers(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userGroupID, _ := strconv.Atoi(c.Query("group"))
	q := c.Query("q")
	filter := models.UserModel{}
	if userGroupID > 0 {
		filter.UserGroupID = uint64(userGroupID)
	}
	if page < 1 {
		response, _ := userRepository.GetUserListBy(&filter, q)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	users, _ := userRepository.GetPaginatedUserListBy(&filter, q, page, limit)
	c.JSON(200, users)
}

func (a *AdminControllerStruct) GetOrganizations(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	professionID, _ := strconv.Atoi(c.Query("prof"))
	filter := models.OrganizationModel{}
	if professionID > 0 {
		filter.ProfessionID = uint64(professionID)
	}
	if page < 1 {
		response, _ := organizationRepository.GetOrganizationListBy(&filter, q)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := organizationRepository.GetPaginatedOrganizationListBy(&filter, q, page, limit)
	for i := 0; i < len(response.Data.([]models.OrganizationModel)); i++ {
		route := fmt.Sprintf("img/org/%d/%s.jpg", response.Data.([]models.OrganizationModel)[i].ID, response.Data.([]models.OrganizationModel)[i].Logo)
		response.Data.([]models.OrganizationModel)[i].Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, route)
	}
	c.JSON(200, response)
}

func (a *AdminControllerStruct) UpdateOrganization(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, "incorrect id")
		return
	}
	var request requests.CreateOrganizationRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, err.Error())
		return
	}
	organization, err := organizationRepository.GetOrganizationByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "not found")
			return
		} else {
			c.JSON(500, err.Error())
			return
		}
	}
	tx := repository.DB.MySQL.Begin()
	if request.New != nil {
		location := fmt.Sprintf("./res/img/org/%d", organization.ID)
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
		_, logo, err := helpers.SaveImageToDisk(location, names, *request.New)
		if err != nil {
			fmt.Println(err.Error())
		}
		request.Logo = logo
	}
	err = organizationRepository.UpdateOrganization(&request)
	if err != nil {
		tx.Rollback()
		c.JSON(500, err.Error())
		return
	}
	if len(request.RelOrganizations) > 0 {
		err = organizationRepository.UpdateRelOrganizations(organization.RelOrganizations, request.RelOrganizations)
		if err != nil {
			tx.Rollback()

			c.JSON(200, err.Error())
			return
		}
	}
	tx.Commit()
	c.JSON(200, organization)
	return
}

func (a *AdminControllerStruct) CreateOrganization(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	var request requests.CreateOrganizationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, err.Error())
		return
	}
	tx := repository.DB.MySQL.Begin()
	request.StaffID = staff.ID
	organization, err := organizationRepository.CreateOrganization(&request)
	if err != nil {
		tx.Rollback()
		c.JSON(500, err.Error())
		return
	}
	if request.New != nil {
		location := fmt.Sprintf("./res/img/org/%d", organization.ID)
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
		_, logo, err := helpers.SaveImageToDisk(location, names, *request.New)
		if err != nil {
			fmt.Println(err.Error())
		}
		request.Logo = logo
		request.CreatedAt = nil
		request.ID = organization.ID
		_ = organizationRepository.UpdateOrganization(&request)
	}
	if len(request.RelOrganizations) > 0 {
		for i := 0; i < len(request.RelOrganizations); i++ {
			request.RelOrganizations[i].Organization = organization
			request.RelOrganizations[i].OrganizationID = organization.ID
		}
		err = organizationRepository.UpdateRelOrganizations(organization.RelOrganizations, request.RelOrganizations)
		if err != nil {
			tx.Rollback()
			c.JSON(200, err.Error())
			return
		}
	}
	tx.Commit()
	c.JSON(200, organization)
	return
}

func (a *AdminControllerStruct) GetOrganizationsByProfession(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, "Invalid id")
		return
	}
	staffOrg, _ := organizationRepository.GetOrganizationByID(uint64(id))
	response := responses.OrganizationByProfessionResponse{}
	filter := models.OrganizationModel{}
	filter.ProfessionID = uint64(3)
	radiologies, _ := organizationRepository.GetOnlyRadiologyOrganizationList(&filter, staffOrg.ID)
	filter.ProfessionID = uint64(2)
	laboratories, _ := organizationRepository.GetOnlyLaboratoryOrganizationList(&filter, staffOrg.ID)
	filter.ProfessionID = uint64(1)
	photographies, _ := organizationRepository.GetOnlyPhotographyOrganizationList(&filter, staffOrg.ID)
	filter = models.OrganizationModel{}
	doctors, _ := organizationRepository.GetOnlyDoctorOrganizationList(&filter, staffOrg.ID)
	for i := 0; i < len(radiologies); i++ {
		route := fmt.Sprintf("img/org/%d/%s.jpg", radiologies[i].ID, radiologies[i].Logo)
		radiologies[i].Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, route)
	}
	for i := 0; i < len(laboratories); i++ {
		route := fmt.Sprintf("img/org/%d/%s.jpg", laboratories[i].ID, laboratories[i].Logo)
		laboratories[i].Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, route)
	}
	for i := 0; i < len(photographies); i++ {
		route := fmt.Sprintf("img/org/%d/%s.jpg", photographies[i].ID, photographies[i].Logo)
		photographies[i].Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, route)
	}
	for i := 0; i < len(doctors); i++ {
		route := fmt.Sprintf("img/org/%d/%s.jpg", doctors[i].ID, doctors[i].Logo)
		doctors[i].Logo = fmt.Sprintf("http://%s/%s", c.Request.Host, route)
	}
	response.Radiologies = radiologies
	response.Laboratories = laboratories
	response.Photographies = photographies
	response.Doctors = doctors
	c.JSON(200, response)
}

func (a *AdminControllerStruct) GetProfessions(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	filter := models.ProfessionModel{}
	if page < 1 {
		response, _ := professionRepository.GetProfessionListBy(&filter, q)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := professionRepository.GetPaginatedProfessionListBy(&filter, q, page, limit)
	c.JSON(200, response)
}

func (a *AdminControllerStruct) GetUserGroups(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	filter := models.UserGroupModel{}
	if page < 1 {
		response, _ := groupRepository.GetGroupListBy(&filter, q)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := groupRepository.GetPaginatedGroupListBy(&filter, q, page, limit)
	c.JSON(200, response)
}

func (a *AdminControllerStruct) GetMessages(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	filter := models.SmsModel{}
	if page < 1 {
		response, _ := messageRepository.GetMessageListBy(&filter, q)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := messageRepository.GetPaginatedMessageListBy(&filter, q, page, limit)
	c.JSON(200, response)
}

func (a *AdminControllerStruct) DeleteMessages(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	var request requests.DeleteMultipleItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, "parse failed")
		return
	}
	err := messageRepository.DeleteMessages(request.Ids)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
	return
}

func (a *AdminControllerStruct) GetHolidays(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startDate := c.Query("start")
	endDate := c.Query("end")
	q := c.Query("q")
	filter := models.HolidayModel{}
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

func (a *AdminControllerStruct) CreateHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
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
	holiday := models.HolidayModel{
		Title:          request.Title,
		Hdate:          hdate,
		OrganizationID: request.OrganizationID,
	}
	err = holidayRepository.CreateHoliday(&holiday)
	c.JSON(200, err)
	return
}

func (a *AdminControllerStruct) UpdateHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
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
		ID:             uint64(id),
		Title:          request.Title,
		Hdate:          hdate,
		OrganizationID: request.OrganizationID,
	}
	err = holidayRepository.UpdateHoliday(&holiday)
	c.JSON(200, err)
	return
}

func (a *AdminControllerStruct) DeleteHoliday(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	if !staff.IsAdmin() {
		c.JSON(403, "Access Denied!")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	holiday := models.HolidayModel{
		ID: uint64(id),
	}
	err = holidayRepository.DeleteHoliday(&holiday)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "")
	return
}
