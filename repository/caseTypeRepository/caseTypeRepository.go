package caseTypeRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetCaseTypeBy(conditions *models.CaseType) (*models.CaseType, error) {
	var caseType models.CaseType
	err := repository.DB.MySQL.Preload("Organization").First(&caseType, &conditions).Error
	if err != nil {
		return nil, err
	}
	return &caseType, nil
}

func GetCaseTypeByID(id uint64) (*models.CaseType, error) {
	caseType := models.CaseType{
		ID: id,
	}
	err := repository.DB.MySQL.Preload("Organization").First(&caseType, &caseType).Error
	if err != nil {
		return nil, err
	}
	return &caseType, nil
}

func GetCaseTypeListBy(conditions *models.CaseType) ([]models.CaseType, error) {
	caseTypes := []models.CaseType{}
	err := repository.DB.MySQL.Preload("Organization").Find(&caseTypes, &conditions).Error
	if err != nil {
		return caseTypes, err
	}
	return caseTypes, nil
}

func GetPaginatedCaseTypeListBy(conditions *models.CaseType, page int, limit int) (pagination.Pagination, error) {
	caseTypes := []models.CaseType{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	repository.DB.MySQL.Find(&caseTypes, &conditions).Count(&count)
	err := repository.DB.MySQL.Scopes(pagination.PaginateScope(count, &paginate)).Preload("Organization").Find(&caseTypes, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = caseTypes
	return paginate, nil
}

func CreateCaseType(request *requests.CreateCaseTypeRequest) error {
	var IsLimited = 0
	if request.IsLimited {
		IsLimited = 1
	}
	caseType := models.CaseType{
		OrganizationID: request.OrganizationID,
		Name:           request.Name,
		Duration:       request.Duration,
		IsLimited:      IsLimited,
		Limitation:     request.Limitation,
	}
	err := repository.DB.MySQL.Create(caseType).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCaseType(id uint64, request *requests.CreateCaseTypeRequest) error {
	var IsLimited = 0
	if request.IsLimited {
		IsLimited = 1
	}
	caseType := models.CaseType{
		ID:             id,
		IsLimited:      IsLimited,
		Limitation:     request.Limitation,
		Duration:       request.Duration,
		OrganizationID: request.OrganizationID,
		Name:           request.Name,
	}
	err := repository.DB.MySQL.
		Model(&caseType).
		Updates(map[string]interface{}{
			"IsLimited":      IsLimited,
			"Limitation":     request.Limitation,
			"Duration":       request.Duration,
			"OrganizationID": request.OrganizationID,
			"Name":           request.Name,
		}).
		Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func DeleteCaseType(caseType *models.CaseType) error {
	return repository.DB.MySQL.Delete(caseType).Error
}
