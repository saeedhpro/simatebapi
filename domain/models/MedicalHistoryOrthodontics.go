package models

type MedicalHistoryOrthodontics struct {
	ID                      uint64     `json:"id" gorm:"primarykey"`
	UserID                  uint64     `json:"user_id" gorm:"user_id"`
	User                    *UserModel `json:"user" gorm:"foreignkey:UserID"`
	AdenoidTonsileReduction string     `json:"adenoid_tonsile_reduction" gorm:"adenoid_tonsile_reduction"`
	MedicalCondition        string     `json:"medical_condition" gorm:"medical_condition"`
	ConsumableMedicine      string     `json:"consumable_medicine" gorm:"consumable_medicine"`
	GeneralHealth           string     `json:"general_health" gorm:"general_health"`
	UnderPhysicianCare      string     `json:"under_physician_care" gorm:"under_physician_care"`
	AccidentToHead          string     `json:"accident_to_Head" gorm:"accident_to_Head"`
	Operations              string     `json:"operations" gorm:"operations"`
	ChiefComplaint          string     `json:"chief_complaint" gorm:"chief_complaint"`
	PreviousOrthodontic     string     `json:"previous_orthodontic" gorm:"previous_orthodontic"`
	OralHygiene             string     `json:"oral_hygiene" gorm:"oral_hygiene"`
	Frontal                 string     `json:"frontal" gorm:"frontal"`
	Profile                 string     `json:"profile" gorm:"profile"`
	TeethPresent            string     `json:"teeth_present" gorm:"teeth_present"`
	UnErupted               string     `json:"un_erupted" gorm:"un_erupted"`
	IeMissing               string     `json:"ie_missing" gorm:"ie_missing"`
	IeExtracted             string     `json:"ie_extracted" gorm:"ie_extracted"`
	IeImpacted              string     `json:"ie_impacted" gorm:"ie_impacted"`
	IeSupernumerary         string     `json:"ie_supernumerary" gorm:"ie_supernumerary"`
	IeCaries                string     `json:"ie_caries" gorm:"ie_caries"`
	IeRct                   string     `json:"ie_rct" gorm:"ie_rct"`
	IeAnomalies             string     `json:"ie_anomalies" gorm:"ie_anomalies"`
	LeftMolar               string     `json:"left_molar" gorm:"left_molar"`
	RightMolar              string     `json:"right_molar" gorm:"right_molar"`
	LeftCanine              string     `json:"left_canine" gorm:"left_canine"`
	RightCanine             string     `json:"right_canine" gorm:"right_canine"`
	Overjet                 string     `json:"overjet" gorm:"overjet"`
	Overbite                string     `json:"overbite" gorm:"overbite"`
	Crossbite               string     `json:"crossbite" gorm:"crossbite"`
	CrowdingMd              string     `json:"crowding_md" gorm:"crowding_md"`
	CrowdingMx              string     `json:"crowding_mx" gorm:"crowding_mx"`
	SpacingMx               string     `json:"spacing_mx" gorm:"spacing_mx"`
	SpacingMd               string     `json:"spacing_md" gorm:"spacing_md"`
	Diagnosis               string     `json:"diagnosis" gorm:"diagnosis"`
	TreatmentPlan           string     `json:"treatment_plan" gorm:"treatment_plan"`
	LengthActiveTreatment   string     `json:"length_active_treatment" gorm:"length_active_treatment"`
	Retention               string     `json:"retention" gorm:"retention"`
}

func (MedicalHistoryOrthodontics) TableName() string {
	return "medical_history_orthodontics"
}
