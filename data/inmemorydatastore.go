package data

import (
	"github.com/devgek/webskeleton/helper/password"
	"github.com/devgek/webskeleton/models"
)

// ...
var (
	MessiName          = "Lionel"
	PassSecret         = "Secret00"
	PassSecretEcrypted = password.EncryptPassword(PassSecret)
	MessiEmail         = "lionel.messi@fcb.com"
	MessiEmail2        = "lm@barcelona.es"
	MessiID            = uint(0)
)

//NewInMemoryDatastore ...
func NewInMemoryDatastore() (Datastore, error) {
	ds, err := NewDatastore("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	//init the db with test data
	messi := &models.User{Name: MessiName, Pass: PassSecretEcrypted, Email: MessiEmail}
	err = ds.CreateEntity(messi)
	if err != nil {
		panic(err)
	}

	MessiID = messi.Entity.ID

	return ds, nil
}
