package models

import "kahrersoftware.at/webskeleton/types"

//User ...
type User struct {
	Entity
	Name  string         `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass  string         `gorm:"type:text;not null" form:"gkvPass"`
	Email string         `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Role  types.RoleType `gorm:"type:integer;not null" form:"gkvRole"`
}
