package services

import (
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/msg"
)

//Services the business serives
type Services struct {
	DS data.Datastore
}

//NewServices ...
func NewServices(ds data.Datastore) *Services {
	return &Services{ds}
}

//ServiceError ...
type ServiceError struct {
	key string
}

func (se *ServiceError) Error() string {
	return msg.Messages.GetString(se.key)
}
