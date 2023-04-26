package data

import (
	"github.com/devgek/webskeleton/config"
	entitydata "github.com/devgek/webskeleton/entity/data"
	"github.com/devgek/webskeleton/helper/password"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/types"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite" // gorm for sqlite3
	"gorm.io/gorm"
)

// Datastore interface to datastore
type Datastore interface {
	entitydata.EntityDatastore
	GetUser(name string) (*models.User, error)
	GetDB() *gorm.DB
}

// DatastoreImpl the Datastore implementation
type DatastoreImpl struct {
	*entitydata.GormEntityDatastore
}

// GetDB ...
func (ds DatastoreImpl) GetDB() *gorm.DB {
	return ds.DB
}

// NewPostgres ...
func NewPostgres() (Datastore, error) {
	dialectArgs := "host=" + config.DatastoreHost()
	dialectArgs = dialectArgs + " port=" + config.DatastorePort()
	dialectArgs = dialectArgs + " user=" + config.DatastoreUser()
	dialectArgs = dialectArgs + " password=" + config.DatastorePassword()
	dialectArgs = dialectArgs + " dbname=" + config.DatastoreDb()
	dialectArgs = dialectArgs + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dialectArgs), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return initNewEntityDatastore(db)
}

// NewSqlite ...
func NewSqlite(dbName string) (Datastore, error) {
	//?_foreign_keys=1 ... to handle foreign keys with golang
	db, err := gorm.Open(sqlite.Open(dbName+"?_foreign_keys=1"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return initNewEntityDatastore(db)
}

// Initialize new entity datastore with some initial data records
func initNewEntityDatastore(db *gorm.DB) (Datastore, error) {
	if config.IsDatastoreLog() {
		//log gorm db statements
		db.Debug()
	}

	//GEN-BEGIN:create entity tables
	db.AutoMigrate(&models.User{}, &models.Contact{}, &models.ContactAddress{}, &models.Account{})
	//GEN-END:create entity tables

	pass := password.EncryptPassword("xyz")
	admin := &models.User{Name: "admin", Pass: pass, Email: "admin@webskeleton.com", Role: types.RoleTypeAdmin}

	// err = db.FirstOrCreate(admin, &models.User{Name: "admin"}).Error
	err := db.FirstOrCreate(admin, "name = ?", "admin").Error
	if err != nil {
		return nil, err
	}

	contactAddress := &models.ContactAddress{Street: "Short Street", StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	contact := &models.Contact{OrgType: types.OrgTypeOrg, Name: "Mustermann GesmbH", NameExt: "Max Mustermann", ContactType: types.ContactTypeK, ContactAddresses: []models.ContactAddress{*contactAddress}}

	err = db.FirstOrCreate(contact, "name = ?", "Mustermann GesmbH").Error

	return &DatastoreImpl{GormEntityDatastore: &entitydata.GormEntityDatastore{DB: db}}, err
}
