package data_test

import (
	"testing"

	"github.com/devgek/webskeleton/data"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestGetUserLionel(t *testing.T) {
	inMemoryDS, err := data.NewInMemoryDatastore()

	user, err := inMemoryDS.GetUser("Lionel")

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, data.MessiID, user.ID, "User id not %v", data.MessiID)
	assert.Equal(t, data.MessiEmail, user.Email, "Email not %v", data.MessiEmail)
}

func TestDatastore_GetUser(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("After calling NewInMemoryDatastore", t, func() {
		ds, err := data.NewInMemoryDatastore()

		Convey("Datastore must be created without an error", func() {
			So(ds, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("And user "+data.MessiName+" must be available", func() {
				messi, err := ds.GetUser(data.MessiName)
				So(err, ShouldBeNil)
				So(messi, ShouldNotBeNil)
				So(messi.Name, ShouldEqual, data.MessiName)
			})
		})
	})
}
