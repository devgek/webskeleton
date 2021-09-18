package entitymodel

import (
	"github.com/devgek/webskeleton/entity/dto"
	"strconv"

	"github.com/jinzhu/gorm"
)

//Entity ...
type Entity struct {
	gorm.Model
}

//EntityHolder struct that holds entities
type EntityHolder interface {
	LoadRelated(db *gorm.DB) error
}

//EntityOptionBuilder struct that can build entity options
type EntityOptionBuilder interface {
	BuildEntityOption() dto.EntityOption
}

//LoadRelatedEntities implement this method in concrete entity
func (e *Entity) LoadRelatedEntities(db *gorm.DB) error {
	return nil
}

//BuildEntityOption ...
func (e Entity) BuildEntityOption() dto.EntityOption {
	o := dto.EntityOption{}
	o.ID = e.ID
	o.Value = "Entity with ID " + strconv.Itoa(int(e.ID))

	return o
}

//AddNewEntityOption
/*
	Creates a new entity option and adds it to the given entityOptionList
*/
func AddNewEntityOption(builder EntityOptionBuilder, params ...interface{}) {
	option := builder.BuildEntityOption()
	*(params[0].(*[]dto.EntityOption)) = append(*(params[0].(*[]dto.EntityOption)), option)
}
