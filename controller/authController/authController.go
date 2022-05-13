package authController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/constant"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"gorm.io/gorm"
	"log"
	"time"
)

type AuthControllerInterface interface {
	Login(c *gin.Context)
}

type AuthControllerStruct struct {
}

func NewAuthController() AuthControllerInterface {
	x := &AuthControllerStruct{}
	return x
}

func (auth AuthControllerStruct) Login(c *gin.Context) {
	var request requests.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error(), "bind")
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := userRepository.GetUserBy(&models.UserModel{Tel: request.Tel})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, user)
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if helpers.PasswordVerify(request.Password, user.Pass) != true {
		c.JSON(422, gin.H{
			"message": constant.InvalidPassword,
		})
		return
	}
	repository.DB.MySQL.Model(&user).Update("last_login", time.Now())
	response, _ := token.GenerateToken(user)
	c.JSON(200, response)
	return
}
