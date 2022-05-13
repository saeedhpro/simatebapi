package models

import "time"

type SmsModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Sent           bool               `json:"sent" gorm:"sent"`
	Incoming       bool               `json:"incoming" gorm:"incoming"`
	Msg            string             `json:"msg" gorm:"msg"`
	Number         string             `json:"number" gorm:"number"`
	Amount         float64            `json:"amount" gorm:"amount"`
	UserID         uint64             `json:"user_id" gorm:"user_id"`
	User           *UserModel         `json:"user" gorm:"foreignkey:UserID"`
	StaffID        uint64             `json:"staff_id" gorm:"staff_id"`
	Staff          *UserModel         `json:"staff" gorm:"foreignkey:StaffID"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id"`
	Organization   *OrganizationModel `json:"organization" gorm:"foreignkey:OrganizationID"`
	Created        *time.Time         `json:"created" gorm:"created"`
}

func (SmsModel) TableName() string {
	return "sms"
}
