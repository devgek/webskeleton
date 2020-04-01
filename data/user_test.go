package data_test

import (
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserLionel(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	user, err := inMemoryDS.GetUser("Lionel")

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, data.MessiID, user.ID, "User id not %v", data.MessiID)
	assert.Equal(t, data.MessiEmail, user.Email, "Email not %v", data.MessiEmail)
}

func TestCreateUser(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	roger := &models.User{Name: "Roger", Pass: []byte{'s', 'e', 'c', 'r', 'e', 't'}, Email: "roger.federer@atp.com", Admin: false}
	err := inMemoryDS.CreateEntity(roger)

	assert.Nil(t, err, "No error expected")
	expectedID := data.MessiID + 1
	assert.Equal(t, expectedID, roger.ID, "User id not %v", expectedID)
}

func TestSaveUser(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	messi, err := inMemoryDS.GetUser("Lionel")

	assert.Nil(t, err, "No error expected")

	oldMessi := *messi

	messi.Email = data.MessiEmail2
	err = inMemoryDS.SaveEntity(messi)

	assert.Nil(t, err, "No error expected")
	assert.NotEqual(t, oldMessi.Email, messi.Email, "New Email not saved")
	assert.Equal(t, oldMessi.CreatedAt, messi.CreatedAt, "CreatedAt changed")
	assert.NotEqual(t, oldMessi.UpdatedAt, messi.UpdatedAt, "UpdatedAt not saved")
}
