package data

import (
	"github.com/devgek/webskeleton/config"
	entitydata "github.com/devgek/webskeleton/entity/data"
	"github.com/devgek/webskeleton/helper/password"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
)

//Datastore interface to datastore
type Datastore interface {
	entitydata.EntityDatastore
	GetUser(name string) (*models.User, error)
	GetDB() *gorm.DB
}

//DatastoreImpl the Datastore implementation
type DatastoreImpl struct {
	*entitydata.GormEntityDatastoreImpl
}

//GetDB ...
func (ds DatastoreImpl) GetDB() *gorm.DB {
	return ds.DB
}

//NewPostgres ...
func NewPostgres() (Datastore, error) {
	dialectArgs := "host=" + config.DatastoreHost()
	dialectArgs = dialectArgs + " port=" + config.DatastorePort()
	dialectArgs = dialectArgs + " user=" + config.DatastoreUser()
	dialectArgs = dialectArgs + " password=" + config.DatastorePassword()
	dialectArgs = dialectArgs + " dbname=" + config.DatastoreDb()
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

	if config.IsDatastoreLog() {
		//log gorm db statements
		db.LogMode(true)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Contact{}, &models.ContactAddress{})

	pass := password.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Role: types.RoleTypeAdmin}

	// err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error
	err = db.FirstOrCreate(admin, "name = ?", "admin").Error
	if err != nil {
		return nil, err
	}

	contactAddress := &models.ContactAddress{Street: "Short Street", StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	contact := &models.Contact{OrgType: types.OrgTypeOrg, Name: "Mustermann GesmbH", NameExt: "Max Mustermann", ContactType: types.ContactTypeK, ContactAddresses: []models.ContactAddress{*contactAddress}}

	err = db.FirstOrCreate(contact, "name = ?", "Mustermann GesmbH").Error

	return &DatastoreImpl{GormEntityDatastoreImpl: &entitydata.GormEntityDatastoreImpl{DB: db}}, err
}
