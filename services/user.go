package services

import (
	"errors"

	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/devgek/webskeleton/models"
)

//
var (
	ErrorLoginNotAllowed = errors.New("Login nicht erlaubt")
	ErrorUserNotCreated  = errors.New("User kann nicht angelegt werden")
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
		if err = bcrypt.CompareHashAndPassword(user.Pass, []byte(password)); err == nil {
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
	user.Pass, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err == nil {
		user, err = s.DS.CreateUser(user)
		if err == nil {
			return user, err
		}
	}

	log.Println("CreateUser:", err.Error())
	return user, ErrorUserNotCreated
}
