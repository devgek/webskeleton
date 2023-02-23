package entitymodel

import (
	"strconv"

	entitydto "github.com/devgek/webskeleton/entity/dto"
	"gorm.io/gorm"
)

// Entity ...
type Entity interface {
	EntityID() uint
	EntityDesc() string
	EntityOption() entitydto.EntityOption
	MustEmbed() []string
}

// EntityOptionBuilder struct that can build entity options
type EntityOptionBuilder interface {
	BuildEntityOption() entitydto.EntityOption
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

// EntityDesc the description of the entity
func (e GormEntity) EntityDesc() string {
	return "GormEntity" + strconv.Itoa(int(e.ID))
}

// EntityOption ...
func (e GormEntity) EntityOption() entitydto.EntityOption {
	o := entitydto.EntityOption{}
	o.ID = e.ID
	o.Value = e.EntityDesc()

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
	*(params[0].(*[]entitydto.EntityOption)) = append(*(params[0].(*[]entitydto.EntityOption)), option)
}
