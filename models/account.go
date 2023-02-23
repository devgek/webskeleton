package models

import (
	entitydto "github.com/devgek/webskeleton/entity/dto"
	entitymodel "github.com/devgek/webskeleton/entity/model"
)

// Account ...
type Account struct {
	entitymodel.GormEntity `entity:"type:Account;name:account"`
	Name                   string `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Nr                     string `gorm:"type:text;not null" form:"gkvNr"`
}

// EntityOption ...
func (a Account) EntityOption() entitydto.EntityOption {
	o := entitydto.EntityOption{}
	o.ID = a.GormEntity.EntityID()
	o.Value = a.Name + ":" + a.Nr

	return o
}
