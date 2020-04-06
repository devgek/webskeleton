package models

import (
	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass  []byte `form:"gkvPass"`
	Email string `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Admin bool   `gorm:"not null" form:"gkvAdmin"`
}
