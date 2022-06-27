package requests

type UserCreateRequest struct {
	Fname          string  `json:"fname" binding:"required"`
	Lname          string  `json:"lname" binding:"required"`
	OrganizationID *uint64 `json:"organization_id"`
	UserGroupID    *uint64 `json:"user_group_id"`
	KnownAs        string  `json:"known_as"`
	Gender         string  `json:"gender"`
	Tel            string  `json:"tel"`
	Tel1           string  `json:"tel1"`
	Cardno         string  `json:"cardno"`
	BirthDate      string  `json:"birth_date"`
	FileID         string  `json:"file_id"`
	Address        string  `json:"address"`
	Info           string  `json:"info"`
	Introducer     string  `json:"introducer"`
	CityID         *uint64 `json:"city_id"`
}

type UserUpdateRequest struct {
	ID             uint64  `json:"id"`
	Fname          string  `json:"fname" binding:"required"`
	Lname          string  `json:"lname" binding:"required"`
	OrganizationID *uint64 `json:"organization_id"`
	UserGroupID    *uint64 `json:"user_group_id"`
	KnownAs        string  `json:"known_as"`
	Gender         string  `json:"gender"`
	Tel            string  `json:"tel"`
	Tel1           string  `json:"tel1"`
	Cardno         string  `json:"cardno"`
	BirthDate      string  `json:"birth_date"`
	FileID         string  `json:"file_id"`
	Address        string  `json:"address"`
	Info           string  `json:"info"`
	Introducer     string  `json:"introducer"`
	CityID         *uint64 `json:"city_id"`
}

type UserListDeleteRequest struct {
	IDs []uint64 `json:"ids"`
}

type CreateUserMedicalHistoryRequest struct {
	AdenoidTonsileReduction string `json:"adenoid_tonsile_reduction"`
	MedicalCondition        string `json:"medical_condition"`
	ConsumableMedicine      string `json:"consumable_medicine"`
	GeneralHealth           string `json:"general_health"`
	UnderPhysicianCare      string `json:"under_physician_care"`
	AccidentToHead          string `json:"accident_to_Head"`
	Operations              string `json:"operations"`
	ChiefComplaint          string `json:"chief_complaint"`
	PreviousOrthodontic     string `json:"previous_orthodontic"`
	OralHygiene             string `json:"oral_hygiene"`
	Frontal                 string `json:"frontal"`
	Profile                 string `json:"profile"`
	TeethPresent            string `json:"teeth_present"`
	UnErupted               string `json:"un_erupted"`
	IeMissing               string `json:"ie_missing"`
	IeExtracted             string `json:"ie_extracted"`
	IeImpacted              string `json:"ie_impacted"`
	IeSupernumerary         string `json:"ie_supernumerary"`
	IeCaries                string `json:"ie_caries"`
	IeRct                   string `json:"ie_rct"`
	IeAnomalies             string `json:"ie_anomalies"`
	LeftMolar               string `json:"left_molar"`
	RightMolar              string `json:"right_molar"`
	LeftCanine              string `json:"left_canine"`
	RightCanine             string `json:"right_canine"`
	Overjet                 string `json:"overjet"`
	Overbite                string `json:"overbite"`
	Crossbite               string `json:"crossbite"`
	CrowdingMd              string `json:"crowding_md"`
	CrowdingMx              string `json:"crowding_mx"`
	SpacingMx               string `json:"spacing_mx"`
	SpacingMd               string `json:"spacing_md"`
	Diagnosis               string `json:"diagnosis"`
	TreatmentPlan           string `json:"treatment_plan"`
	LengthActiveTreatment   string `json:"length_active_treatment"`
	Retention               string `json:"retention"`
	UserID                  uint64 `json:"user_id"`
}
