package fileController

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/fileRepository"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type FileControllerInterface interface {
	GetUserFileList(c *gin.Context)
	CreateFile(c *gin.Context)
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

func (f FileControllerStruct) CreateFile(c *gin.Context) {
	var request requests.FileCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, err.Error())
		return
	}
	if request.File != "" {
		name := "file"
		location := fmt.Sprintf("./res/%s", name)
		_, name, err := saveFileToDisk(location, request.File, request.Ext)
		if err != nil {
			fmt.Println(err.Error())
		}
		request.Path = name
	} else {
		c.JSON(422, "file is mandatory")
		return
	}
	staff := token.GetStaffUser(c)
	request.StaffID = staff.UserID
	request.OrganizationID = staff.OrganizationID
	response, err := fileRepository.CreateFile(&request)
	if err != nil {
		c.JSON(500, response)
		return
	}
	file, _ := fileRepository.GetFileByID(response.ID)
	organization := models.OrganizationModel{
		ID: file.OrganizationID,
	}
	_ = repository.DB.MySQL.Find(&organization, &organization).Error
	file.Organization = organization
	user := models.UserModel{
		ID: file.UserID,
	}
	_ = repository.DB.MySQL.Find(&user, &user).Error
	file.User = user
	user = models.UserModel{
		ID: file.StaffID,
	}
	_ = repository.DB.MySQL.Find(&user, &user).Error
	file.Staff = user
	file.Path = fmt.Sprintf("http://%s/file/%s/1.%s", c.Request.Host, file.Path, file.Ext)
	c.JSON(200, response)
	return
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

func saveFileToDisk(location string, data string, ext string) (string, string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", "", fmt.Errorf("invalid image")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		log.Println("errpeed")
		return "", "", err
	}
	name := ""
	var names []string
	files, err := ioutil.ReadDir(location)
	for i := 0; i < len(files); i++ {
		names = append(names, files[i].Name())
	}
	for {
		name = helpers.RandomString(8)
		fmt.Println(fmt.Sprintf("%s", name), "location")
		if !helpers.ItemExists(names, fmt.Sprintf("%s", name)) {
			err = os.MkdirAll(fmt.Sprintf("%s/%s", location, name), os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
			}
			break
		}
	}
	fileName := fmt.Sprintf("%s/%s/1.%s", location, name, ext)
	err = ioutil.WriteFile(fileName, buff.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error(), "cf")
		return "", "", fmt.Errorf("cant save file")
	}
	return fileName, name, err
}
