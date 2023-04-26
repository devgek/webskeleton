package models

import (
	entitydto "github.com/devgek/webskeleton/entity/dto"
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/types"
)

// Contact ...
type Contact struct {
	entitymodel.GormEntity `entity:"type:Contact;name:contact;gui:no;nav:yes"`
	OrgType                types.OrgType     `gorm:"type:integer;not null" form:"gkvOrgType"`
	Name                   string            `gorm:"type:varchar(100);not null" form:"gkvName"`
	NameExt                string            `gorm:"type:varchar(100)" form:"gkvNameExt"`
	ContactType            types.ContactType `gorm:"type:integer;not null" form:"gkvContactType"`
	ContactAddresses       []ContactAddress
}

// EntityOption ...
func (c Contact) EntityOption() entitydto.EntityOption {
	o := entitydto.EntityOption{}
	o.ID = c.GormEntity.EntityID()
	o.Value = c.Name

	return o
}

// MustEmbed returns the names of the fields that must be embedded (oneToMany)
func (c Contact) MustEmbed() []string {
	return []string{"ContactAddresses"}
}
