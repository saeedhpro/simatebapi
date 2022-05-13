package models

import "time"

type OperationModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Info           string             `json:"info" gorm:"info"`
	CaseType       string             `json:"case_type" gorm:"case_type"`
	Prescription   string             `json:"prescription" gorm:"prescription"`
	Code           string             `json:"code" gorm:"code"`
	Status         uint8              `json:"status" gorm:"status"`
	UserID         uint64             `json:"user_id" gorm:"user_id"`
	User           *UserModel         `json:"user" gorm:"foreignkey:UserID"`
	StaffID        uint64             `json:"staff_id" gorm:"staff_id"`
	Staff          *UserModel         `json:"staff" gorm:"foreignkey:StaffID"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id"`
	Organization   *OrganizationModel `json:"organization" gorm:"foreignkey:OrganizationID"`
	CreatedAt      *time.Time         `json:"created_at" gorm:"created_at"`
	UpdatedAt      *time.Time         `json:"updated_at" gorm:"updated_at"`
}

func (OperationModel) TableName() string {
	return "operation"
}
