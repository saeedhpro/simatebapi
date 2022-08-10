package paymentRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/domain/responses"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
	"time"
)

func GetPaymentsListBy(conditions *models.PaymentModel) ([]models.PaymentModel, error) {
	payments := []models.PaymentModel{}
	err := repository.DB.MySQL.Find(&payments, &conditions).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func GetPaginatedPaymentsListBy(conditions *models.PaymentModel, page int, limit int) (pagination.Pagination, error) {
	payments := []models.PaymentModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	query.Find(&payments, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate)).
		Preload("Staff").
		Preload("User")
	err := query.Find(&payments, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = payments
	return paginate, nil
}

func GetUserPaymentsTotal(userId uint64) (*responses.PaymentTotalResponse, error) {
	response := responses.PaymentTotalResponse{}
	var due float64 = 0
	query := repository.DB.MySQL.
		Model(models.PaymentModel{}).
		Preload("Organization").
		Preload("Staff").
		Preload("User")
	err := query.Where("income = ?", 1).
		Where("paid_for = ?", 1).
		Where("user_id = ?", userId).
		Select("sum(amount) as due").
		Scan(&due).Error
	if err != nil {
		return &response, err
	}
	response.DueTotal = due
	query = repository.DB.MySQL.
		Model(models.PaymentModel{}).
		Preload("Organization").
		Preload("Staff").
		Preload("User")
	err = query.Where("income = ?", 1).
		Where("user_id = ?", userId).
		Select("sum(amount) as due").
		Scan(&due).Error
	if err != nil {
		return &response, err
	}
	response.Total = due
	return &response, nil
}

func CreatePayment(request *requests.CreatePaymentRequest, user *models.UserModel) (*models.PaymentModel, error) {
	payment := models.PaymentModel{
		StaffID:   user.ID,
		UserID:    *request.UserID,
		Income:    1,
		Info:      request.Info,
		Paytype:   request.Paytype,
		PaidTo:    request.PaidTo,
		PaidFor:   request.PaidFor,
		CheckBank: request.CheckBank,
		CheckNum:  request.CheckNum,
		TraceCode: request.TraceCode,
		Amount:    request.Amount,
	}
	if request.Created != "" {
		Created, err := time.Parse("2006-01-02 00:00:00", request.Created)
		if err == nil {
			payment.Created = &Created
		} else {
			fmt.Println(err.Error())
			now := time.Now()
			payment.Created = &now
		}
	}
	if request.CheckDate != "" {
		CheckDate, err := time.Parse("2006-01-02 00:00:00", request.CheckDate)
		if err == nil {
			payment.CheckDate = &CheckDate
		} else {
			fmt.Println(err.Error())
			now := time.Now()
			payment.CheckDate = &now
		}
	}
	err := repository.DB.MySQL.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func DeletePayments(ids []uint64) error {
	return repository.DB.MySQL.Delete(&models.PaymentModel{}, ids).Error
}
