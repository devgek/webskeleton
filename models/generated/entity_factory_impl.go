/*
Package genmodels contains all entities and must also have a struct which implemnents entitymodel.EntityFactory.
*/
package genmodels

import (
	"errors"
	"log"

	entitymodel "github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/models"

	"strings"
)

//EntityFactoryImpl create Entities by name
type EntityFactoryImpl struct {
}

//GetEntity return entity struct by name
func (ef EntityFactoryImpl) GetEntity(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeContact:
		return &models.Contact{}, nil
	case EntityTypeContactAddress:
		return &models.ContactAddress{}, nil
	case EntityTypeUser:
		return &models.User{}, nil

	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//GetEntitySlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetEntitySlice(entityName string) (interface{}, error) {
	entityType := ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case EntityTypeContact:
		return &[]models.Contact{}, nil
	case EntityTypeContactAddress:
		return &[]models.ContactAddress{}, nil
	case EntityTypeUser:
		return &[]models.User{}, nil

	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}
}

//DoWithAllEntities
/*
	Method ranges over entities and calls entityFunc with each entity. You can serve parameters with each call to entityFunc.
    Attention! Maybe params should be pointers to change things outside entityFunc.
*/
func (ef EntityFactoryImpl) DoWithAllEntities(entityList interface{}, entityFunc entitymodel.DoWithEntityFunc, params ...interface{}) {
	switch entityListType := entityList.(type) {
	case *[]models.Contact:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}
	case *[]models.ContactAddress:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}
	case *[]models.User:
		for _, entity := range *entityListType {
			entityFunc(entity, params...)
		}

	default:
		log.Println("DoWithAll::unknown entityList", entityListType)
	}
}
