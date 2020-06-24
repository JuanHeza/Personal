package main

import (
	"github.com/stretchr/testify/mock"
)

//MockStore contains methods for inspection
type MockStore struct {
	mock.Mock
}

//CreateProyect returns the result
func (m *MockStore) CreateProyect(pr *Proyect) error {
	rets := m.Called(pr)
	return rets.Error(0)
}

//CreateProject is something
func (m *MockStore) CreateProject(pr *Projects) error {
	rets := m.Called(pr)
	return rets.Error(0)
}

//DeleteProject is something
func (m *MockStore) DeleteProject(pr string) error {
	rets := m.Called(pr)
	return rets.Error(0)
}

//GetProyect returns a proyect
// func (m *MockStore) GetProyect() ([]*Proyect, error) {
// 	rets := m.Called()
// 	return rets.Get(0).([]*Proyect), rets.Error(1)
// }
func (m *MockStore) GetProyect(id ...string) ([]*Projects, error) {
	rets := m.Called()
	return rets.Get(0).([]*Projects), rets.Error(1)
}

//InitMockStore initialice the store
func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
