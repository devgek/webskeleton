package data

import (
	"github.com/devgek/webskeleton/models"
	"github.com/stretchr/testify/mock"
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

//SaveUser ...
func (m *MockedDatastore) SaveUser(user *models.User) (*models.User, error) {
	args := m.Called(user)

	return args.Get(0).(*models.User), nil
}

//DeleteEntityByID ...
func (m *MockedDatastore) DeleteEntityByID(entity interface{}) error {

	args := m.Called(entity)
	return args.Error(1)
}

//GetAllUser ...
func (m *MockedDatastore) GetAllUser() ([]models.User, error) {
	// args := m.Called()

	return []models.User{}, nil
}
