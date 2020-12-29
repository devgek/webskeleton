package data

import (
	"github.com/devgek/webskeleton/global"
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
	GetDB() *gorm.DB
}

//DatastoreImpl the Datastore implementation
type DatastoreImpl struct {
	*gorm.DB
}

//GetDB ...
func (ds DatastoreImpl) GetDB() *gorm.DB {
	return ds.DB
}

//NewPostgres ...
func NewPostgres() (Datastore, error) {
	dialectArgs := "host=" + global.DatastoreHost()
	dialectArgs = dialectArgs + " port=" + global.DatastorePort()
	dialectArgs = dialectArgs + " user=" + global.DatastoreUser()
	dialectArgs = dialectArgs + " password=" + global.DatastorePassword()
	dialectArgs = dialectArgs + " dbname=" + global.DatastoreDb()
	dialectArgs = dialectArgs + " sslmode=disable"

	return NewDatastore("postgres", dialectArgs)
}

//NewSqlite ...
func NewSqlite(dbName string) (Datastore, error) {
	//?_foreign_keys=1 ... to handle foreign keys with golang
	return NewDatastore("sqlite3", dbName+"?_foreign_keys=1")
}

//NewDatastore create datastore DS
func NewDatastore(driver string, databaseName string) (Datastore, error) {
	db, err := gorm.Open(driver, databaseName)
	if err != nil {
		return nil, err
	}

	if global.IsDatastoreLog() {
		//log gorm db statements
		db.LogMode(true)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Contact{}, &models.ContactAddress{})

	pass := helper.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Role: types.RoleTypeAdmin}

	// err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error
	err = db.FirstOrCreate(admin, "name = ?", "admin").Error
	if err != nil {
		return nil, err
	}

	contactAddress := &models.ContactAddress{Street: "Short Street", StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	contact := &models.Contact{OrgType: types.OrgTypeOrg, Name: "Mustermann GesmbH", NameExt: "Max Mustermann", ContactType: types.ContactTypeK, ContactAddresses: []models.ContactAddress{*contactAddress}}

	err = db.FirstOrCreate(contact, "name = ?", "Mustermann GesmbH").Error

	return &DatastoreImpl{db}, err
}
