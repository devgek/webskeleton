package services

import (
	"kahrersoftware.at/webskeleton/dtos"
	"kahrersoftware.at/webskeleton/helper"
	"kahrersoftware.at/webskeleton/models"
	"kahrersoftware.at/webskeleton/types"
)

//CreateEntity create new user
func (s Services) CreateEntity(entity interface{}, entityName string) error {
	if entityName == "user" {
		user := entity.(*models.User)
		user.Pass = helper.EncryptPassword(user.Pass)
		return s.DS.CreateEntity(user)
	}

	return s.DS.CreateEntity(entity)
}

//GetEntityOptions ...
func (s Services) GetEntityOptions(entityType types.EntityType) ([]dtos.EntityOption, error) {
	options := []dtos.EntityOption{}

	entities := s.EF.GetSlice(entityType.Val())
	err := s.DS.GetAllEntities(entities)
	if err != nil {
		return options, err
	}

	switch entityType := entities.(type) {
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
	options := []dtos.EntityOption{}

	entity := s.EF.Get(entityType.Val())
	err := s.DS.GetEntityByID(entity, id)
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
