package models

import (
	"errors"
	"strings"

	"github.com/devgek/webskeleton/types"
)

//EntityFactory create Entities by name
type EntityFactory struct {
}

//Get return entity struct by name
func (ef EntityFactory) Get(entityName string) (interface{}, error) {
	entityType := types.ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case types.EntityTypeUser:
		return &User{}, nil
	case types.EntityTypeContact:
		return &Contact{}, nil
	case types.EntityTypeContactAddress:
		return &ContactAddress{}, nil
	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//GetSlice return slice of entity struct by name
func (ef EntityFactory) GetSlice(entityName string) (interface{}, error) {
	switch strings.ToLower(entityName) {
	case "user":
		return &[]User{}, nil
	case "contact":
		return &[]Contact{}, nil
	case "contactaddress":
		return &[]ContactAddress{}, nil
	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}
