package entityservices

import (
	"github.com/devgek/webskeleton/data/entity"

	"github.com/devgek/webskeleton/models"
)

//EntityService the service to handle entity data
type EntityService struct {
	DS entitydata.EntityDatastore
	EF models.EntityFactory
}
