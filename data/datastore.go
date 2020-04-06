package data

import (
	"github.com/devgek/webskeleton/helper"
	"github.com/devgek/webskeleton/types"

	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
)

//CRUDDatastore CRUD operations with abstract entity type
type CRUDDatastore interface {
	GetOneEntityBy(entity interface{}, key string, val interface{}) error
	GetEntityByID(entity interface{}, id uint) error
	GetAllEntities(entitySlice interface{}) error
	CreateEntity(entity interface{}) error
	SaveEntity(entity interface{}) error
	DeleteEntityByID(entity interface{}, id uint) error
}

//Datastore interface to datastore
type Datastore interface {
	CRUDDatastore
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
	db.AutoMigrate(&models.User{}, &models.Contact{}, &models.ContactAddress{})

	pass, _ := helper.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Admin: true}

	// err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error
	err = db.FirstOrCreate(admin, "name = ?", "admin").Error
	if err != nil {
		return nil, err
	}

	contactAddress := &models.ContactAddress{Street: "Short Street", StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	customer := &models.Contact{OrgType: types.OrgTypeOrg, Name: "Mustermann GesmbH", CustomerType: types.CustomerTypeK, ContactAddresses: []models.ContactAddress{*contactAddress}}

	err = db.FirstOrCreate(customer, "ID = ?", 1).Error

	return &DatastoreImpl{db}, err
}
