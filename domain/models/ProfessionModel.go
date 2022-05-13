package models

type ProfessionModel struct {
	ID               uint64 `json:"id" gorm:"primarykey"`
	Name             string `json:"name" gorm:"name,index"`
	LaboratoryCases  string `json:"laboratory_cases" gorm:"laboratory_cases"`
	PhotographyCases string `json:"photography_cases" gorm:"photography_cases"`
	RadiologyCases   string `json:"radiology_cases" gorm:"radiology_cases"`
}

func (ProfessionModel) TableName() string {
	return "profession"
}
