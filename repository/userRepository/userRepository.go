package userRepository

import (
	"fmt"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/helpers"
	"github.com/saeedhpro/apisimateb/helpers/pagination"
	"github.com/saeedhpro/apisimateb/repository"
	"time"
)

func GetUserBy(conditions *models.UserModel) (*models.UserModel, error) {
	var user models.UserModel
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Staff").
		Preload("UserGroup").
		First(&user, &conditions).Error
	if err != nil {
		return nil, err
	}
	if user.BirthDate != nil {
		year, _, _, _, _, _ := helpers.TimeDiff(*user.BirthDate, time.Now())
		user.Age = year
	}
	return &user, nil
}

func GetLastOnlineUsers() ([]models.UserModel, error) {
	users := []models.UserModel{}
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Staff").
		Preload("UserGroup").
		Where("user_group_id in (3,4,5)").
		Order("last_login desc").
		Limit(10).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	return users, nil
}

func GetLastOnlinePatients() ([]models.UserModel, error) {
	users := []models.UserModel{}
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Staff").
		Preload("UserGroup").
		Order("last_login desc").
		Limit(10).
		Find(&users, &models.UserModel{
			UserGroupID: 1,
		}).
		Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	return users, nil
}

func GetOrganizationUserListBy(conditions *models.UserModel, q string, organization *models.OrganizationModel) ([]models.UserModel, error) {
	users := []models.UserModel{}
	query := repository.DB.MySQL.
		Preload("Organization").
		Preload("Staff").
		Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	query = query.
		Order("created desc")
	err := query.Find(&users, &conditions).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	return users, nil
}

func GetPaginatedOrganizationUserListBy(conditions *models.UserModel, q string, organization *models.OrganizationModel, page int, limit int) (pagination.Pagination, error) {
	users := []models.UserModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	if !organization.IsDoctor() {
		q2 := repository.DB.MySQL.Model(models.AppointmentModel{})
		if organization.IsPhotography() {
			q2 = q2.Where("appointment.photography_id = ? ", organization.ID)
		} else if organization.IsLaboratory() {
			q2 = q2.Where("appointment.laboratory_id = ? ", organization.ID)
		} else if organization.IsRadiology() {
			q2 = q2.Where("appointment.radiology_id = ? ", organization.ID)
		}
		query = query.Where("id in (?)", q2.Select("user_id"))
	}
	query = query.
		Order("created desc")
	query.Find(&users, &conditions).Count(&count)
	query = repository.DB.MySQL.Scopes(pagination.PaginateScope(count, &paginate)).Preload("Organization").Preload("Staff").Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	if !organization.IsDoctor() {
		q2 := repository.DB.MySQL.Model(models.AppointmentModel{})
		if organization.IsPhotography() {
			q2 = q2.Where("appointment.photography_id = ? ", organization.ID)
		} else if organization.IsLaboratory() {
			q2 = q2.Where("appointment.laboratory_id = ? ", organization.ID)
		} else if organization.IsRadiology() {
			q2 = q2.Where("appointment.radiology_id = ? ", organization.ID)
		}
		query = query.Where("id in (?)", q2.Select("user_id"))
	}
	query = query.
		Order("created desc")
	err := query.Find(&users, &conditions).Error
	if err != nil {
		return paginate, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	paginate.Data = users
	return paginate, nil
}

func GetUserListBy(conditions *models.UserModel, q string) ([]models.UserModel, error) {
	users := []models.UserModel{}
	query := repository.DB.MySQL.Preload("Organization").Preload("Staff").Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	query = query.
		Order("created desc")
	err := query.Find(&users, &conditions).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	return users, nil
}

func GetPaginatedUserListBy(conditions *models.UserModel, q string, page int, limit int) (pagination.Pagination, error) {
	users := []models.UserModel{}
	paginate := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}
	var count int64 = 0
	query := repository.DB.MySQL
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	query = query.
		Order("created desc")
	query.Find(&users, &conditions).Count(&count)
	query = repository.DB.MySQL.Scopes(pagination.PaginateScope(count, &paginate)).Preload("Organization").Preload("Staff").Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.
			Where("fname LIKE ?", "%"+q+"%").
			Or("lname LIKE ?", "%"+q+"%").
			Or("Concat(Concat(`fname`, ' ' ),`lname`) LIKE ?", "%"+q+"%"))
	}
	query = query.
		Order("created desc")
	err := query.Find(&users, &conditions).Error
	if err != nil {
		return paginate, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].BirthDate != nil {
			year, _, _, _, _, _ := helpers.TimeDiff(*users[i].BirthDate, time.Now())
			users[i].Age = year
		}
	}
	paginate.Data = users
	return paginate, nil
}

