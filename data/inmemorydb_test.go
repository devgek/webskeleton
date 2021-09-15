package data_test

import (
	"github.com/devgek/webskeleton/data"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatastoreImpl_GetUser(t *testing.T) {

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
