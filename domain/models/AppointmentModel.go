package models

import (
	"time"
)

type AppointmentModel struct {
	ID                   uint64             `json:"id" gorm:"primarykey"`
	UserID               uint64             `json:"user_id" gorm:"user_id"`
	User                 *UserModel         `json:"user" gorm:"ForeignKey:UserID"`
	CreatedAt            *time.Time         `json:"created_at" gorm:"created_at,autoCreateTime"`
	Info                 string             `json:"info" gorm:"info"`
	StaffID              uint64             `json:"staff_id" gorm:"staff_id"`
	Staff                *UserModel         `json:"staff" gorm:"ForeignKey:UserID"`
	StartAt              string             `json:"start_at" gorm:"start_at"`
	EndAt                *time.Time         `json:"end_at" gorm:"end_at"`
	Status               uint8              `json:"status" gorm:"status"`
	UpdatedAt            *time.Time         `json:"updated_at" gorm:"updated_at,autoUpdateTime"`
	Income               float64            `json:"income" gorm:"income"`
	Subject              string             `json:"subject" gorm:"subject"`
	CaseType             string             `json:"case_type" gorm:"case_type"`
	LaboratoryCases      string             `json:"laboratory_cases" gorm:"laboratory_cases"`
	PhotographyCases     string             `json:"photography_cases" gorm:"photography_cases"`
	RadiologyCases       string             `json:"radiology_cases" gorm:"radiology_cases"`
	LastLaboratoryCases  string             `json:"last_laboratory_cases" gorm:"last_laboratory_cases"`
	LastPhotographyCases string             `json:"last_photography_cases" gorm:"last_photography_cases"`
	LastRadiologyCases   string             `json:"last_radiology_cases" gorm:"last_radiology_cases"`
	Prescription         string             `json:"prescription" gorm:"prescription"`
	FuturePrescription   string             `json:"future_prescription" gorm:"future_prescription"`
	LaboratoryMsg        string             `json:"laboratory_msg" gorm:"laboratory_msg"`
	PhotographyMsg       string             `json:"photography_msg" gorm:"photography_msg"`
	RadiologyMsg         string             `json:"radiology_msg" gorm:"radiology_msg"`
	OrganizationID       uint64             `json:"organization_id" gorm:"organization_id,index"`
	Organization         *OrganizationModel `json:"organization" gorm:"ForeignKey:OrganizationID"`
	LaboratoryID         uint64             `json:"laboratory_id" gorm:"laboratory_id"`
	Laboratory           *OrganizationModel `json:"laboratory" gorm:"ForeignKey:LaboratoryID"`
	PhotographyID        uint64             `json:"photography_id" gorm:"photography_id"`
	Photography          *OrganizationModel `json:"photography" gorm:"ForeignKey:PhotographyID"`
	RadiologyID          uint64             `json:"radiology_id" gorm:"radiology_id"`
	Radiology            *OrganizationModel `json:"radiology" gorm:"ForeignKey:RadiologyID"`
	LAdmissionAt         *time.Time         `json:"l_admission_at" gorm:"l_admission_at"`
	PAdmissionAt         *time.Time         `json:"p_admission_at" gorm:"p_admission_at"`
	RAdmissionAt         *time.Time         `json:"r_admission_at" gorm:"r_admission_at"`
	LResultAt            *time.Time         `json:"l_result_at" gorm:"l_result_at"`
	PResultAt            *time.Time         `json:"p_result_at" gorm:"p_result_at"`
	RResultAt            *time.Time         `json:"r_result_at" gorm:"r_result_at"`
	LRndImg              string             `json:"l_rnd_img" gorm:"l_rnd_img"`
	PRndImg              string             `json:"p_rnd_img" gorm:"p_rnd_img"`
	RRndImg              string             `json:"r_rnd_img" gorm:"r_rnd_img"`
	LImgs                int                `json:"l_imgs" gorm:"l_imgs"`
	PImgs                int                `json:"p_imgs" gorm:"p_imgs"`
	RImgs                int                `json:"r_imgs" gorm:"r_imgs"`
	Code                 string             `json:"code" gorm:"code"`
	IsVip                bool               `json:"is_vip" gorm:"is_vip"`
	VipIntroducer        uint64             `json:"vip_introducer" gorm:"vip_introducer"`
	Absence              bool               `json:"absence" gorm:"absence"`
	FileID               string             `json:"file_id" gorm:"file_id"`
	Price                float64            `json:"price" gorm:"price"`
	PhotographyStatus    bool               `json:"photography_status" gorm:"photography_status"`
	RadiologyStatus      bool               `json:"radiology_status" gorm:"radiology_status"`
	OfficeID             uint64             `json:"office_id" gorm:"office_id"`
}

func (AppointmentModel) TableName() string {
	return "appointment"
}

type QueStruct struct {
	DefaultDuration int                `json:"default_duration"`
	Limits          []CaseType         `json:"limits"`
	Ques            []AppointmentModel `json:"ques"`
	WorkHour        WorkHour           `json:"work_hour"`
}

type WorkHour struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
