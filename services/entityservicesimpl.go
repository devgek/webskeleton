package services

import (
	"github.com/devgek/webskeleton/entity/dto"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	genmodels "github.com/devgek/webskeleton/models/generated"
)

//GetEntityOptions ...
func (s Services) GetEntityOptions(entityType genmodels.EntityType) ([]dto.EntityOption, error) {
	options := []dto.EntityOption{}

	entities, err := s.EF.GetEntitySlice(entityType.Val())
	if err != nil {
		return options, err
	}
	entity, err := s.EF.GetEntity(entityType.Val())
	if err != nil {
		return options, err
	}

	err = s.DS.GetAllEntities(entity, entities)
	if err != nil {
		return options, err
	}

	s.EF.DoWithAllEntities(entities, entitymodel.AddNewEntityOption, &options)

	return options, nil
}
