package organizationRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
	"github.com/saeedhpro/apisimateb/repository/relOrganizationRepository"
	"time"
)

func GetOrganizationByID(ID uint64) (*models.OrganizationModel, error) {
	var organization models.OrganizationModel
	organization.ID = ID
	err := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff").
		First(&organization, &organization).Error
	organization.RelOrganizations, _ = relOrganizationRepository.GetOrganizationRelListByID(organization.ID)
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

func GetOnlyDoctorOrganizationList(conditions *models.OrganizationModel, organizationID uint64) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	query := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff")
	query = query.Not("profession_id", []uint64{1, 2, 3})
	query = query.Not("id", []uint64{organizationID})
	err := query.Find(&organizations, &conditions).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetOnlyRadiologyOrganizationList(conditions *models.OrganizationModel, organizationID uint64) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	query := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff")
	query = query.Not("id", []uint64{organizationID})
	err := query.Find(&organizations, &conditions).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetOnlyLaboratoryOrganizationList(conditions *models.OrganizationModel, organizationID uint64) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	query := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff")
	query = query.Not("id", []uint64{organizationID})
	err := query.Find(&organizations, &conditions).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetOnlyPhotographyOrganizationList(conditions *models.OrganizationModel, organizationID uint64) ([]models.OrganizationModel, error) {
	organizations := []models.OrganizationModel{}
	query := repository.DB.MySQL.
		Preload("Profession").
		Preload("Staff")
	query = query.Not("id", []uint64{organizationID})
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

func UpdateOrganization(request *requests.CreateOrganizationRequest) error {
	organization := models.OrganizationModel{
		ID:            request.ID,
		Name:          request.Name,
		KnownAs:       request.KnownAs,
		ProfessionID:  request.ProfessionID,
		Phone:         request.Phone,
		Info:          request.Info,
		CaseTypes:     request.CaseTypes,
		SmsCredit:     request.SmsCredit,
		SmsPrice:      request.SmsPrice,
		SliderRndImg:  request.SliderRndImg,
		SliderImgs:    request.SliderImgs,
		WorkHourStart: request.WorkHourStart,
		WorkHourEnd:   request.WorkHourEnd,
		Website:       request.Website,
		Instagram:     request.Instagram,
		Text1:         request.Text1,
		Image1:        request.Image1,
		Text2:         request.Text2,
		Image2:        request.Image2,
		Text3:         request.Text3,
		Image3:        request.Image3,
	}
	if request.CreatedAt != nil {
		created := time.Time{}
		cr := *request.CreatedAt
		t, err := time.Parse(time.RFC3339, cr)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			created = t
		}
		organization.CreatedAt = &created
	}
	if request.Logo != "" {
		organization.Logo = request.Logo
	}
	err := repository.DB.MySQL.
		Model(&organization).
		Where("id = ?", &organization.ID).
		Updates(organization).
		Error
	if err != nil {
		fmt.Println(err.Error(), "err")
		return err
	}
	return nil
}

func UpdateOrganizationAbout(id uint64, request *requests.UpdateOrganizationAboutNames) error {
	organization := models.OrganizationModel{
		ID:     id,
		Text1:  request.Text1,
		Image1: request.Image1,
		Text2:  request.Text2,
		Image2: request.Image2,
		Text3:  request.Text3,
		Image3: request.Image3,
		Text4:  request.Text4,
		Image4: request.Image4,
	}
	err := repository.DB.MySQL.
		Model(&organization).
		Where("id = ?", &organization.ID).
		Updates(organization).
		Error
	if err != nil {
		fmt.Println(err.Error(), "err")
		return err
	}
	return nil
}

func CreateOrganization(request *requests.CreateOrganizationRequest) (*models.OrganizationModel, error) {
	created := time.Now()
	organization := models.OrganizationModel{
		Name:          request.Name,
		KnownAs:       request.KnownAs,
		ProfessionID:  request.ProfessionID,
		StaffID:       request.StaffID,
		Phone:         request.Phone,
		CreatedAt:     &created,
		Info:          request.Info,
		CaseTypes:     request.CaseTypes,
		SmsCredit:     request.SmsCredit,
		SmsPrice:      request.SmsPrice,
		SliderRndImg:  request.SliderRndImg,
		SliderImgs:    request.SliderImgs,
		WorkHourStart: request.WorkHourStart,
		WorkHourEnd:   request.WorkHourEnd,
		Website:       request.Website,
		Instagram:     request.Instagram,
		Text1:         request.Text1,
		Image1:        request.Image1,
		Text2:         request.Text2,
		Image2:        request.Image2,
		Text3:         request.Text3,
		Image3:        request.Image3,
	}
	err := repository.DB.MySQL.Create(&organization).
		Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func UpdateRelOrganizations(relList []models.RelOrganizationModel, newRelList []models.RelOrganizationModel) error {
	removed := differenceItems(relList, newRelList)
	added := differenceItems(newRelList, relList)
	for i := 0; i < len(removed); i++ {
		_, err := relOrganizationRepository.DeleteRelOrganization(removed[i].OrganizationID, removed[i].RelOrganizationID, removed[i].ProfessionID)
		if err != nil {
			fmt.Println(err.Error(), "del")
		}
	}
	for i := 0; i < len(added); i++ {
		_, err := relOrganizationRepository.CreateRelOrganization(added[i].OrganizationID, added[i].RelOrganizationID, added[i].ProfessionID)
		if err != nil {
			fmt.Println(err.Error(), "create")
		}
	}
	return nil
}

func differenceItems(firstList []models.RelOrganizationModel, secondList []models.RelOrganizationModel) []models.RelOrganizationModel {
	list := []models.RelOrganizationModel{}
	for i := 0; i < len(firstList); i++ {
		found := false
		for j := 0; j < len(secondList); j++ {
			if firstList[i].RelOrganizationID == secondList[j].RelOrganizationID {
				found = true
				break
			}
		}
		if !found {
			list = append(list, firstList[i])
		}
	}
	return list
}
