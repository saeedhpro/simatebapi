package models

import "time"

type VipScheduleModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Count          uint               `json:"count" gorm:"count"`
	Site           int                `json:"site" gorm:"site"`
	App            int                `json:"app" gorm:"app"`
	StartAt        *time.Time         `json:"start_at" gorm:"start_at"`
	EndAt          *time.Time         `json:"end_at" gorm:"end_at"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id"`
	Organization   *OrganizationModel `json:"organization" gorm:"foreignkey:OrganizationID"`
}

func (VipScheduleModel) TableName() string {
	return "vip_schedule"
}
