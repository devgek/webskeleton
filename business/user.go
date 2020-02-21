package business

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"kahrersoftware.at/bmdexport/logs"

	"kahrersoftware.at/webskeleton/models"
)

//LoginUser check user and pwd
func (ctx Context) LoginUser(username string, password string) (*models.User, error) {
	user, err := ctx.DS.GetUser(username)
	if err == nil {
		if err = bcrypt.CompareHashAndPassword(user.Pass, []byte(password)); err == nil {
			return user, nil
		}
	}

	logs.Debug("LoginUser:", err.Error())
	return nil, errors.New("Login nicht erlaubt")
}

//CreateUser create new user
func (ctx Context) CreateUser(username string, password string, email string) (*models.User, error) {
	user := &models.User{}
	user.Name = username
	user.Email = email
	var err error
	user.Pass, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err == nil {
		user, err = ctx.DS.CreateUser(user)
		if err == nil {
			return user, err
		}
	}

	logs.Debug("CreateUser:", err.Error())
	return user, errors.New("User kann nicht angelegt werden")
}
