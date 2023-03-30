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
	//entity := ef.entityRegistry[entityName]
	//return entity, nil

	switch strings.ToLower(entityName) {
	case "contact":
		return &models.Contact{}, nil
	case "contactaddress":
		return &models.ContactAddress{}, nil
	case "user":
		return &models.User{}, nil
	case "account":
		return &models.Account{}, nil
	default:
		return nil, errors.New("Unknown entity '" + entityName + "'")
	}

}

// GetEntitySlice return slice of entity struct by name
func (ef EntityFactoryImpl) GetEntitySlice(entityName string) (interface{}, error) {
	//entitySlice := ef.entitySliceRegistry[entityName]
	//return entitySlice, nil

	switch strings.ToLower(entityName) {
	case "contact":
		return &[]models.Contact{}, nil
	case "contactaddress":
		return &[]models.ContactAddress{}, nil
	case "user":
		return &[]models.User{}, nil
	case "account":
		return &[]models.Account{}, nil
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
