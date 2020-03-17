package models

import (
	"github.com/jinzhu/gorm"
)

//User login user
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50);not null;unique"`
	Pass  []byte
	Email string `gorm:"type:varchar(100);not null"`
	Admin bool   `gorm:"not null"`
}
