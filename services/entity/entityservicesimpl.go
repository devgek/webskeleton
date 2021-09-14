package entityservices

import (
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/models"
	"github.com/devgek/webskeleton/types"
)

//GetEntityOptions ...
func (s EntityService) GetEntityOptions(entityType types.EntityType) ([]dtos.EntityOption, error) {
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
func (s EntityService) GetEntityOptionsByID(entityType types.EntityType, id uint) ([]dtos.EntityOption, error) {
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
