package models

import "time"

type CreditModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Verified       bool               `json:"verified" gorm:"verified"`
	TraceCode      string             `json:"trace_code" gorm:"trace_code"`
	RefNum         string             `json:"ref_num" gorm:"ref_num"`
	Comment        string             `json:"comment" gorm:"comment"`
	PaidFor        string             `json:"paid_for" gorm:"paid_for"`
	Amount         float64            `json:"amount" gorm:"amount"`
	StaffID        uint64             `json:"staff_id" gorm:"staff_id"`
	Staff          *UserModel         `json:"staff" gorm:"foreignkey:StaffID"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id"`
	Organization   *OrganizationModel `json:"organization" gorm:"foreignkey:OrganizationID"`
	CreatedAt      *time.Time         `json:"created_at" gorm:"created_at"`
}

func (CreditModel) TableName() string {
	return "credit"
}
