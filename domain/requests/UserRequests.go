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

type UserListDeleteRequest struct {
	IDs []uint64 `json:"ids"`
}