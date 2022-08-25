package sms

import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/env"
	"log"
)

var ApiKey = ""
var Sender = ""

var Kavenegar *kavenegar.Kavenegar

func init() {
	ApiKey = env.GetEnv("SMS_KEY")
	Sender = env.GetEnv("SMS_SENDER")
	Kavenegar = kavenegar.New(ApiKey)
}

func SendByPackage(receptor []string, message string) (bool, *string, error) {
	if _, err := Kavenegar.Message.Send(Sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			log.Println(err.Error())
			break
		case *kavenegar.HTTPError:
			log.Println(err.Error())
			break
		default:
			log.Println(err.Error())
			break
		}
		r := err.Error()
		return false, &r, err
	} else {
		//for _, r := range res {
		//}
		return true, nil, nil
	}
}

type TemplateSMS struct {
	Type     string
	Template string
	Token    string
	Token2   string
	Token3   string
	Receptor string
	Tokens   map[string]string
}

func (s *TemplateSMS) Send() (bool, *string, error) {
	receptor := helpers.NormalizePhoneNumber(s.Receptor)
	params := &kavenegar.VerifyLookupParam{
		Tokens: s.Tokens,
		Token2: s.Token2,
		Token3: s.Token3,
	}
	if res, err := Kavenegar.Verify.Lookup(receptor, s.Template, s.Token, params); err != nil {
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
