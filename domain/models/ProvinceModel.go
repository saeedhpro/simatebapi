package models

type ProvinceModel struct {
	ID       uint64        `json:"id" gorm:"primarykey"`
	Name     string        `json:"name" gorm:"name,index"`
	Counties []CountyModel `json:"counties" gorm:"foreignKey:ProvinceID"`
}

func (ProvinceModel) TableName() string {
	return "province"
}


