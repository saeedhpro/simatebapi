package fileController

import (
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/fileRepository"
	"strconv"
)

type FileControllerInterface interface {
	GetUserFileList(c *gin.Context)
	DeleteFile(c *gin.Context)
}

type FileControllerStruct struct {

}

func NewFileController() FileControllerInterface {
	x := FileControllerStruct{}
	return &x
}

func (f FileControllerStruct) GetUserFileList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	staff := token.GetStaffUser(c)
	conditions := models.FileModel{UserID: uint64(userID), OrganizationID: staff.OrganizationID}
	response, _ := fileRepository.GetFilesListBy(c, &conditions)
	c.JSON(200, response)
}

func (f FileControllerStruct) DeleteFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	conditions := models.FileModel{ID: uint64(id)}
	response, err := fileRepository.DeleteFileBy(&conditions)
	if err != nil {
		c.JSON(500, response)
		return
	}
	c.JSON(200, response)
	return
}

