package services

import (
	"errors"
	"github.com/devgek/webskeleton/data"
	"github.com/devgek/webskeleton/models"
)

//Services the business serives
type Services struct {
	DS data.Datastore
	EF models.EntityFactory
}

//NewServices ...
func NewServices(ds data.Datastore) *Services {
	return &Services{ds, models.EntityFactory{}}
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
