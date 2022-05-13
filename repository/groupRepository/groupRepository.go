package groupRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetGroupListBy(conditions *models.UserGroupModel, q string) ([]models.UserGroupModel, error) {
	groups := []models.UserGroupModel{}
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&groups, &conditions).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func GetPaginatedGroupListBy(conditions *models.UserGroupModel, q string, page int, limit int) (pagination.Pagination, error) {
	groups := []models.UserGroupModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	query.Find(&groups, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate))
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&groups, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = groups
	return paginate, nil
}
