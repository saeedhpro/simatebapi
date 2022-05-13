package models

type CountyModel struct {
	ID         uint64        `json:"id" gorm:"primarykey"`
	Name       string        `json:"name" gorm:"name,index"`
	ProvinceID uint64        `json:"province_id" gorm:"province_id"`
	Province   ProvinceModel `json:"province" gorm:"foreignkey:ProvinceID"`
	Cities []CityModel `json:"cities" gorm:"foreignKey:CountyID"`
}

func (CountyModel) TableName() string {
	return "county"
}
