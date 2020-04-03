package models

import (
	"github.com/devgek/webskeleton/types"
	"github.com/jinzhu/gorm"
)

//Contact ...
type Contact struct {
	gorm.Model
	OrgType          types.OrgType      `gorm:"not null"`
	Name             string             `gorm:"type:varchar(100);not null"`
	NameExt          string             `gorm:"type:varchar(100)"`
	CustomerNr       string             `gorm:"type:varchar(10)"`
	CustomerType     types.CustomerType `gorm:"not null"`
	ContactAddresses []ContactAddress
}
