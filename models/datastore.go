package models

import (
	"github.com/jinzhu/gorm"
)

//Datastore interface to datastore
type Datastore interface {
	GetUser(userID string) (User, error)
}

//DS the Datastore implementation
type DS struct {
	*gorm.DB
}

//NewDS create datastore DS
func NewDS(driver string, databaseName string) (*DS, error) {
	db, err := gorm.Open(driver, databaseName)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})
	user := &User{Name: "admin", Email: "admin@webskeleton.com"}
	db.FirstOrCreate(user, user)

	return &DS{db}, nil
}

//GetUser return User data
func (ds *DS) GetUser(userID string) (User, error) {
	var user = &User{}
	ds.Where("name = ?", userID).First(&user)
	return *user, nil
}
