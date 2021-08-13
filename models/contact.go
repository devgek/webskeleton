package models

import (
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/types"
	"github.com/jinzhu/gorm"
)

//Contact ...
type Contact struct {
	Entity
	OrgType          types.OrgType     `gorm:"type:integer;not null" form:"gkvOrgType"`
	Name             string            `gorm:"type:varchar(100);not null" form:"gkvName"`
	NameExt          string            `gorm:"type:varchar(100)" form:"gkvNameExt"`
	ContactType      types.ContactType `gorm:"type:integer;not null" form:"gkvContactType"`
	ContactAddresses []ContactAddress
}

//BuildEntityOption ...
func (c *Contact) BuildEntityOption() dtos.EntityOption {
	o := dtos.EntityOption{}
	o.ID = c.ID
	o.Value = c.Name

	return o
}

//LoadRelated load related entities (implements EntityHolder)
func (c *Contact) LoadRelated(db *gorm.DB) error {
	c.ContactAddresses = []ContactAddress{}
	db.Model(c).Related(&c.ContactAddresses)

	return nil
}
