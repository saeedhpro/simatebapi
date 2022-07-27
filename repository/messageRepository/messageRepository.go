package messageRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"time"
)

func GetMessageListBy(conditions *models.SmsModel, q string) ([]models.SmsModel, error) {
	messages := []models.SmsModel{}
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("msg LIKE ?", "%"+q+"%")
	}
	err := query.
		Preload("User").
		Preload("Staff").
		Preload("Organization").
		Find(&messages, &conditions).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetMessageListByIds(conditions *models.SmsModel, q string, ids []uint64) ([]models.SmsModel, error) {
	messages := []models.SmsModel{}
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("msg LIKE ?", "%"+q+"%")
	}
	query = query.
		Where(ids)
	err := query.
		Preload("User").
		Preload("Staff").
		Preload("Organization").
		Find(&messages, &conditions).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetPaginatedMessageListBy(conditions *models.SmsModel, q string, page int, limit int) (pagination.Pagination, error) {
	messages := []models.SmsModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("msg LIKE ?", "%"+q+"%")
	}
	query.Find(&messages, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate))
	if q != "" {
		query = query.Where("msg LIKE ?", "%"+q+"%")
	}
	err := query.
		Preload("User").
		Preload("Staff").
		Preload("Organization").
		Find(&messages, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = messages
	return paginate, nil
}

func DeleteMessages(ids []uint64) error {
	return repository.DB.MySQL.Delete(&models.SmsModel{}, ids).Error
}

func SendSMS(request *requests.SendSMSRequest, staffID uint64, organizationID uint64, send bool) error {
	var numbers []string
	var smsList []models.SmsModel
	for i := 0; i < len(request.Numbers); i++ {
		n := helpers.NormalizePhoneNumber(request.Numbers[i])
		if n != "" {
			user, _ := userRepository.GetUserBy(&models.UserModel{
				Tel: request.Numbers[i],
			})
			if user == nil {
				user, _ = userRepository.GetUserBy(&models.UserModel{
					Tel: fmt.Sprintf("0%s", request.Numbers[i][3:]),
				})
			}
			if user != nil {
				now := time.Now()
				sms := models.SmsModel{
					UserID:         user.ID,
					OrganizationID: organizationID,
					StaffID:        staffID,
					Incoming:       true,
					Msg:            request.Msg,
					Number:         user.Tel,
					Sent:           send,
					Created:        &now,
				}
				smsList = append(smsList, sms)
				numbers = append(numbers, n)
			}
		}
	}
	for i := 0; i < len(smsList); i++ {
		_ = repository.DB.MySQL.Create(&smsList[i]).Error
	}
	return nil
}
