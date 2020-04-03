package models

import (
	"github.com/jinzhu/gorm"
)

//ContactAddress ...
type ContactAddress struct {
	gorm.Model
	ContactID uint
	Street    string `gorm:"type:varchar(100);not null"`
	StreetNr  string `gorm:"type:varchar(10);not null"`
	StreetExt string `gorm:"type:varchar(50)"`
	Zip       string `gorm:"type:varchar(10)"`
	City      string `gorm:"not null"`
}
