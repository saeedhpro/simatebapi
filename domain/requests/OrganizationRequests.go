package requests

import (
	"github.com/saeedhpro/apisimateb/domain/models"
)

type CreateOrganizationRequest struct {
	ID               uint64                        `json:"id"`
	Name             string                        `json:"name"`
	New              *string                       `json:"new"`
	ProfessionID     uint64                        `json:"profession_id"`
	StaffID          uint64                        `json:"staff_id"`
	Website          string                        `json:"website"`
	Phone            string                        `json:"phone"`
	Phone1           string                        `json:"phone1"`
	Instagram        string                        `json:"instagram"`
	SmsPrice         float64                       `json:"sms_price"`
	SmsCredit        int                           `json:"sms_credit"`
	CaseTypes        string                        `json:"case_types"`
	KnownAs          string                        `json:"known_as"`
	CreatedAt        *string                       `json:"created_at"`
	Logo             string                        `json:"logo"`
	Info             string                        `json:"info"`
	SliderRndImg     string                        `json:"slider_rnd_img"`
	SliderImgs       int                           `json:"slider_imgs"`
	WorkHourStart    string                        `json:"work_hour_start"`
	WorkHourEnd      string                        `json:"work_hour_end"`
	AboutUsHtml      string                        `json:"about_us_html"`
	Text1            string                        `json:"text1"`
	Image1           string                        `json:"image1"`
	Text2            string                        `json:"text2"`
	Image2           string                        `json:"image2"`
	Text3            string                        `json:"text3"`
	Image3           string                        `json:"image3"`
	Text4            string                        `json:"text4"`
	Image4           string                        `json:"image4"`
	RelOrganizations []models.RelOrganizationModel `json:"rel_organizations"`
}

type UpdateOrganizationAbout struct {
	Text1  string `json:"text1"`
	Image1 string `json:"image1"`
	Text2  string `json:"text2"`
	Image2 string `json:"image2"`
	Text3  string `json:"text3"`
	Image3 string `json:"image3"`
	Text4  string `json:"text4"`
	Image4 string `json:"image4"`
}
