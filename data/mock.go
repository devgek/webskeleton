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

//CreateEntity ...
func (m *MockedDatastore) CreateEntity(entity interface{}) error {
	args := m.Called(entity)

	return args.Error(1)
}

//SaveEntity ...
func (m *MockedDatastore) SaveEntity(entity interface{}) error {
	args := m.Called(entity)

	return args.Error(1)
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
