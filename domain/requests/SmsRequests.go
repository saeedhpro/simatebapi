package requests

import (
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/sms"
)

type SendSMSRequest struct {
	UserID         uint64   `json:"user_id"`
	Number         []string `json:"numbers"`
	Msg            string   `json:"msg"`
	OrganizationID uint64   `json:"organization_id"`
}

func (s *SendSMSRequest) SendSMS() (bool, *string, error) {
	var receptor []string
	for i := 0; i < len(s.Number); i++ {
		n := helpers.NormalizePhoneNumber(s.Number[i])
		if n != "" {
			receptor = append(receptor, n)
		}
	}
	message := s.Msg
	send, res, error := sms.SendByPackage(receptor, message)
	return send, res, error
}
