package models

type CaseType struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Name           string             `json:"name" gorm:"name,index"`
	Limitation     uint64             `json:"limitation" gorm:"limitation"`
	Duration       uint64             `json:"duration" gorm:"duration"`
	IsLimited      bool               `json:"is_limited" gorm:"is_limited"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id"`
	Organization   *OrganizationModel `json:"organization" gorm:"ForeignKey:OrganizationID"`
}

func (CaseType) TableName() string {
	return "case_type"
}
