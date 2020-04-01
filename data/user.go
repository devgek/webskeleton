package data

import (
	"errors"
	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
)

//GetUser return User data
func (ds *DatastoreImpl) GetUser(username string) (*models.User, error) {
	var user models.User
	if err := ds.Where("name = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User not found")
		}

		return nil, err
	}

	return &user, nil
}

//GetAllUser select * from user
func (ds *DatastoreImpl) GetAllUser() ([]models.User, error) {
	var users = []models.User{}
	ds.Find(&users)

	return users, ds.Error
}

//DeleteUser delete user object
func (ds *DatastoreImpl) DeleteUser(id uint) error {
	user := &models.User{Model: gorm.Model{ID: id}}
	err := ds.Unscoped().Delete(user).Error
	if err != nil {
		return err
	}

	if ds.RowsAffected < 1 {
		return errors.New("No user deleted")
	}

	return nil
}
