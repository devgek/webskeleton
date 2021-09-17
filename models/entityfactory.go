/*
Package models contains all entities and must also have a struct which implemnents entitymodel.EntityFactory.

*/
package models

import (
	"errors"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"strings"
)

//EntityFactoryImpl create Entities by name
type EntityFactoryImpl struct {
}

//Get return entity struct by name
func (ef EntityFactoryImpl) Get(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeUser:
		return &User{}, nil
	case EntityTypeContact:
		return &Contact{}, nil
	case EntityTypeContactAddress:
		return &ContactAddress{}, nil
	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//GetSlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetSlice(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeUser:
		return &[]User{}, nil
	case EntityTypeContact:
		return &[]Contact{}, nil
	case EntityTypeContactAddress:
		return &[]ContactAddress{}, nil
	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//DoWithAll
/*
	Method ranges over entities and calls entityFunc with each entity. You can serve parameters with each call to entityFunc.
    Attention! Maybe params should be pointers to change things outside entityFunc.
*/
func (ef EntityFactoryImpl) DoWithAll(entities interface{}, entityFunc entitymodel.DoWithEntityFunc, params ...interface{}) {
	switch entityType := entities.(type) {
	case *[]User:
		for _, e := range *entityType {
			entityFunc(e, params...)
		}
	case *[]Contact:
		for _, e := range *entityType {
			entityFunc(e, params...)
		}
	case *[]ContactAddress:
		for _, e := range *entityType {
			entityFunc(e, params...)
		}
	}
}
