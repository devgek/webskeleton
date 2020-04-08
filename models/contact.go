package models

import (
	"github.com/devgek/webskeleton/types"
	"github.com/jinzhu/gorm"
)

//Contact ...
type Contact struct {
	gorm.Model
	OrgType          types.OrgType      `gorm:"type:char(1);not null" form:"gkvOrgType"`
	Name             string             `gorm:"type:varchar(100);not null" form:"gkvName"`
	NameExt          string             `gorm:"type:varchar(100)" form:"gkvNameExt"`
	CustomerNr       string             `gorm:"type:varchar(10)" form:"gkvCustomerNr"`
	CustomerType     types.CustomerType `gorm:"type:char(1);not null" form:"gkvCustomerType"`
	ContactAddresses []ContactAddress
}
