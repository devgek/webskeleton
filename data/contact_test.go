package data_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kahrersoftware.at/webskeleton/data"
	"kahrersoftware.at/webskeleton/models"
	"kahrersoftware.at/webskeleton/types"
)

func TestCreateContact(t *testing.T) {
	inMemoryDS := data.NewInMemoryDatastore()

	contactAddress := &models.ContactAddress{Street: "Short Street", StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	customer := &models.Contact{OrgType: types.OrgTypeOrg, Name: "Mustermann GesmbH", ContactType: types.ContactTypeK, ContactAddresses: []models.ContactAddress{*contactAddress}}
	err := inMemoryDS.CreateEntity(customer)

	assert.Nil(t, err, "No error expected")
	expectedID := uint(2)
	assert.Equal(t, expectedID, customer.ID, "Customer id not %v", expectedID)
	assert.Equal(t, expectedID, customer.ContactAddresses[0].ContactID, "ContactAddesses.ContactID id not %v", expectedID)
}
