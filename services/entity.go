package services

import (
	"errors"
	"log"
)

//
var (
	ErrorEntityNotCreated = errors.New("Entity not created")
	ErrorEntityNotSaved   = errors.New("Entity not saved")
	ErrorEntityNotDeleted = errors.New("Entity not deleted")
)

//CreateEntity create new user
func (s Services) CreateEntity(entity interface{}) error {
	err := s.DS.CreateEntity(entity)
	if err == nil {
		return err
	}

	log.Println("CreateEntity:", err.Error())
	return ErrorEntityNotCreated
}
