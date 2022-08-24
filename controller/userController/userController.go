package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	user2 "github.com/saeedhpro/apisimateb/helpers/user"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/countyRepository"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"github.com/saeedhpro/apisimateb/repository/provinceRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type UserControllerInterface interface {
	Own(c *gin.Context)
	GetOrganizationUsersList(c *gin.Context)
	GetOrganizationUserList(c *gin.Context)
	GetOrganizationPatientList(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	DeleteUsers(c *gin.Context)
}

type UserControllerStruct struct {
}

func NewUserController() UserControllerInterface {
	x := UserControllerStruct{}
	return &x
}

func (u *UserControllerStruct) Own(c *gin.Context) {
	staff := token.GetStaffUser(c)
	user, _ := userRepository.GetUserByID(staff.UserID)
	c.JSON(200, user)
	return
}

func (u *UserControllerStruct) GetOrganizationUsersList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userGroupID, _ := strconv.Atoi(c.Query("group"))
	q := c.Query("q")
	staff := token.GetStaffUser(c)
	organization, _ := organizationRepository.GetOrganizationByID(staff.OrganizationID)
	filter := models.UserModel{}
	isDoctor := organization.IsDoctor()
	if isDoctor {
		filter.OrganizationID = organization.ID
	}
	if userGroupID > 0 {
		filter.UserGroupID = uint64(userGroupID)
	}
	if page < 1 {
		response, _ := userRepository.GetOrganizationUserListBy(&filter, q, organization)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := userRepository.GetPaginatedOrganizationUserListBy(&filter, q, organization, page, limit)
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) GetOrganizationUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userGroupID, _ := strconv.Atoi(c.Query("group"))
	organizationID, _ := strconv.Atoi(c.Param("id"))
	q := c.Query("q")
	filter := models.UserModel{
		OrganizationID: uint64(organizationID),
	}
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
	response, _ := userRepository.GetPaginatedUserListBy(&filter, q, page, limit)
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) GetOrganizationPatientList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userGroupID, _ := strconv.Atoi(c.Query("group"))
	organizationID, _ := strconv.Atoi(c.Param("id"))
	q := c.Query("q")
	filter := models.UserModel{
		OrganizationID: uint64(organizationID),
	}
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
	response, _ := userRepository.GetPaginatedUserListBy(&filter, q, page, limit)
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := userRepository.GetUserByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, nil)
			return
		}
		c.JSON(500, err.Error())
		return
	}
	if response.City != nil {
		response.County, _ = countyRepository.GetCountyByID(response.City.CountyID)
		response.Province, _ = provinceRepository.GetProvinceByID(response.County.ProvinceID)
	}
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	staff := token.GetStaffUser(c)
	if staff.UserGroupID == 1 {
		c.JSON(403, "access denied")
	}
	response, err := userRepository.DeleteUserByID(uint64(id))
	if err != nil {
		c.JSON(500, response)
		return
	}
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) DeleteUsers(c *gin.Context) {
	var request requests.UserListDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	response, err := userRepository.DeleteUserListByID(request.IDs)
	if err != nil {
		c.JSON(500, response)
		return
	}
	c.JSON(200, nil)
	return
}

func (u *UserControllerStruct) CreateUser(c *gin.Context) {
	var request requests.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	staff := token.GetStaffUser(c)
	user, err := userRepository.CreateUser(&request, staff.UserID, staff.OrganizationID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	go user2.SendRegisterSms(user)
	c.JSON(200, &user)
	return
}

func (u *UserControllerStruct) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userRepository.GetUserByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, "not found")
			return
		} else {
			c.JSON(500, err.Error())
			return
		}
	}
	var request requests.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	tx := repository.DB.MySQL.Begin()
	err = userRepository.UpdateUser(&request)
	if err != nil {
		tx.Rollback()
		c.JSON(500, err.Error())
		return
	}
	tx.Commit()
	if user.Tel != request.Tel || strings.ToLower(user.Gender) != strings.ToLower(request.Gender) {
		go user2.SendRegisterSms(user)
	}
	c.JSON(200, true)
	return
}
