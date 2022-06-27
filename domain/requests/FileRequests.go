package requests

type FileCreateRequest struct {
	File           string `json:"file"`
	Path           string `json:"path"`
	Ext            string `json:"ext"`
	Info           string `json:"info"`
	Comment        string `json:"comment"`
	UserID         uint64 `json:"user_id"`
	OrganizationID uint64 `json:"organization_id"`
	StaffID        uint64 `json:"staff_id"`
}

type FileUpdateRequest struct {
	ID      uint64 `json:"id"`
	Path    string `json:"path"`
	Info    string `json:"info"`
	Comment string `json:"comment"`
}
