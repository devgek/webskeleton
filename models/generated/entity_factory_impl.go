/*
Package genmodels contains all entities and must also have a struct which implements entitymodel.EntityFactory.
*/
package genmodels

import (
	"errors"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/models"
	"strings"
)

type EntityFactoryImpl struct{}

func NewEntityFactoryImpl() EntityFactoryImpl {
	return EntityFactoryImpl{}
}

// GetEntity return entity struct by name
func (ef EntityFactoryImpl) GetEntity(entityName string) (interface{}, error) {
	switch strings.ToLower(entityName) {

	case "account":
		return &models.Account{}, nil
	case "contact":
		return &models.Contact{}, nil
	case "contactaddress":
		return &models.ContactAddress{}, nil
	case "user":
		return &models.User{}, nil

	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}

}

// GetEntitySlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetEntitySlice(entityName string) (interface{}, error) {
	switch strings.ToLower(entityName) {

	case "account":
		return &[]models.Account{}, nil
	case "contact":
		return &[]models.Contact{}, nil
	case "contactaddress":
		return &[]models.ContactAddress{}, nil
	case "user":
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
	if val, ok := entityList.([]interface{}); ok {
		for _, e := range val {
			if entity, ok := e.(entitymodel.Entity); ok {
				entityFunc(entity, params...)
			}
		}
	}
}
