package relOrganizationRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetOrganizationRelListByID(ID uint64) ([]models.RelOrganizationModel, error) {
	organizations := []models.RelOrganizationModel{}
	organization := models.RelOrganizationModel{
		OrganizationID: ID,
	}
	err := repository.DB.MySQL.
		Preload("Profession").
		Preload("Organization").
		Preload("RelOrganization").
		Find(&organizations, &organization).
		Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetRelOrganization(organizationID uint64, relOrganizationID uint64, professionID uint64) (*models.RelOrganizationModel, error) {
	rel := models.RelOrganizationModel{
		OrganizationID:    organizationID,
		RelOrganizationID: relOrganizationID,
		ProfessionID:      professionID,
	}
	err := repository.DB.MySQL.
		Preload("Profession").
		Preload("Organization").
		Preload("RelOrganization").
		Find(&rel, &rel).
		Error
	if err != nil {
		return nil, err
	}
	return &rel, nil
}

func DeleteRelOrganization(organizationID uint64, relOrganizationID uint64, professionID uint64) (bool, error) {
	rel := models.RelOrganizationModel{
		OrganizationID:    organizationID,
		RelOrganizationID: relOrganizationID,
		ProfessionID:      professionID,
	}
	err := repository.DB.MySQL.
		Delete(&rel, &rel).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateRelOrganization(organizationID uint64, relOrganizationID uint64, professionID uint64) (bool, error) {
	rel := models.RelOrganizationModel{
		OrganizationID:    organizationID,
		RelOrganizationID: relOrganizationID,
		ProfessionID:      professionID,
	}
	err := repository.DB.MySQL.
		Create(&rel).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}
