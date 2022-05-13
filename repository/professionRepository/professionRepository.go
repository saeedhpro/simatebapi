package professionRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetProfessionListBy(conditions *models.ProfessionModel, q string) ([]models.ProfessionModel, error) {
	professions := []models.ProfessionModel{}
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&professions, &conditions).Error
	if err != nil {
		return nil, err
	}
	return professions, nil
}

func GetPaginatedProfessionListBy(conditions *models.ProfessionModel, q string, page int, limit int) (pagination.Pagination, error) {
	professions := []models.ProfessionModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	query.Find(&professions, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate))
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&professions, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = professions
	return paginate, nil
}
