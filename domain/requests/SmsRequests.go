package requests

import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
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

type SendTemplateSMS struct {
	Token    string `json:"token"`
	Token2   string `json:"token2"`
	Token3   string `json:"token3"`
	Receptor string `json:"receptor"`
	Template string `json:"template"`
}

func (s *SendTemplateSMS) Send() (bool, *string, error) {
	receptor := helpers.NormalizePhoneNumber(s.Receptor)
	params := &kavenegar.VerifyLookupParam{
		Token2: s.Token2,
		Token3: s.Token3,
	}
	if res, err := sms.Kavenegar.Verify.Lookup(receptor, s.Template, s.Token, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
		return false, nil, err
	} else {
		fmt.Println("MessageID 	= ", res.MessageID)
		fmt.Println("Status    	= ", res.Status)
		return true, &res.Message, nil
	}
}
