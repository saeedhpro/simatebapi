package holidayRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetHolidayBy(conditions *models.HolidayModel) (*models.HolidayModel, error) {
	var holiday models.HolidayModel
	err := repository.DB.MySQL.Preload("Organization").First(&holiday, &conditions).Error
	if err != nil {
		return nil, err
	}
	return &holiday, nil
}

func GetHolidayListBy(conditions *models.HolidayModel, q string, startDate string, endDate string) ([]models.HolidayModel, error) {
	holidays := []models.HolidayModel{}
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	query = query.
		Preload("Organization").
		Order("hdate desc")
	if startDate != "" && endDate != "" {
		query = query.
			Where("hdate between ? and ?", startDate, endDate)
	}
	err := query.
		Find(&holidays, &conditions).Error
	if err != nil {
		fmt.Println(err.Error())
		return holidays, err
	}
	return holidays, nil
}

func GetPaginatedHolidayListBy(conditions *models.HolidayModel, q string, startDate string, endDate string, page int, limit int) (pagination.Pagination, error) {
	holidays := []models.HolidayModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	query.
		Preload("Organization").
		Where("hdate between ? and ?", startDate, endDate).
		Find(&holidays, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate))
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.
		Where("hdate between ? and ?", startDate, endDate).
		Preload("Organization").
		Order("hdate desc").
		Find(&holidays, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = holidays
	return paginate, nil
}

func CreateHoliday(holiday *models.HolidayModel) error {
	return repository.DB.MySQL.Create(holiday).Error
}

func UpdateHoliday(holiday *models.HolidayModel) error {
	return repository.DB.MySQL.Model(holiday).Updates(models.HolidayModel{
		Title:          holiday.Title,
		Hdate:          holiday.Hdate,
		OrganizationID: holiday.OrganizationID,
	}).Error
}

func DeleteHoliday(holiday *models.HolidayModel) error {
	return repository.DB.MySQL.Delete(holiday).Error
}
