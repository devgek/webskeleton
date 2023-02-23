package genmodels

import (
	entitymodel "github.com/devgek/webskeleton/entity/model"
	"github.com/devgek/webskeleton/models"
)

type EntityFactoryCreator struct{}

func (er *EntityFactoryCreator) Create() entitymodel.EntityFactory {
	ef := entitymodel.NewDefaultEntityFactory()

	ef.RegisterType("contact", &models.Contact{})
	ef.RegisterType("contactaddress", &models.ContactAddress{})
	ef.RegisterType("user", &models.User{})

	ef.RegisterSliceType("contact", &[]models.Contact{})
	ef.RegisterSliceType("contactaddress", &[]models.ContactAddress{})
	ef.RegisterSliceType("user", &[]models.User{})

	return ef
}
