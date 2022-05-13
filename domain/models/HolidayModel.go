package models

import "time"

type HolidayModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Hdate          time.Time          `json:"hdate" gorm:"hdate"`
	Title          string             `json:"title" gorm:"title"`
	OrganizationID *uint64            `json:"organization_id" gorm:"organization_id,index"`
	Organization   *OrganizationModel `json:"organization" gorm:"ForeignKey:OrganizationID"`
}

func (HolidayModel) TableName() string {
	return "holiday"
}
