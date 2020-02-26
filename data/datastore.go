package data

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"kahrersoftware.at/webskeleton/models"
)

//Datastore interface to datastore
type Datastore interface {
	GetUser(userID string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

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

	db.AutoMigrate(&models.User{})
	pass, _ := bcrypt.GenerateFromPassword([]byte("xyz"), bcrypt.MinCost)
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com"}

	err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error

	return &DatastoreImpl{db}, err
}

//GetUser return User data
func (ds *DatastoreImpl) GetUser(username string) (*models.User, error) {
	var user = &models.User{}
	if err := ds.Where("name = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User nicht vorhanden")
		}

		return nil, err
	}

	return user, nil
}

//CreateUser create new user
func (ds *DatastoreImpl) CreateUser(user *models.User) (*models.User, error) {
	ret := *user
	err := ds.Create(user).Error
	return &ret, err
}
