package user

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	sms2 "github.com/saeedhpro/apisimateb/helpers/sms"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"strings"
)

func SendRegisterSms(u *models.UserModel) {
	organization, _ := organizationRepository.GetOrganizationByID(u.OrganizationID)
	gender := ""
	if u.Gender != "" {
		if strings.ToLower(u.Gender) == "male" {
			gender = "آقای "
		} else {
			gender = "خانم "
		}
	}
	sms := sms2.TemplateSMS{
		Receptor: u.Tel,
		Template: "Information",
		Token:    organization.Phone,
		Tokens: map[string]string{
			"token20": fmt.Sprintf("%s %s %s", gender, u.Fname, u.Lname),
			"token10": organization.Name,
		},
	}
	sms.Send()
}
