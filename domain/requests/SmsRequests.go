package requests

import (
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/sms"
)

type SendSMSRequest struct {
	UserID         uint64   `json:"user_id"`
	Numbers        []string `json:"numbers"`
	PhoneNumber    string   `json:"phone_number"`
	Msg            string   `json:"msg"`
	Type           uint     `json:"type"`
	OrganizationID uint64   `json:"organization_id"`
}

func (s *SendSMSRequest) SendSMS() (bool, *string, error) {
	var receptor []string
	for i := 0; i < len(s.Numbers); i++ {
		n := helpers.NormalizePhoneNumber(s.Numbers[i])
		if n != "" {
			receptor = append(receptor, n)
		}
	}
	message := s.Msg
	send, res, error := sms.SendByPackage(receptor, message)
	return send, res, error
}
