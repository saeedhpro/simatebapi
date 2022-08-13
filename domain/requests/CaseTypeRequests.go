package requests

type CreateCaseTypeRequest struct {
	Name           string `json:"name" binding:"required"`
	Duration       uint64 `json:"duration" binding:"required"`
	IsLimited      bool   `json:"is_limited"`
	Limitation     uint64 `json:"limitation" binding:"required"`
	OrganizationID uint64 `json:"organization_id" binding:"required"`
}
