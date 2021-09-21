/*
Package models contains all entities and must also have a struct which implemnents entitymodel.EntityFactory.
*/
package models

import (
	"errors"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"log"
	"strings"
)

//EntityFactoryImpl create Entities by name
type EntityFactoryImpl struct {
}

//Get return entity struct by name
func (ef EntityFactoryImpl) Get(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeContact:
		return &Contact{}, nil
	case EntityTypeContactAddress:
		return &ContactAddress{}, nil
	case EntityTypeUser:
		return &User{}, nil

	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//GetSlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetSlice(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeContact:
		return &[]Contact{}, nil
	case EntityTypeContactAddress:
		return &[]ContactAddress{}, nil
	case EntityTypeUser:
		return &[]User{}, nil

	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//DoWithAll
/*
	Method ranges over entities and calls entityFunc with each entity. You can serve parameters with each call to entityFunc.
    Attention! Maybe params should be pointers to change things outside entityFunc.
*/
func (ef EntityFactoryImpl) DoWithAll(entityList interface{}, entityFunc entitymodel.DoWithEntityFunc, params ...interface{}) {
	switch entityListType := entityList.(type) {
	case *[]Contact:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}
	case *[]ContactAddress:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}
	case *[]User:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}

	default:
		log.Println("DoWithAll::unknown entityList", entityListType)
	}
}
