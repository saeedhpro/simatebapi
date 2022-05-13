package requests

type UserLoginRequest struct {
	Tel      string `json:"tel" binding:"required"`
	Password string `json:"password" binding:"required"`
}

