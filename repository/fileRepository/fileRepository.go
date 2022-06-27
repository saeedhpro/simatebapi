package fileRepository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetFilesListBy(c *gin.Context, conditions *models.FileModel) ([]models.FileModel, error) {
	files := []models.FileModel{}
	err := repository.DB.MySQL.Find(&files, &conditions).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(files); i++ {
		organization := models.OrganizationModel{
			ID: files[i].OrganizationID,
		}
		_ = repository.DB.MySQL.Find(&organization, &organization).Error
		files[i].Organization = organization
		user := models.UserModel{
			ID: files[i].UserID,
		}
		_ = repository.DB.MySQL.Find(&user, &user).Error
		files[i].User = user
		user = models.UserModel{
			ID: files[i].StaffID,
		}
		_ = repository.DB.MySQL.Find(&user, &user).Error
		files[i].Staff = user
		files[i].Path = fmt.Sprintf("http://%s/file/%s/1.%s", c.Request.Host, files[i].Path, files[i].Ext)
	}
	return files, nil
}

func GetFileByID(id uint64) (*models.FileModel, error) {
	file := models.FileModel{
		ID: id,
	}
	err := repository.DB.MySQL.First(&file, &file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func CreateFile(request *requests.FileCreateRequest) (*models.FileModel, error) {
	file := models.FileModel{
		UserID:         request.UserID,
		OrganizationID: request.OrganizationID,
		StaffID:        request.StaffID,
		Ext:            request.Ext,
		Comment:        request.Comment,
		Info:           request.Info,
		Path:           request.Path,
	}
	err := repository.DB.MySQL.
		Create(&file).
		Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func DeleteFileBy(conditions *models.FileModel) (bool, error) {
	err := repository.DB.MySQL.Delete(&conditions, &conditions).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
