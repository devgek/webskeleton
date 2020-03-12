package data

import (
	"golang.org/x/crypto/bcrypt"
	"kahrersoftware.at/webskeleton/models"
)

//NewInMemoryDatastore ...
func NewInMemoryDatastore() (Datastore, error) {
	ds, err := NewDatastore("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	//init the db with test data
	impl := ds.(*DatastoreImpl)
	passEncrypted, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.MinCost)
	err = impl.DB.Create(&models.User{Name: "Lionel", Pass: passEncrypted, Email: "lionel.messi@fcb.com"}).Error
	if err != nil {
		panic(err)
	}
	// impl.DB.Create(&models.User{Name: "Gerald", Pass: passEncrypted, Email: "gerald.kahrer@gmail.com"})

	return ds, err
}
