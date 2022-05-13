package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type UserControllerInterface interface {
	Own(c *gin.Context)
	GetOrganizationUsersList(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
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
	filter := models.UserModel{
		OrganizationID: staff.OrganizationID,
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
	c.JSON(200, response)
	return
}

func (u *UserControllerStruct) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	staff := token.GetStaffUser(c)
	if staff.UserGroupID == 1 {

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
	if err !=nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, &user)
	return
}
