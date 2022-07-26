package requests

type AppointmentCreateRequest struct {
	CaseType string  `json:"case_type"`
	Income   float64 `json:"income"`
	Info     string  `json:"info"`
	StartAt  string  `json:"start_at" binding:"required"`
	UserID   uint64  `json:"user_id" binding:"required"`
}

type AppointmentUpdateRequest struct {
	ID                 uint64   `json:"id"`
	UserID             uint64   `json:"user_id"`
	CreatedAt          *string  `json:"created_at"`
	Info               string   `json:"info"`
	StaffID            uint64   `json:"staff_id"`
	StartAt            string   `json:"start_at"`
	EndAt              *string  `json:"end_at"`
	Status             uint8    `json:"status"`
	UpdatedAt          *string  `json:"updated_at"`
	Income             float64  `json:"income"`
	Subject            string   `json:"subject"`
	CaseType           string   `json:"case_type"`
	LaboratoryCases    string   `json:"laboratory_cases"`
	PhotographyCases   string   `json:"photography_cases"`
	RadiologyCases     string   `json:"radiology_cases"`
	Prescription       string   `json:"prescription"`
	FuturePrescription string   `json:"future_prescription"`
	LaboratoryMsg      string   `json:"laboratory_msg"`
	PhotographyMsg     string   `json:"photography_msg"`
	RadiologyMsg       string   `json:"radiology_msg"`
	OrganizationID     uint64   `json:"organization_id"`
	LaboratoryID       uint64   `json:"laboratory_id"`
	PhotographyID      uint64   `json:"photography_id"`
	RadiologyID        uint64   `json:"radiology_id"`
	LAdmissionAt       *string  `json:"l_admission_at"`
	PAdmissionAt       *string  `json:"p_admission_at"`
	RAdmissionAt       *string  `json:"r_admission_at"`
	LResultAt          *string  `json:"l_result_at"`
	PResultAt          *string  `json:"p_result_at"`
	RResultAt          *string  `json:"r_result_at"`
	LRndImg            string   `json:"l_rnd_img"`
	PRndImg            string   `json:"p_rnd_img"`
	RRndImg            string   `json:"r_rnd_img"`
	LImgs              int      `json:"l_imgs"`
	PImgs              int      `json:"p_imgs"`
	RImgs              int      `json:"r_imgs"`
	Code               string   `json:"code"`
	IsVip              bool     `json:"is_vip"`
	VipIntroducer      uint64   `json:"vip_introducer"`
	Results            []string `json:"results"`
}

type AddAppointmentResultsRequest struct {
	ID      uint64   `json:"id"`
	Results []string `json:"results"`
}
