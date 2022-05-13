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

func GetUserListBy(conditions *models.UserModel, q string) ([]models.UserModel, error) {
	users := []models.UserModel{}
	query := repository.DB.MySQL.Preload("Organization").Preload("Staff").Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.Where("fname LIKE ?", "%"+q+"%").Or("lname LIKE ?", "%"+q+"%"))
	}
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
		query = query.Where(repository.DB.MySQL.Where("fname LIKE ?", "%"+q+"%").Or("lname LIKE ?", "%"+q+"%"))
	}
	query.Find(&users, &conditions).Count(&count)
	query = repository.DB.MySQL.Scopes(pagination.PaginateScope(count, &paginate)).Preload("Organization").Preload("Staff").Preload("UserGroup")
	if q != "" {
		query = query.Where(repository.DB.MySQL.Where("fname LIKE ?", "%"+q+"%").Or("lname LIKE ?", "%"+q+"%"))
	}
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
	err := repository.DB.MySQL.Preload("Organization").Preload("Staff").Preload("UserGroup").First(&user, &user).Error
	if err != nil {
		return nil, err
	}
	if user.BirthDate != nil {
		year, _, _, _, _, _ := helpers.TimeDiff(*user.BirthDate, time.Now())
		user.Age = year
	}
	return &user, nil
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
	birthDate, err := time.Parse("2014/11/23", request.BirthDate)
	if err != nil {
		fmt.Println(err.Error())
	}
	CityID := request.CityID
	if request.CityID == nil {
		CityID = nil
	}
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
		BirthDate:      &birthDate,
		FileID:         request.FileID,
		Address:        request.Address,
		Info:           request.Info,
		Introducer:     request.Introducer,
	}
	err = repository.DB.MySQL.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
