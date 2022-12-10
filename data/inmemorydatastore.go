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
	MustermannName     = "Mustermann GesmbH"
	MustermannID       = uint(1)
	MustermannStreet   = "Short Street"
)

// NewInMemoryDatastore ...
func NewInMemoryDatastore() (Datastore, error) {
	ds, err := NewSqlite(":memory:")
	if err != nil {
		return nil, err
	}
	//init the db with test data
	messi := &models.User{Name: MessiName, Pass: PassSecretEcrypted, Email: MessiEmail}
	err = ds.CreateEntity(messi)
	if err != nil {
		return nil, err
	}

	MessiID = messi.EntityID()

	return ds, nil
}
