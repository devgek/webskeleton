package data

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"kahrersoftware.at/webskeleton/models"
)

//MockedDatastore ...
type MockedDatastore struct {
	mock.Mock
}

//GetUser ...
func (m *MockedDatastore) GetUser(username string) (*models.User, error) {

	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

//CreateUser ...
func (m *MockedDatastore) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)

	return args.Get(0).(*models.User), nil
}

//NewInMemoryDatastore ...
func NewInMemoryDatastore() (Datastore, error) {
	ds, err := NewDatastore("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	//init the db with test data
	impl := ds.(*DatastoreImpl)
	passEncrypted, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.MinCost)
	err = impl.DB.Create(&models.User{Name: "Lionel", Pass: passEncrypted, Email: "lionel.messi@fcb.com"}).Error
	if err != nil {
		panic(err)
	}
	// impl.DB.Create(&models.User{Name: "Gerald", Pass: passEncrypted, Email: "gerald.kahrer@gmail.com"})

	return ds, err
}
