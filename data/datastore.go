package data

import (
	"errors"
	"github.com/devgek/webskeleton/helper"

	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
)

//Datastore interface to datastore
type Datastore interface {
	GetAllUser() ([]models.User, error)
	GetUser(name string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	SaveUser(user *models.User) (*models.User, error)
	DeleteEntityByID(entity interface{}) error
}

//
var (
	ErrorEntityNotDeleted = errors.New("Entity not deleted")
)

//DatastoreImpl the Datastore implementation
type DatastoreImpl struct {
	*gorm.DB
}

//NewDatastore create datastore DS
func NewDatastore(driver string, databaseName string) (Datastore, error) {
	db, err := gorm.Open(driver, databaseName)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.AutoMigrate(&models.User{})

	pass, _ := helper.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Admin: true}

	err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error

	return &DatastoreImpl{db}, err
}

//GetUser return User data
func (ds *DatastoreImpl) GetUser(username string) (*models.User, error) {
	var user = &models.User{}
	if err := ds.Where("name = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User not found")
		}

		return nil, err
	}

	return user, nil
}

//CreateUser create new user
func (ds *DatastoreImpl) CreateUser(user *models.User) (*models.User, error) {
	err := ds.Create(user).Error
	return user, err
}

//GetAllUser select * from user
func (ds *DatastoreImpl) GetAllUser() ([]models.User, error) {
	var users = []models.User{}
	ds.Find(&users)

	return users, ds.Error
}

//SaveUser update user data
func (ds *DatastoreImpl) SaveUser(user *models.User) (*models.User, error) {
	return user, ds.Save(user).Error
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

//DeleteEntityByID delete entity by id (primary key)
//ID must be provided
func (ds *DatastoreImpl) DeleteEntityByID(entity interface{}) error {
	db := ds.Unscoped().Delete(entity)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected != 1 {
		return ErrorEntityNotDeleted
	}

	return nil
}
