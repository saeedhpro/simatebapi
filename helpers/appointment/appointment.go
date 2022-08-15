package appointment

import (
	"fmt"
	"github.com/jalaali/go-jalaali"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"strings"
	"time"
)

func SendAppointmentCreatedSMS(request *requests.AppointmentCreateRequest, appointment *models.AppointmentModel) {
	user, _ := userRepository.GetUserByID(request.UserID)
	organization, _ := organizationRepository.GetOrganizationByID(appointment.OrganizationID)
	t, err := time.Parse("2006-01-02 15:04:05", request.StartAt)
	if err == nil {
		date, err1 := jalaali.From(t).JFormat("2006 Jan 02")
		if err1 != nil {
			fmt.Println(err1.Error())
			date = ""
		}
		dateStr := fmt.Sprintf("%s %s", GetPersianDay(jalaali.From(t).Weekday().String()), date)
		sms := requests.SendTemplateSMS{
			Receptor: user.Tel,
			Template: "reserve",
			Token:    strings.Split(request.StartAt, " ")[1],
			Token3:   organization.Name,
			Token2:   dateStr,
		}
		sms.Send()
	} else {
		fmt.Println(err.Error())
	}
}

func SendFileSentSMS(appointment *models.AppointmentModel) {
	user, _ := userRepository.GetUserByID(appointment.UserID)
	organization, _ := organizationRepository.GetOrganizationByID(appointment.OrganizationID)
	sms := requests.SendTemplateSMS{
		Receptor: user.Tel,
		Template: "filesend",
		Token:    user.Tel,
		Token3:   organization.Name,
		Token2:   appointment.Appcode,
	}
	sms.Send()
}

func GetPersianDay(day string) string {
	var days = []string{
		"شنبه", "یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه",
	}
	switch day {
	case "Saturday":
		return days[0]
	case "Sunday":
		return days[1]
	case "Monday":
		return days[2]
	case "Tuesday":
		return days[3]
	case "Wednesday":
		return days[4]
	case "Thursday":
		return days[5]
	case "Friday":
		return days[6]
	default:
		return ""
	}
}
