package models

type RelOrganizationModel struct {
	ID                uint64             `json:"id" gorm:"primarykey"`
	ProfessionID      uint64             `json:"profession_id" gorm:"profession_id"`
	Profession        *ProfessionModel   `json:"profession" gorm:"foreignkey:ProfessionID"`
	OrganizationID    uint64             `json:"organization_id" gorm:"organization_id"`
	Organization      *OrganizationModel `json:"organization" gorm:"foreignkey:OrganizationID"`
	RelOrganizationID uint64             `json:"rel_organization_id" gorm:"rel_organization_id"`
	RelOrganization   *OrganizationModel `json:"rel_organization" gorm:"foreignkey:RelOrganizationID"`
}

func (RelOrganizationModel) TableName() string {
	return "rel_organization"
}
