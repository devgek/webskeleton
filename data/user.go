package data

import (
	"github.com/devgek/webskeleton/models"
)

//GetUser return User data
func (ds *DatastoreImpl) GetUser(username string) (*models.User, error) {
	var user models.User
	if err := ds.GetOneEntityBy(&user, "name", username); err != nil {
		return nil, err
	}

	return &user, nil
}
