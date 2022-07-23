package messageController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"github.com/saeedhpro/apisimateb/repository/messageRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"strconv"
)

type MessageControllerInterface interface {
	GetOrganizationMessages(c *gin.Context)
	DeleteMessages(c *gin.Context)
	SendSms(c *gin.Context)
}

type MessageControllerStrunct struct {
}

func NewMessageController() MessageControllerInterface {
	x := MessageControllerStrunct{}
	return &x
}

func (m *MessageControllerStrunct) GetOrganizationMessages(c *gin.Context) {
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	filter := models.SmsModel{
		OrganizationID: staff.OrganizationID,
	}
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

func (m *MessageControllerStrunct) DeleteMessages(c *gin.Context) {

	var request requests.DeleteMultipleItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, "parse failed")
		return
	}
	messages, err := messageRepository.GetMessageListByIds(&models.SmsModel{}, "", request.Ids)
	if err != nil {
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(500, "get messages")
			return
		}
	}
	staff, _ := userRepository.GetUserByID(token.GetStaffUser(c).UserID)
	for i := 0; i < len(messages); i++ {
		if messages[i].OrganizationID != staff.OrganizationID {
			c.JSON(403, "Access Denied!")
			return
		}
	}
	err = messageRepository.DeleteMessages(request.Ids)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
	return
}

func (m *MessageControllerStrunct) SendSms(c *gin.Context) {
	staff := token.GetStaffUser(c)
	var request requests.SendSMSRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, err.Error())
		return
	}
	if request.Type == 1 {
		request.Numbers = append(request.Numbers, request.PhoneNumber)
	} else if request.Type == 3 {
		allNumbers, _ := userRepository.GetNumberListByOrganizationID(0)
		request.Numbers = allNumbers
	}
	ok, _, err := request.SendSMS()
	if err != nil {
		fmt.Println(err.Error())
		ok = false
	}
	organizationID := staff.OrganizationID
	if request.OrganizationID != 0 {
		organizationID = request.OrganizationID
	}
	err = messageRepository.SendSMS(&request, staff.UserID, organizationID, ok)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, true)
}
