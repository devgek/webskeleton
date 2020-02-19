package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Datastore interface to datastore
type Datastore interface {
	GetUser(userID string) (*User, error)
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

	db.AutoMigrate(&User{})
	user := &User{Name: "admin", Pass: []byte("xyz"), Email: "admin@webskeleton.com"}
	db.FirstOrCreate(user, user)

	return &DatastoreImpl{db}, nil
}

//GetUser return User data
func (ds *DatastoreImpl) GetUser(userID string) (*User, error) {
	var user = &User{}
	if err := ds.Where("name = ?", userID).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User nicht vorhanden")
		}

		return nil, err
	}

	return user, nil
}
