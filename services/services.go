package services

import (
	"github.com/devgek/webskeleton/data"
	entityservices "github.com/devgek/webskeleton/services/entity"
	"github.com/pkg/errors"
)

//Services the business services
type Services struct {
	ES entityservices.EntityServices
	DS data.Datastore
}

//NewServices ...
func NewServices(es entityservices.EntityServices, ds data.Datastore) *Services {
	return &Services{es, ds}
}

//Do ... just for test mocking
func (s Services) Do(par1 int, par2 int) (int, error) {
	sum := par1 + par2
	//useless, but for testing errors
	if sum < 5 {
		return -1, nil
	}

	if sum > 5 {
		return sum, errors.New("invalid: sum > 5")
	}
	return sum, nil
}

//ServiceError ...
type ServiceError struct {
	key string
}

//ServiceError implements error
func (se *ServiceError) Error() string {
	return se.key
}
