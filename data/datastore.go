package data

import (
	"github.com/devgek/webskeleton/helper"

	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
)

//CRUDDatastore CRUD operations with abstract entity type
type CRUDDatastore interface {
	CreateEntity(entity interface{}) error
	SaveEntity(entity interface{}) error
	DeleteEntityByID(entity interface{}) error
}

//Datastore interface to datastore
type Datastore interface {
	CRUDDatastore
	GetAllUser() ([]models.User, error)
	GetUser(name string) (*models.User, error)
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

	db.LogMode(true)
	db.AutoMigrate(&models.User{})

	pass, _ := helper.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Admin: true}

	err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error

	return &DatastoreImpl{db}, err
}
