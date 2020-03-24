package services

import (
	"errors"
	"github.com/devgek/webskeleton/helper"

	"log"

	"github.com/devgek/webskeleton/models"
)

//
var (
	ErrorLoginNotAllowed = &ServiceError{"msg.m0001"}
	ErrorUserNotCreated  = &ServiceError{"msg.m0002"}
	ErrorUserNotSaved    = &ServiceError{"msg.m0003"}
	ErrorUserNotDeleted  = &ServiceError{"msg.m0004"}
)

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

//LoginUser check user and pwd
func (s Services) LoginUser(username string, password string) (*models.User, error) {
	user, err := s.DS.GetUser(username)
	if err == nil {
		if err = helper.ComparePassword(user.Pass, []byte(password)); err == nil {
			return user, nil
		}
	}

	log.Println("LoginUser:", err.Error())
	return nil, ErrorLoginNotAllowed
}

//CreateUser create new user
func (s Services) CreateUser(username string, password string, email string) (*models.User, error) {
	user := &models.User{}
	user.Name = username
	user.Email = email
	var err error
	user.Pass, err = helper.EncryptPassword(password)
	if err == nil {
		user, err = s.DS.CreateUser(user)
		if err == nil {
			return user, err
		}
	}

	log.Println("CreateUser:", err.Error())
	return user, ErrorUserNotCreated
}

//UpdateUser update user data
func (s Services) UpdateUser(username string, email string, admin bool) (*models.User, error) {
	oldUser, err := s.DS.GetUser(username)
	if err != nil {
		log.Println("UpdateUser:", err.Error())
		return &models.User{}, ErrorUserNotSaved
	}

	oldUser.Email = email
	oldUser.Admin = admin

	newUser, err := s.DS.SaveUser(oldUser)
	if err == nil {
		return newUser, err
	}

	log.Println("UpdateUser:", err.Error())
	return &models.User{}, ErrorUserNotSaved
}

//DeleteUser delete user
func (s Services) DeleteUser(id uint) error {
	user := &models.User{}
	user.ID = id
	err := s.DS.DeleteEntityByID(user)
	if err == nil {
		return err
	}

	log.Println("DeleteUser:", err.Error())
	return ErrorUserNotDeleted
}

//GetAllUsers ...
func (s Services) GetAllUsers() ([]models.User, error) {
	users, err := s.DS.GetAllUser()
	if err == nil {
		return users, err
	}

	log.Println("GetAllUses:", err.Error())
	return users, err
}
