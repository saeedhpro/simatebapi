package messageRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
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
