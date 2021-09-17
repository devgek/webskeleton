package services

import (
	"errors"
	"github.com/devgek/webskeleton/data"
	entitymodel "github.com/devgek/webskeleton/entity/model"
)

//Services the business services
type Services struct {
	DS data.Datastore
	EF entitymodel.EntityFactory
}

//NewServices ...
func NewServices(ef entitymodel.EntityFactory, ds data.Datastore) *Services {
	return &Services{DS: ds, EF: ef}
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
