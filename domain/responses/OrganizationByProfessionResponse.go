package responses

import "github.com/saeedhpro/apisimateb/domain/models"

type OrganizationByProfessionResponse struct {
	Radiologies   []models.OrganizationModel `json:"radiologies"`
	Photographies []models.OrganizationModel `json:"photographies"`
	Laboratories  []models.OrganizationModel `json:"laboratories"`
	Doctors       []models.OrganizationModel `json:"doctors"`
}
