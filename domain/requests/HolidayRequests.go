package requests

type HolidayCreateRequest struct {
	Hdate          string  `json:"hdate"`
	Title          string  `json:"title"`
	OrganizationID *uint64 `json:"organization_id"`
}