func GetUserByID(ID uint64) (*models.UserModel, error) {
	user := models.UserModel{ID: ID}
	err := repository.DB.MySQL.
		Preload("Organization").
		Preload("Staff").
		Preload("UserGroup").
		First(&user, &user).
		Error
	if err != nil {
		return nil, err
	}
	if user.BirthDate != nil {
		year, _, _, _, _, _ := helpers.TimeDiff(*user.BirthDate, time.Now())
		user.Age = year
	}
	return &user, nil
}

func GetNumberListByOrganizationID(ID uint64) ([]string, error) {
	var numbers []string
	query := repository.DB.MySQL.
		Table("user").
		Where("tel IS NOT NULL")
	if ID != 0 {
		query = query.
			Where("organization_id = ?", ID)
	}
	err := query.
		Select("tel").
		Find(&numbers).
		Error
	if err != nil {
		fmt.Println(err.Error())
		return numbers, err
	}
	return numbers, nil
}

func DeleteUserByID(ID uint64) (bool, error) {
	var user models.UserModel
	user.ID = ID
	err := repository.DB.MySQL.Delete(&user, &user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteUserListByID(IDs []uint64) (bool, error) {
	err := repository.DB.MySQL.Where("id IN ?", IDs).Delete(&models.UserModel{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateUser(request *requests.UserCreateRequest, staffID uint64, organizationID uint64) (*models.UserModel, error) {
	CityID := request.CityID
	if request.CityID == nil {
		CityID = nil
	}
	created := time.Now()
	user := models.UserModel{
		OrganizationID: organizationID,
		StaffID:        staffID,
		CityID:         *CityID,
		UserGroupID:    *request.UserGroupID,
		Fname:          request.Fname,
		Lname:          request.Lname,
		KnownAs:        request.KnownAs,
		Gender:         request.Gender,
		Tel:            request.Tel,
		Tel1:           request.Tel1,
		Cardno:         request.Cardno,
		FileID:         request.FileID,
		Address:        request.Address,
		Info:           request.Info,
		Introducer:     request.Introducer,
		Surgery:        request.Surgery,
		HasSurgery:     request.HasSurgery,
		Created:        &created,
	}
	if request.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", request.BirthDate)
		if err == nil {
			user.BirthDate = &birthDate
		} else {
			fmt.Println(err.Error())
		}
	}
	if request.Pass != "" {
		pass, _ := helpers.PasswordHash(request.Pass)
		user.Pass = pass
	}
	fmt.Println(user.OrganizationID)
	err := repository.DB.MySQL.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(request *requests.UserUpdateRequest) error {
	CityID := request.CityID
	if request.CityID == nil {
		CityID = nil
	}
	user := models.UserModel{
		ID:          request.ID,
		CityID:      *CityID,
		UserGroupID: *request.UserGroupID,
		Fname:       request.Fname,
		Lname:       request.Lname,
		KnownAs:     request.KnownAs,
		Gender:      request.Gender,
		Tel:         request.Tel,
		Tel1:        request.Tel1,
		Cardno:      request.Cardno,
		FileID:      request.FileID,
		Address:     request.Address,
		Info:        request.Info,
		Introducer:  request.Introducer,
		Surgery:     request.Surgery,
		HasSurgery:  request.HasSurgery,
	}
	if request.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", request.BirthDate)
		if err == nil {
			user.BirthDate = &birthDate
		} else {
			fmt.Println(err.Error())
		}
	}
	err := repository.DB.MySQL.
		Model(&user).
		Where("id = ?", user.ID).
		Updates(&user).
		Error
	fmt.Println(user.ID)
	if err != nil {
		fmt.Println(err.Error(), "err")
		return err
	}
	return nil
}
func UpdateUserPass(user *models.UserModel) error {
	err := repository.DB.MySQL.
		Model(&user).
		Where("id = ?", user.ID).
		Update("pass", user.Pass).
		Error
	fmt.Println(user.ID)
	if err != nil {
		fmt.Println(err.Error(), "err")
		return err
	}
	return nil
}
