package medicalHistoryController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/repository/medicalHistoryRepository"
	"gorm.io/gorm"
	"strconv"
)

type MedicalHistoryControllerInterface interface {
	GetUserMedicalHistory(c *gin.Context)
	CreateUserMedicalHistory(c *gin.Context)
}

type MedicalHistoryControllerStruct struct {
}

func NewMedicalHistoryController() MedicalHistoryControllerInterface {
	x := MedicalHistoryControllerStruct{}
	return &x
}

func (m *MedicalHistoryControllerStruct) GetUserMedicalHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, "incorrect id")
		return
	}
	medical, err := medicalHistoryRepository.GetUserMedicalHistory(uint64(id))
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, medical)
	return
}

func (m *MedicalHistoryControllerStruct) CreateUserMedicalHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, "incorrect id")
		return
	}
	var request requests.CreateUserMedicalHistoryRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, err.Error())
		return
	}
	medical, err := medicalHistoryRepository.GetUserMedicalHistory(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = medicalHistoryRepository.CreateUserMedicalHistory(&request)
			if err != nil {
				c.JSON(500, err.Error())
				return
			}
			c.JSON(200, "done")
			return
		} else {
			if err != nil {
				c.JSON(500, err.Error())
				return
			}
		}
	} else {
		err = medicalHistoryRepository.UpdateUserMedicalHistory(&request)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
	}
	c.JSON(200, medical)
	return
}
