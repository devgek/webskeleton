package services

import (
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

 */
func TestServices_GetEntityOptions(t *testing.T) {
	mockedDB := data.MockedDatastore{}
	ef := models.EntityFactoryImpl{}
	services := NewServices(ef, &mockedDB)

	u := []models.User{}
	users := []models.User{}
	users = append(users, models.User{Name: "Maxi", Pass: "xyz", Email: "maxi@test.at", Role: 0})
	users = append(users, models.User{Name: "Franzi", Pass: "xyz", Email: "franzi@test.at", Role: 0})
	// setup expectations
	mockedDB.On("GetAllEntities", &u).Return(users)

	// call the code we are testing
	userOptions, err := services.GetEntityOptions(models.EntityTypeUser)
	assert.Nil(t, err, "No error expected")
	assert.NotNil(t, userOptions, "Returned options expected")
	assert.Equal(t, 2, len(userOptions), "Expected 2 user options")
	assert.Equal(t, "Maxi:maxi@test.at", userOptions[0].Value, "Expected Maxi option")
	assert.Equal(t, "Franzi:franzi@test.at", userOptions[1].Value, "Expected Franzi option")

	// assert that the expectations were met
	mockedDB.AssertNumberOfCalls(t, "GetAllEntities", 1)
	mockedDB.AssertExpectations(t)
}
