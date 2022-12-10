package entitymodel

import (
	"strconv"

	"github.com/devgek/webskeleton/entity/dto"

	"gorm.io/gorm"
)

// Entity ...
type Entity interface {
	EntityID() uint
	EntityName() string
	EntityOption() dto.EntityOption
	MustEmbed() []string
}

// EntityHolder struct that holds entities
type EntityHolder interface {
	LoadRelated(db *gorm.DB) error
}

// EntityOptionBuilder struct that can build entity options
type EntityOptionBuilder interface {
	BuildEntityOption() dto.EntityOption
}

type GormEntity struct {
	gorm.Model
}

// LoadRelatedEntities implement this method in concrete entity
func (e *GormEntity) LoadRelatedEntities(db *gorm.DB) error {
	return nil
}

// EntityID the ID of the entity
func (e GormEntity) EntityID() uint {
	return e.ID
}

// EntityName the name of the entity
func (e GormEntity) EntityName() string {
	return "GormEntity" + strconv.Itoa(int(e.ID))
}

// EntityOption ...
func (e GormEntity) EntityOption() dto.EntityOption {
	o := dto.EntityOption{}
	o.ID = e.ID
	o.Value = e.EntityName()

	return o
}

func (e GormEntity) MustEmbed() []string {
	return []string{}
}

//AddNewEntityOption
/*
	Creates a new entity option and adds it to the given entityOptionList
*/
func AddNewEntityOption(entity Entity, params ...interface{}) {
	option := entity.EntityOption()
	*(params[0].(*[]dto.EntityOption)) = append(*(params[0].(*[]dto.EntityOption)), option)
}
