package scheduleRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetScheduleListBy(conditions *models.VipScheduleModel, startDate string, endDate string) ([]models.VipScheduleModel, error) {
	schedules := []models.VipScheduleModel{}
	query := repository.DB.MySQL
	query = query.
		Preload("Organization")
	if startDate != "" {
		query = query.
			Where("start_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.
			Where("end_at <= ?", endDate)
	}
	err := query.
		Find(&schedules, &conditions).Error
	if err != nil {
		fmt.Println(err.Error())
		return schedules, err
	}
	return schedules, nil
}

func GetPaginatedScheduleListBy(conditions *models.VipScheduleModel, startDate string, endDate string, page int, limit int) (pagination.Pagination, error) {
	schedules := []models.VipScheduleModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	query = query.
		Preload("Organization")
	if startDate != "" {
		query = query.
			Where("start_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.
			Where("end_at <= ?", endDate)
	}
	query.
		Find(&schedules, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate))

	query = query.
		Preload("Organization")
	if startDate != "" {
		query = query.
			Where("start_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.
			Where("end_at <= ?", endDate)
	}
	err := query.
		Find(&schedules, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = schedules
	return paginate, nil
}
