package models

type UserGroupModel struct {
	ID             uint64            `json:"id" gorm:"primarykey"`
	Name          string            `json:"name" gorm:"name,index"`
}

func (UserGroupModel) TableName() string {
	return "user_group"
}

func (u *UserGroupModel) IsAdmin() bool {
	return u.ID == 2
}

func (u *UserGroupModel) IsPatient() bool {
	return u.ID == 1
}

func (u *UserGroupModel) IsDoctor() bool {
	return u.ID == 3
}

func (u *UserGroupModel) IsDoctorSecretary() bool {
	return u.ID == 4
}

func (u *UserGroupModel) IsLabAdmin() bool {
	return u.ID == 5
}

func (u *UserGroupModel) IsSupport() bool {
	return u.ID == 100
}
