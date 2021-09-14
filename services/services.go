package services

import (
	"errors"
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/types"
)

//Services the business services
type Services struct {
	EF models.EntityFactory
	DS data.Datastore
}

//NewServices ...
func NewServices(ef models.EntityFactory, ds data.Datastore) *Services {
	return &Services{ef, ds}
}

//Do ... just for test mocking
func (s Services) Do(par1 int, par2 int) (int, error) {
	sum := par1 + par2
	//useless, but for testing errors
	if sum < 5 {
		return -1, nil
	}

	if sum > 5 {
		return sum, errors.New("invalid: sum > 5")
	}
	return sum, nil
}

//GetEntityOptions ...
func (s Services) GetEntityOptions(entityType types.EntityType) ([]dtos.EntityOption, error) {
	var options []dtos.EntityOption

	entities, err := s.EF.GetSlice(entityType.Val())
	if err != nil {
		return options, err
	}

	err = s.DS.GetAllEntities(entities)
	if err != nil {
		return options, err
	}

	switch entityType := entities.(type) {
	case *[]models.User:
		for _, e := range *entityType {
			option := e.BuildEntityOption()
			options = append(options, option)
		}
	case *[]models.Contact:
		for _, e := range *entityType {
			option := e.BuildEntityOption()
			options = append(options, option)
		}
	}
	return options, nil
}

//GetEntityOptionsByID ...
func (s Services) GetEntityOptionsByID(entityType types.EntityType, id uint) ([]dtos.EntityOption, error) {
	var options []dtos.EntityOption
	entity, err := s.EF.Get(entityType.Val())
	if err != nil {
		return options, err
	}
	err = s.DS.GetEntityByID(entity, id)
	if err != nil {
		return options, err
	}
	switch entity.(type) {
	case *models.Contact:
		contact := entity.(*models.Contact)
		option := contact.BuildEntityOption()
		options = append(options, option)
	}
	return options, nil
}

//ServiceError ...
type ServiceError struct {
	key string
}

//ServiceError implements error
func (se *ServiceError) Error() string {
	return se.key
}
