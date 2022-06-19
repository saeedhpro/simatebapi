package appointmentRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
	"time"
)

func GetAppointmentBy(conditions *models.AppointmentModel) (*models.AppointmentModel, error) {
	var appointment models.AppointmentModel
	err := repository.DB.MySQL.Preload("Organization").Preload("Photography").Preload("Radiology").Preload("Staff").Preload("User").First(&appointment, &conditions).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func GetAppointmentListBy(conditions *models.AppointmentModel, start string, end string, isDoctor bool) ([]models.AppointmentModel, error) {
	appointments := []models.AppointmentModel{}
	query := repository.DB.MySQL.
		Preload("Organization").
		Preload("Photography").
		Preload("Radiology").
		Preload("Staff").
		Preload("User")
	fmt.Println(isDoctor)
	if !isDoctor {
		if conditions.Photography != nil {
			query = query.Where("p_admission_at IS NOT NULL and p_result_at IS NULL and status = 2")
		} else if conditions.Radiology != nil {
			query = query.Where("r_admission_at IS NOT NULL and r_result_at IS NULL and status = 2")
		} else if conditions.Laboratory != nil {
			query = query.Where("l_admission_at IS NOT NULL and l_result_at IS NULL and status = 2")
		}
	} else {
		if start != "" {
			query = query.Where("start_at >= ?", fmt.Sprintf("%s 00:00:00", start))
		}
		if end != "" {
			query = query.Where("start_at < ?", fmt.Sprintf("%s 23:59:59", end))
		}
	}
	err := query.
		Find(&appointments, &conditions).Error
	if err != nil {
		return appointments, err
	}
	return appointments, nil
}

func GetPaginatedAppointmentListBy(conditions *models.AppointmentModel, start string, end string, isDoctor bool, page int, limit int) (pagination.Pagination, error) {
	appointments := []models.AppointmentModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if !isDoctor {
		if conditions.Photography != nil {
			query = query.Where("p_admission_at IS NOT NULL and p_result_at IS NULL and status = 2")
		} else if conditions.Radiology != nil {
			query = query.Where("r_admission_at IS NOT NULL and r_result_at IS NULL and status = 2")
		} else if conditions.Laboratory != nil {
			query = query.Where("l_admission_at IS NOT NULL and l_result_at IS NULL and status = 2")
		}
	} else {
		if start != "" {
			query = query.Where("start_at >= ?", fmt.Sprintf("%s 00:00:00", start))
		}
		if end != "" {
			query = query.Where("start_at < ?", fmt.Sprintf("%s 23:59:59", end))
		}
	}
	query.Find(&appointments, &conditions).Count(&count)
	query = repository.DB.MySQL.Scopes(pagination.PaginateScope(count, &paginate)).Preload("Organization").Preload("Photography").Preload("Radiology").Preload("Staff").Preload("User")
	if !isDoctor {
		if conditions.Photography != nil {
			query = query.Where("p_admission_at IS NOT NULL and p_result_at IS NULL and status = 2")
		} else if conditions.Radiology != nil {
			query = query.Where("r_admission_at IS NOT NULL and r_result_at IS NULL and status = 2")
		} else if conditions.Laboratory != nil {
			query = query.Where("l_admission_at IS NOT NULL and l_result_at IS NULL and status = 2")
		}
	} else {
		if start != "" {
			query = query.Where("start_at >= ?", fmt.Sprintf("%s 00:00:00", start))
		}
		if end != "" {
			query = query.Where("start_at < ?", fmt.Sprintf("%s 23:59:59", end))
		}
	}
	err := query.Find(&appointments, &conditions).Error
	if err != nil {
		return paginate, err
	}
	paginate.Data = appointments
	return paginate, nil
}

func FilterOrganizationAppointment(organizationID uint64, status []string, q string, start string, end string, isDoctor bool, page int, limit int) (pagination.Pagination, error) {
	appointments := []models.AppointmentModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	query := repository.DB.MySQL.Preload("Organization").Preload("Photography").Preload("Radiology").Preload("Staff").Preload("User")
	query = query.Where("organization_id", organizationID)
	if len(status) > 0 {
		query = query.Where("status IN ?", status)
	}
	if q != "" {
		query = query.
			Joins("left join (select id, fname, lname from user) user on appointment.user_id = user.id").
			Where(repository.DB.MySQL.
				Where("fname LIKE ?", "%"+q+"%").
				Or("lname LIKE ?", "%"+q+"%"))
	}
	if start != "" {
		query = query.
			Where("start_at >= ?", fmt.Sprintf("%s 00:00:00", start))
	}
	if end != "" {
		query = query.
			Where("start_at < ?", fmt.Sprintf("%s 23:59:59", end))
	}
	var count int64 = 0
	query.Find(&appointments).Count(&count)
	err := query.Scopes(pagination.PaginateScope(count, &paginate)).Find(&appointments).Error
	if err != nil {
		fmt.Println(err.Error())
		return paginate, err
	}
	paginate.Data = appointments
	return paginate, nil
}

func GetAppointmentByID(ID uint64) (*models.AppointmentModel, error) {
	var appointment models.AppointmentModel
	appointment.ID = ID
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Photography").
		Preload("Radiology").
		Preload("Staff").
		Preload("User").
		First(&appointment, &appointment).
		Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func GetAppointmentListBetweenDates(organizationID *uint64, startDate string, endDate string) ([]models.AppointmentModel, error) {
	appointments := []models.AppointmentModel{}
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Photography").
		Preload("Radiology").
		Preload("Staff").
		Preload("User").
		Where("start_at between ? and ?", startDate, endDate).
		Find(&appointments, models.AppointmentModel{OrganizationID: *organizationID}).Error
	if err != nil {
		fmt.Println(err.Error())
		return appointments, err
	}
	return appointments, nil
}

func CreateAppointment(request *requests.AppointmentCreateRequest, staffID uint64, organizationID uint64) (*models.AppointmentModel, error) {
	appointment := models.AppointmentModel{
		OrganizationID: organizationID,
		StaffID:        staffID,
		UserID:         request.UserID,
		Income:         request.Income,
		Info:           request.Info,
		CaseType:       request.CaseType,
		StartAt:        request.StartAt,
		Status:         1,
	}
	err := repository.DB.MySQL.Create(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func AcceptAppointment(request *requests.AppointmentUpdateRequest) (bool, error) {
	appointment := models.AppointmentModel{}
	appointment.ID = request.ID
	appointment.Status = 2
	appointment.Info = request.Info
	appointment.LaboratoryCases = request.LaboratoryCases
	appointment.PhotographyCases = request.PhotographyCases
	appointment.RadiologyCases = request.RadiologyCases
	appointment.Prescription = request.Prescription
	appointment.FuturePrescription = request.FuturePrescription
	appointment.LaboratoryMsg = request.LaboratoryMsg
	appointment.PhotographyMsg = request.PhotographyMsg
	appointment.RadiologyMsg = request.RadiologyMsg
	appointment.LaboratoryID = request.LaboratoryID
	appointment.PhotographyID = request.PhotographyID
	appointment.RadiologyID = request.RadiologyID
	err := repository.DB.MySQL.
		Model(&appointment).
		Updates(&appointment).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateAppointment(request *requests.AppointmentUpdateRequest) (bool, error) {
	appointment := models.AppointmentModel{}
	appointment.ID = request.ID
	appointment.Status = request.Status
	appointment.Info = request.Info
	appointment.LaboratoryCases = request.LaboratoryCases
	appointment.PhotographyCases = request.PhotographyCases
	appointment.RadiologyCases = request.RadiologyCases
	appointment.Prescription = request.Prescription
	appointment.FuturePrescription = request.FuturePrescription
	appointment.LaboratoryMsg = request.LaboratoryMsg
	appointment.PhotographyMsg = request.PhotographyMsg
	appointment.RadiologyMsg = request.RadiologyMsg
	appointment.LaboratoryID = request.LaboratoryID
	appointment.PhotographyID = request.PhotographyID
	appointment.RadiologyID = request.RadiologyID
	r, _ := time.Parse("2006-04-01 11:35:54", *request.RResultAt)
	l, _ := time.Parse("2006-04-01 11:35:54", *request.LResultAt)
	p, _ := time.Parse("2006-04-01 11:35:54", *request.PResultAt)
	appointment.RResultAt = &r
	appointment.LResultAt = &l
	appointment.PResultAt = &p
	appointment.RRndImg = request.RRndImg
	appointment.PRndImg = request.PRndImg
	appointment.LRndImg = request.LRndImg
	err := repository.DB.MySQL.
		Model(&appointment).
		Updates(&appointment).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CancelAppointment(appointment *models.AppointmentModel) (bool, error) {
	appointment.Status = 3
	appointment.Organization = nil
	err := repository.DB.MySQL.
		Model(&appointment).
		Updates(&appointment).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReserveAppointment(appointment *models.AppointmentModel) (bool, error) {
	appointment.Status = 1
	appointment.Organization = nil
	err := repository.DB.MySQL.
		Model(&appointment).
		Updates(&appointment).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}
