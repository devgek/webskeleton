package models

import (
	"kahrersoftware.at/webskeleton/dtos"
	"kahrersoftware.at/webskeleton/types"
)

//Contact ...
type Contact struct {
	Entity
	OrgType          types.OrgType      `gorm:"type:integer;not null" form:"gkvOrgType"`
	Name             string             `gorm:"type:varchar(100);not null" form:"gkvName"`
	NameExt          string             `gorm:"type:varchar(100)" form:"gkvNameExt"`
	CustomerNr       string             `gorm:"type:varchar(10)" form:"gkvCustomerNr"`
	CustomerType     types.CustomerType `gorm:"type:integer;not null" form:"gkvCustomerType"`
	ContactAddresses []ContactAddress
}

//BuildEntityOption ...
func (c *Contact) BuildEntityOption() dtos.EntityOption {
	o := dtos.EntityOption{}
	o.ID = c.ID
	o.Value = c.Name

	return o
}
