package services

import (
	"unicode"

	"github.com/devgek/webskeleton/helper"
	"github.com/devgek/webskeleton/types"

	"log"

	"github.com/devgek/webskeleton/models"
)

//
var (
	ErrorLoginNotAllowed   = &ServiceError{"msg.error.login"}
	ErrorUserNotCreated    = &ServiceError{"msg.error.user.create"}
	ErrorUserNotSaved      = &ServiceError{"msg.error.user.update"}
	ErrorUserNotDeleted    = &ServiceError{"msg.error.user.delete"}
	ErrorUserPasswordRules = &ServiceError{"msg.error.user.passwordrules"}
)

//LoginUser check user and pwd
func (s Services) LoginUser(username string, password string) (*models.User, error) {
	user, err := s.DS.GetUser(username)
	if err == nil {
		if err = helper.ComparePassword(user.Pass, password); err == nil {
			return user, nil
		}
	}

	log.Println("LoginUser:", err.Error())
	return nil, ErrorLoginNotAllowed
}

//CreateUser create new user
func (s Services) CreateUser(username string, password string, email string, role types.RoleType) (*models.User, error) {
	user := &models.User{}
	user.Name = username
	user.Email = email
	user.Role = role
	user.Pass = password
	var err error

	if !isValidPassword(password) {
		log.Println("CreateUser: Password not valid")
		return user, ErrorUserPasswordRules
	}

	user.Pass = helper.EncryptPassword(password)

	if err = s.DS.CreateEntity(user); err == nil {
		return user, err
	}

	log.Println("CreateUser:", err.Error())
	return user, ErrorUserNotCreated
}

// minLen=8, 1 uppercase, 1 lowercase, 1 number
func isValidPassword(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 8 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}

//UpdateUser update user data
func (s Services) UpdateUser(username string, email string, role types.RoleType) (*models.User, error) {
	oldUser, err := s.DS.GetUser(username)
	if err != nil {
		log.Println("UpdateUser:", err.Error())
		return nil, ErrorUserNotSaved
	}

	oldUser.Email = email
	oldUser.Role = role

	if err = s.DS.SaveEntity(oldUser); err == nil {
		return oldUser, err
	}

	log.Println("UpdateUser:", err.Error())
	return nil, ErrorUserNotSaved
}
