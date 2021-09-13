package entityservices

import (
	"github.com/devgek/webskeleton/data/entity"

	"github.com/devgek/webskeleton/models"
)

//EntityServices the services to handle entity data
type EntityServices struct {
	DS entitydata.EntityDatastore
	EF models.EntityFactory
}
