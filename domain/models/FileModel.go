package models

type FileModel struct {
	ID             uint64            `json:"id" gorm:"primarykey"`
	Path           string            `json:"path" gorm:"path"`
	Ext            string            `json:"ext" gorm:"ext"`
	Info           string            `json:"info" gorm:"info"`
	Comment        string            `json:"comment" gorm:"comment"`
	OrganizationID uint64            `json:"organization_id" gorm:"organization_id,index"`
	Organization   OrganizationModel `json:"organization" gorm:"ForeignKey:OrganizationID"`
	StaffID        uint64            `json:"staff_id" gorm:"staff_id"`
	Staff          UserModel         `json:"staff" gorm:"foreignKey:staff_id"`
	UserID         uint64            `json:"user_id" gorm:"user_id"`
	User           UserModel         `json:"user" gorm:"ForeignKey:UserID"`
}

func (FileModel) TableName() string {
	return "file"
}
