package models

import entitymodel "github.com/devgek/webskeleton/entity/model"

// ContactAddress ...
type ContactAddress struct {
	entitymodel.GormEntity `entity:"type:ContactAddress;name:contactaddress;gui:no;nav:no"`
	ContactID              uint
	Street                 string `gorm:"type:varchar(100);not null"`
	StreetNr               string `gorm:"type:varchar(10);not null"`
	StreetExt              string `gorm:"type:varchar(50)"`
	Zip                    string `gorm:"type:varchar(10);not null"`
	City                   string `gorm:"type:varchar(100);not null"`
}
