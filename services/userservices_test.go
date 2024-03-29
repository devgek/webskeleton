package services_test

import (
	"errors"

	entitydata "github.com/devgek/webskeleton/entity/data"
	genmodels "github.com/devgek/webskeleton/models/generated"

	"testing"

	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/helper/password"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/services"
	"github.com/devgek/webskeleton/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_ "gorm.io/driver/sqlite" // gorm for sqlite3
)

/*
func init() {
	fmt.Println("test init")
	fmt.Println(os.Getwd())
	os.Chdir("..")
	fmt.Println(os.Getwd())
}
*/

// TestLoginUser test login service with mocking Datastore
func TestLoginUser(t *testing.T) {
	// create an instance of the mocked Datastore
	mockedDB := &data.MockedDatastore{}
	ef := genmodels.NewEntityFactoryImpl()
	services := services.NewServices(ef, mockedDB)

	passEncrypted := password.EncryptPassword("secret")
	userGerald := &models.User{Name: "Gerald", Pass: passEncrypted, Email: "gerald.kahrer@gmail.com"}
	// setup expectations
	mockedDB.On("GetUser", "Gerald").Return(userGerald, nil)

	// call the code we are testing
	user, err := services.LoginUser("Gerald", "secret")
	assert.Nil(t, err, "Login not allowed")
	assert.Equal(t, userGerald, user, "Wrong user")

	// assert that the expectations were met
	mockedDB.AssertExpectations(t)
}

// TestLoginUserInMemory test login service with inmemory db
func TestLoginUserInMemoryOK(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()
	ef := genmodels.NewEntityFactoryImpl()
	services := services.NewServices(ef, inMemoryDS)

	// happy test, user with correct password
	user, err := services.LoginUser("Lionel", data.PassSecret)
	assert.Nil(t, err, "Login user with error")
	assert.NotNil(t, user, "User is nil")
}
func TestLoginUserInMemoryNOK(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, inMemoryDS)

	// user with wrong password
	user, err := s.LoginUser("Lionel", "wrongsecret")
	assert.Nil(t, user, "User expected nil")
	assert.NotNil(t, err, "Error expected")
	assert.Equal(t, services.ErrorLoginNotAllowed, err, "Expected ErrorLoginNotAllowed")
}

func TestCreateUser(t *testing.T) {
	// create an instance of the mocked Datastore
	mockedDB := &data.MockedDatastore{}
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, mockedDB)

	passEncrypted := password.EncryptPassword(data.PassSecret)
	userRoger := &models.User{Name: "Roger", Pass: passEncrypted, Email: "roger.federer@atp.com", Role: types.RoleTypeUser}
	// setup expectations
	mockedDB.On("CreateEntity", mock.Anything).Return(nil)

	userReturned, err := s.CreateUser("Roger", data.PassSecret, "roger.federer@atp.com", types.RoleTypeUser)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, userRoger.Name, userReturned.Name, "Expected name Roger")
	assert.Equal(t, userRoger.Email, userReturned.Email, "Expected email not returned")
	assert.Equal(t, userRoger.Role, userReturned.Role, "Expected admin not returned")

	mockedDB.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, inMemoryDS)

	messi, err := s.DS.GetUser("Lionel")
	assert.Nil(t, err, "No error expected")
	messi.Email = "lm@barcelona.es"
	messi2, err := s.UpdateUser(messi.Name, messi.Email, messi.Role)
	assert.Equal(t, messi.Email, messi2.Email, "Email not expected")
}

func TestDeleteUser(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, inMemoryDS)

	user, err := s.CreateUser("Rafa", data.PassSecret, "rafael.nadal@atp.com", types.RoleTypeUser)
	assert.Nil(t, err, "No error expected")
	err = s.DS.DeleteEntityByID(user, user.GormEntity.ID)
	assert.Nil(t, err, "No error expected")
}

func TestDeleteUserError(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, inMemoryDS)

	err = s.DS.DeleteEntityByID(&models.User{}, 99)
	assert.Equal(t, entitydata.ErrorEntityNotDeleted, err, "Expected error not returned")
}
func TestDoTableBased(t *testing.T) {
	tests := map[string]struct {
		input1 int
		input2 int
		output int
		err    error
	}{
		"successful addition": {
			input1: 2,
			input2: 3,
			output: 5,
			err:    nil,
		},
		"invalid addition": {
			input1: 2,
			input2: 2,
			output: -1,
			err:    nil,
		},
		"getting error": {
			input1: 2,
			input2: 4,
			output: 6,
			err:    errors.New("invalid: sum > 5"),
		},
	}
	ef := genmodels.NewEntityFactoryImpl()
	s := services.NewServices(ef, nil)

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output, err := s.Do(test.input1, test.input2)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.output, output)
	}
}
