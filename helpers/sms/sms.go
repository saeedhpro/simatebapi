package sms

import (
	"github.com/kavenegar/kavenegar-go"
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
	if res, err := Kavenegar.Message.Send(Sender, receptor, message, nil); err != nil {
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
		for _, r := range res {
			log.Println("MessageID 	= ", r.MessageID)
			log.Println("Status    	= ", r.Status)
		}
		return true, nil, nil
	}
}
