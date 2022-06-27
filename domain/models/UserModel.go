package models

import (
	"time"
)

type UserModel struct {
	ID             uint64             `json:"id" gorm:"primarykey"`
	Fname          string             `json:"fname" gorm:"fname,index"`
	Lname          string             `json:"lname" gorm:"lname,index"`
	Email          string             `json:"email" gorm:"email"`
	Point          uint64             `json:"point" gorm:"point"`
	UserGroupID    uint64             `json:"user_group_id" gorm:"user_group_id"`
	UserGroup      UserGroupModel     `json:"user_group" gorm:"user_group"`
	Created        *time.Time         `json:"created" gorm:"created,autoCreateTime"`
	Tel            string             `json:"tel" gorm:"tel,index"`
	Cardno         string             `json:"cardno" gorm:"cardno"`
	StaffID        uint64             `json:"staff_id" gorm:"staff_id"`
	Staff          *UserModel         `json:"staff" gorm:"foreignkey:StaffID"`
	Pass           string             `json:"-" gorm:"pass"`
	Pin            string             `json:"pin" gorm:"pin"`
	TelegramID     string             `json:"telegram_id" gorm:"telegram_id"`
	Info           string             `json:"info" gorm:"info"`
	Guest          bool               `json:"guest" gorm:"guest"`
	LastLogin      *time.Time         `json:"last_login" gorm:"last_login"`
	Gender         string             `json:"gender" gorm:"gender"`
	RndImg         string             `json:"rnd_img" gorm:"rnd_img"`
	KnownAs        string             `json:"known_as" gorm:"known_as"`
	MaritalStatus  string             `json:"marital_status" gorm:"marital_status"`
	BirthDate      *time.Time         `json:"birth_date" gorm:"birth_date"`
	Tel1           string             `json:"tel1" gorm:"tel1"`
	Tel2           string             `json:"tel2" gorm:"tel2"`
	Tel3           string             `json:"tel3" gorm:"tel3"`
	Relation       string             `json:"relation" gorm:"relation"`
	Nid            string             `json:"nid" gorm:"nid"`
	FileID         string             `json:"file_id" gorm:"file_id"`
	OrganizationID uint64             `json:"organization_id" gorm:"organization_id,index"`
	Organization   *OrganizationModel `json:"organization" gorm:"ForeignKey:OrganizationID"`
	Address        string             `json:"address" gorm:"address"`
	Introducer     string             `json:"introducer" gorm:"introducer"`
	Appcode        string             `json:"appcode" gorm:"appcode"`
	DuePayment     int                `json:"due_payment" gorm:"due_payment"`
	CityID         uint64             `json:"city_id" gorm:"city_id"`
	City           *CityModel         `json:"city" gorm:"foreignkey:CityID"`
	IsVip          bool               `json:"is_vip" gorm:"is_vip"`
	Age            int                `json:"age" gorm:"-"`
	CountyID       *uint64            `json:"county_id" gorm:"-"`
	County         *CountyModel       `json:"county" gorm:"-"`
	ProvinceID     *uint64            `json:"province_id" gorm:"-"`
	Province       *ProvinceModel     `json:"province" gorm:"-"`
}

func (UserModel) TableName() string {
	return "user"
}

func (u *UserModel) IsAdmin() bool {
	return u.UserGroupID == 2
}

func (u *UserModel) IsPatient() bool {
	return u.UserGroupID == 1
}

func (u *UserModel) IsDoctor() bool {
	return u.UserGroupID == 3
}

func (u *UserModel) IsDoctorSecretary() bool {
	return u.UserGroupID == 4
}

func (u *UserModel) IsLabAdmin() bool {
	return u.UserGroupID == 5
}

func (u *UserModel) IsSupport() bool {
	return u.UserGroupID == 100
}
