package data

import (
	"kahrersoftware.at/webskeleton/helper"
	"kahrersoftware.at/webskeleton/models"
)

// ...
var (
	MessiName   = "Lionel"
	MessiPass   = "secret"
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
	passEncrypted := helper.EncryptPassword(MessiPass)
	messi := &models.User{Name: MessiName, Pass: passEncrypted, Email: MessiEmail}
	err = ds.CreateEntity(messi)
	if err != nil {
		panic(err)
	}

	MessiID = messi.ID

	return ds
}
