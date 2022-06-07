package organizationRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetOrganizationByID(ID uint64) (*models.OrganizationModel, error) {
	var organization models.OrganizationModel
	organization.ID = ID
	err := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff").
		First(&organization, &organization).Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func GetOrganizationByType(t string) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	var organization models.OrganizationModel
	if t == "radiology" {
		organization.ProfessionID = 3
	} else if t == "photography" {
		organization.ProfessionID = 1
	}
	err := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff").
		Find(&organizations, &organization).
		Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetOrganizationListBy(conditions *models.OrganizationModel, q string) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	query := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff")
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&organizations, &conditions).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetPaginatedOrganizationListBy(conditions *models.OrganizationModel, q string, page int, limit int) (pagination.Pagination, error) {
	organizations := []models.OrganizationModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	query.Find(&organizations, &conditions).Count(&count)
	query = repository.DB.MySQL.
		Scopes(pagination.PaginateScope(count, &paginate)).
		Preload("Profession").
		Preload("Staff")
	if q != "" {
		query = query.Where("name LIKE ?", "%"+q+"%")
	}
	err := query.Find(&organizations, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = organizations
	return paginate, nil
}
