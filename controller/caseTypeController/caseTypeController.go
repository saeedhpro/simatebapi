package caseTypeController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/caseTypeRepository"
	"gorm.io/gorm"
	"strconv"
)

type CaseTypeControllerInterface interface {
	Get(c *gin.Context)
	GetOrganizationCaseTypeList(c *gin.Context)
	CreateCaseType(c *gin.Context)
	UpdateCaseType(c *gin.Context)
	DeleteCaseType(c *gin.Context)
}

type CaseTypeControllerStruct struct {
}

func NewCaseTypeController() CaseTypeControllerInterface {
	x := CaseTypeControllerStruct{}
	return &x
}

func (ct *CaseTypeControllerStruct) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := caseTypeRepository.GetCaseTypeByID(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, response)
	return
}

func (ct *CaseTypeControllerStruct) GetOrganizationCaseTypeList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	organizationID := token.GetStaffUser(c).OrganizationID
	filter := models.CaseType{OrganizationID: organizationID}
	if page < 1 {
		response, _ := caseTypeRepository.GetCaseTypeListBy(&filter)
		c.JSON(200, response)
		return
	}
	if limit < 1 {
		limit = 10
	}
	response, _ := caseTypeRepository.GetPaginatedCaseTypeListBy(&filter, page, limit)
	c.JSON(200, response)
	return
}

func (ct *CaseTypeControllerStruct) CreateCaseType(c *gin.Context) {
	var request requests.CreateCaseTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, err.Error())
		return
	}
	err := caseTypeRepository.CreateCaseType(&request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "done")
	return
}

func (ct *CaseTypeControllerStruct) UpdateCaseType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var request requests.CreateCaseTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, err.Error())
		return
	}
	err = caseTypeRepository.UpdateCaseType(uint64(id), &request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "done")
	return
}

func (ct *CaseTypeControllerStruct) DeleteCaseType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	caseType := models.CaseType{
		ID: uint64(id),
	}
	err = caseTypeRepository.DeleteCaseType(&caseType)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "")
	return
}
