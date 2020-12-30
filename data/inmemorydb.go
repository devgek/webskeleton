package data

import (
	"github.com/devgek/webskeleton/helper/password"
	"github.com/devgek/webskeleton/models"
)

// ...
var (
	MessiName   = "Lionel"
	PassSecret  = "Secret00"
	MessiEmail  = "lionel.messi@fcb.com"
	MessiEmail2 = "lm@barcelona.es"
	MessiID     = uint(0)
)

//NewInMemoryDatastore ...
func NewInMemoryDatastore() Datastore {
	ds, err := NewDatastore("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	//init the db with test data
	passEncrypted := password.EncryptPassword(PassSecret)
	messi := &models.User{Name: MessiName, Pass: passEncrypted, Email: MessiEmail}
	err = ds.CreateEntity(messi)
	if err != nil {
		panic(err)
	}

	MessiID = messi.ID

	return ds
}
