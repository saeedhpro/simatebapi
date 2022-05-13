package models

type CityModel struct {
	ID         uint64        `json:"id" gorm:"primarykey"`
	Name       string        `json:"name" gorm:"name,index"`
	CountyID uint64        `json:"county_id" gorm:"county_id"`
	County   CountyModel `json:"county" gorm:"foreignkey:CountyID"`
}

func (CityModel) TableName() string {
	return "city"
}
