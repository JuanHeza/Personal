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

//CreateFunction is something
func (m *MockStore) CreateFunction(fn *Function, pr string) error {
	rets := m.Called(fn, pr)
	return rets.Error(0)
}

//UpdateFunction is something
func (m *MockStore) UpdateFunction(fn *Function) error {
	rets := m.Called(fn)
	return rets.Error(0)
}

//DeleteFunction is something
func (m *MockStore) DeleteFunction(fn *Function) error {
	rets := m.Called(fn)
	return rets.Error(0)
}

//CreateModel is something
func (m *MockStore) CreateModel(md *Model, pr string) error {
	rets := m.Called(md, pr)
	return rets.Error(0)
}

//UpdateModel is something
func (m *MockStore) UpdateModel(md *Model, pr string) error {
	rets := m.Called(md, pr)
	return rets.Error(0)
}

//DeleteModel is something
func (m *MockStore) DeleteModel(md *Model, pr string) error {
	rets := m.Called(md, pr)
	return rets.Error(0)
}

//CreateNotas is something
func (m *MockStore) CreateNotas(nt *Note, pr string) error {
	rets := m.Called(nt, pr)
	return rets.Error(0)
}

//UpdateNotas is something
func (m *MockStore) UpdateNotas(nt *Note) error {
	rets := m.Called(nt)
	return rets.Error(0)
}

//DeleteNotas is something
func (m *MockStore) DeleteNotas(nt *Note) error {
	rets := m.Called(nt)
	return rets.Error(0)
}

//CreateTareas is something
func (m *MockStore) CreateTareas(tr *Task, pr string) error {
	rets := m.Called(tr, pr)
	return rets.Error(0)
}

//UpdateTareas is something
func (m *MockStore) UpdateTareas(tr *Task) error {
	rets := m.Called(tr)
	return rets.Error(0)
}

//DeleteTareas is something
func (m *MockStore) DeleteTareas(tr *Task) error {
	rets := m.Called(tr)
	return rets.Error(0)
}

//InitMockStore initialice the store
func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
