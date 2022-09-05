package appointment

import (
	"fmt"
	"github.com/jalaali/go-jalaali"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	sms2 "github.com/saeedhpro/apisimateb/helpers/sms"
	"github.com/saeedhpro/apisimateb/repository/organizationRepository"
	"github.com/saeedhpro/apisimateb/repository/userRepository"
	"strings"
	"time"
)

func SendAppointmentCreatedSMS(request *requests.AppointmentCreateRequest, appointment *models.AppointmentModel) {
	t, err := time.Parse("2006-01-02 15:04:05", request.StartAt)
	if time.Now().Before(t) {
		user, _ := userRepository.GetUserByID(request.UserID)
		organization, _ := organizationRepository.GetOrganizationByID(appointment.OrganizationID)
		if err == nil {
			date, err1 := jalaali.From(t).JFormat("02 Jan")
			if err1 != nil {
				fmt.Println(err1.Error())
				date = ""
			}
			dateStr := fmt.Sprintf("%s %s", GetPersianDay(jalaali.From(t).Weekday().String()), date)
			sms := sms2.TemplateSMS{
				Receptor: user.Tel,
				Template: "Reservation",
				Token:    strings.Split(request.StartAt, " ")[1],
				Tokens: map[string]string{
					"token10": organization.Name,
					"token20": dateStr,
				},
			}
			sms.Send()
		} else {
			fmt.Println(err.Error())
		}
	}
}

func SendAppointmentCodeSMS(appointment *models.AppointmentModel) {
	t, err := time.Parse(time.RFC3339Nano, appointment.StartAt)
	if time.Now().Before(t) {
		if err == nil {
			user, _ := userRepository.GetUserByID(appointment.UserID)
			date, err1 := jalaali.From(t).JFormat("2006 Jan 02")
			if err1 != nil {
				fmt.Println(err1.Error())
				date = ""
			}
			tm, err1 := jalaali.From(t).JFormat("15:04:05")
			if err1 != nil {
				fmt.Println(err1.Error())
				tm = ""
			}
			dateStr := fmt.Sprintf("%s %s", GetPersianDay(jalaali.From(t).Weekday().String()), date)
			sms := sms2.TemplateSMS{
				Receptor: user.Tel,
				Template: "APPOINTMENT",
				Token:    appointment.Code,
				Token2:   tm,
				Tokens: map[string]string{
					"token20": dateStr,
				},
			}
			sms.Send()
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(time.Now().After(t))
	}
}

func SendUserFileSMS(organizationName string, tel string, password string) {
	sms := sms2.TemplateSMS{
		Receptor: tel,
		Template: "filesend",
		Token:    tel,
		Token2:   password,
		Tokens: map[string]string{
			"token20": organizationName,
		},
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
