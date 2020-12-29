package models

import (
	"strconv"

	"github.com/devgek/webskeleton/dtos"
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

//LoadRelatedEntities implement this method in concrete entity
func (e *Entity) LoadRelatedEntities(db *gorm.DB) error {
	return nil
}

//BuildEntityOption ...
func (e *Entity) BuildEntityOption() dtos.EntityOption {
	o := dtos.EntityOption{}
	o.ID = e.ID
	o.Value = "Entity with ID " + strconv.Itoa(int(e.ID))

	return o
}
