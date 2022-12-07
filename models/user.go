package models

import (
	"github.com/devgek/webskeleton/entity/dto"
	"github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/types"
)

// User ...
type User struct {
	entitymodel.GormEntity `entity:"type:User;name:user"`
	Name                   string         `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass                   string         `gorm:"type:text;not null" form:"gkvPass"`
	Email                  string         `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Role                   types.RoleType `gorm:"type:integer;not null" form:"gkvRole"`
}

// EntityOption ...
func (u User) EntityOption() dto.EntityOption {
	o := dto.EntityOption{}
	o.ID = u.GormEntity.EntityId()
	o.Value = u.Name + ":" + u.Email

	return o
}
