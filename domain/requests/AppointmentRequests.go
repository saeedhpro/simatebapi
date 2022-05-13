package requests

type AppointmentCreateRequest struct {
	CaseType string `json:"case_type"`
	Income   float64 `json:"income" binding:"required"`
	Info     string `json:"info"`
	StartAt  string `json:"start_at" binding:"required"`
	UserID   uint64 `json:"user_id" binding:"required"`
}
