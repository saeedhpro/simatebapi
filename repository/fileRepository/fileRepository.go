package fileRepository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/domain/models"
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
		files[i].Path = fmt.Sprintf("http://%s/res/file/%s/1.%s", c.Request.Host, files[i].Path, files[i].Ext)
	}
	return files, nil
}

func DeleteFileBy(conditions *models.FileModel) (bool, error) {
	err := repository.DB.MySQL.Delete(&conditions, &conditions).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
