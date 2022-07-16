package models

import (
	"github.com/saeedhpro/apisimateb/constant"
	"time"
)

type OrganizationModel struct {
	ID               uint64                 `json:"id" gorm:"primarykey"`
	Name             string                 `json:"name" gorm:"name,index"`
	KnownAs          string                 `json:"known_as" gorm:"known_as"`
	ProfessionID     uint64                 `json:"profession_id" gorm:"profession_id"`
	Profession       *ProfessionModel       `json:"profession" gorm:"foreignKey:ProfessionID"`
	Logo             string                 `json:"logo" gorm:"logo"`
	Phone            string                 `json:"phone" gorm:"phone"`
	Phone1           string                 `json:"phone1" gorm:"phone1"`
	CreatedAt        *time.Time             `json:"created_at" gorm:"created_at"`
	StaffID          uint64                 `json:"staff_id" gorm:"staff_id"`
	Staff            *UserModel             `json:"staff" gorm:"foreignKey:StaffID"`
	Info             string                 `json:"info" gorm:"info"`
	CaseTypes        string                 `json:"case_types" gorm:"case_types"`
	SmsCredit        int                    `json:"sms_credit" gorm:"sms_credit"`
	SmsPrice         float64                `json:"sms_price" gorm:"sms_price"`
	SliderRndImg     string                 `json:"slider_rnd_img" gorm:"slider_rnd_img"`
	SliderImgs       int                    `json:"slider_imgs" gorm:"slider_imgs"`
	WorkHourStart    string                 `json:"work_hour_start" gorm:"work_hour_start"`
	WorkHourEnd      string                 `json:"work_hour_end" gorm:"work_hour_end"`
	Website          string                 `json:"website" gorm:"website"`
	Instagram        string                 `json:"instagram" gorm:"instagram"`
	Text1            string                 `json:"text1" gorm:"text1"`
	Image1           string                 `json:"image1" gorm:"image1"`
	Text2            string                 `json:"text2" gorm:"text2"`
	Image2           string                 `json:"image2" gorm:"image2"`
	Text3            string                 `json:"text3" gorm:"text3"`
	Image3           string                 `json:"image3" gorm:"image3"`
	Text4            string                 `json:"text4" gorm:"text4"`
	Image4           string                 `json:"image4" gorm:"image4"`
	RelOrganizations []RelOrganizationModel `json:"rel_organizations" gorm:"-"`
}

func (OrganizationModel) TableName() string {
	return "organization"
}

func (o *OrganizationModel) IsDoctor() bool {
	return o.ProfessionID != 1 && o.ProfessionID != 2 && o.ProfessionID != 3
}

func (o *OrganizationModel) IsPhotography() bool {
	return o.ProfessionID == constant.PhotographyProfession
}

func (o *OrganizationModel) IsRadiology() bool {
	return o.ProfessionID == constant.RadiologyProfession
}

func (o *OrganizationModel) IsLaboratory() bool {
	return o.ProfessionID == constant.LaboratoryProfession
}
