package paymentController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/repository/paymentRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"strconv"
)

type PaymentControllerInterface interface {
	GetUserPayments(c *gin.Context)
	GetUserPaymentsTotal(c *gin.Context)
	CreatePayment(c *gin.Context)
	DeletePayments(c *gin.Context)
}

type PaymentControllerStruct struct {
}

func NewPaymentController() PaymentControllerInterface {
	x := PaymentControllerStruct{}
	return &x
}

func (p *PaymentControllerStruct) GetUserPayments(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userID, _ := strconv.Atoi(c.Param("id"))
	filter := models.PaymentModel{
		UserID: uint64(userID),
	}
	if page < 1 {
		response, _ := paymentRepository.GetPaymentsListBy(&filter)
		c.JSON(200, response)
		return
	}
	response, _ := paymentRepository.GetPaginatedPaymentsListBy(&filter, page, limit)
	c.JSON(200, response)
	return
}

func (p *PaymentControllerStruct) GetUserPaymentsTotal(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	response, _ := paymentRepository.GetUserPaymentsTotal(uint64(userID))
	c.JSON(200, response)
	return
}

func (p *PaymentControllerStruct) CreatePayment(c *gin.Context) {
	var request requests.CreatePaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error(), "bind")
		c.JSON(500, err.Error())
		return
	}
	user, _ := userRepository.GetUserByID(*request.UserID)
	payment, err := paymentRepository.CreatePayment(&request, user)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, &payment)
	return
}

func (p *PaymentControllerStruct) DeletePayments(c *gin.Context) {
	var request requests.DeleteMultipleItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(500, "parse failed")
		return
	}
	err := paymentRepository.DeletePayments(request.Ids)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
	return
}
