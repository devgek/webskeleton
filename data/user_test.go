package data_test

import (
	"github.com/devgek/webskeleton/data"
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
