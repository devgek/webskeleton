package services

import (
	"github.com/devgek/webskeleton/entity/dto"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/models"
)

//GetEntityOptions ...
func (s Services) GetEntityOptions(entityType models.EntityType) ([]dto.EntityOption, error) {
	options := []dto.EntityOption{}

	entities, err := s.EF.GetSlice(entityType.Val())
	if err != nil {
		return options, err
	}

	err = s.DS.GetAllEntities(entities)
	if err != nil {
		return options, err
	}

	s.EF.DoWithAll(entities, MakeEntityOption, &options)

	return options, nil
}

func MakeEntityOption(builder entitymodel.EntityOptionBuilder, params ...interface{}) {
	option := builder.BuildEntityOption()
	*(params[0].(*[]dto.EntityOption)) = append(*(params[0].(*[]dto.EntityOption)), option)
}
