package services

import (
	"github.com/devgek/webskeleton/entity/dto"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	genmodels "github.com/devgek/webskeleton/models/generated"
)

//GetEntityOptions ...
func (s Services) GetEntityOptions(entityType genmodels.EntityType) ([]dto.EntityOption, error) {
	options := []dto.EntityOption{}

	entities, err := s.EF.GetSlice(entityType.Val())
	if err != nil {
		return options, err
	}

	err = s.DS.GetAllEntities(entities)
	if err != nil {
		return options, err
	}

	s.EF.DoWithAll(entities, entitymodel.AddNewEntityOption, &options)

	return options, nil
}
