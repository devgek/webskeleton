package data_test

import (
	"github.com/devgek/webskeleton/data/entity"
	"testing"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/types"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
	"github.com/stretchr/testify/assert"
)

func TestGetOneEntityBy(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	var user = models.User{}
	err := inMemoryDS.GetOneEntityBy(&user, "name", "Lionel")

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, data.MessiName, user.Name, "Expected", data.MessiName)
	assert.Equal(t, data.MessiEmail, user.Email, "Expected", data.MessiEmail)

	err = inMemoryDS.GetOneEntityBy(&user, "name", "Lionex")
	assert.NotNil(t, err, "Error expected")
	assert.Equal(t, entitydata.ErrorEntityNotFountBy, err, "ErrorEntityNotFoundBy expected")
}

func TestGetAllEntities(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	var users []models.User
	err := inMemoryDS.GetAllEntities(&users)

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, 2, len(users), "Expected %v, but got %v", 2, len(users))
}

func TestGetAllEntitiesFiltered(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	var users []models.User
	err := inMemoryDS.GetAllEntities(&users)
	inMemoryDS.GetDB().Where("name = ? AND admin = ?", "admin", false)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, 2, len(users), "Expected %v, but got %v", 2, len(users))
}
func TestCreateEntity(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	roger := &models.User{Name: "Roger", Pass: "secret", Email: "roger.federer@atp.com", Role: types.RoleTypeUser}
	err := inMemoryDS.CreateEntity(roger)

	assert.Nil(t, err, "No error expected")
	expectedID := data.MessiID + 1
	assert.Equal(t, expectedID, roger.ID, "User id not %v", expectedID)
}

func TestSaveEntity(t *testing.T) {
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

func TestDeleteEntityById(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	roger := &models.User{Name: "Roger", Pass: "secret", Email: "roger.federer@atp.com", Role: types.RoleTypeUser}
	err := inMemoryDS.CreateEntity(roger)

	assert.Nil(t, err, "No error expected")
	if err = inMemoryDS.DeleteEntityByID(roger, roger.ID); err != nil {
		t.Errorf("Error while deleting entity: %v", err)
	}

	roger, err = inMemoryDS.GetUser("Roger")
	assert.NotNil(t, err, "Error expected, cause user should be deleted")

	notExistingUser := &models.User{}
	err = inMemoryDS.DeleteEntityByID(notExistingUser, 99)
	assert.NotNil(t, err, "Error expected")
	assert.Equal(t, entitydata.ErrorEntityNotDeleted, err, "Expected dedicated error ErrorEntityNotDeleted")
}
