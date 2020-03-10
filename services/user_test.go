package services_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"kahrersoftware.at/webskeleton/data"
	"kahrersoftware.at/webskeleton/models"
	"kahrersoftware.at/webskeleton/services"
	"testing"
)

func TestLoginUser(t *testing.T) {
	// assert.FailNow(t, "must be implemented")
	// assert.Equal(t, 14, 15)
	// create an instance of our test object
	mockedDB := &data.MockedDatastore{}
	services := services.NewServices(mockedDB)

	passEncrypted, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.MinCost)
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
			output: 4,
			err:    nil,
		},
		"getting error": {
			input1: 2,
			input2: 4,
			output: 6,
			err:    errors.New(""),
		},
	}
	s := services.NewServices(nil)

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output, err := s.Do(test.input1, test.input2)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.output, output)
	}
}
