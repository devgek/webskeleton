package services

import (
	"github.com/devgek/webskeleton/helper"
	"github.com/devgek/webskeleton/models"
)

//CreateEntity create new user
func (s Services) CreateEntity(entity interface{}, entityName string) error {
	if entityName == "user" {
		user := entity.(*models.User)
		user.Pass = helper.EncryptPassword(user.Pass)
		return s.DS.CreateEntity(user)
	}

	return s.DS.CreateEntity(entity)
}
