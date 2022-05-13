package responses

import "github.com/saeedhpro/apisimateb/domain/models"

type UserLoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
	User      *models.UserModel `json:"user"`
}
