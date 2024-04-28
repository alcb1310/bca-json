package mocks

import "github.com/stretchr/testify/mock"

type DatabaseMock struct {
	mock.Mock
}

func NewDatabaseMock() *DatabaseMock {
	return &DatabaseMock{}
}

func (m *DatabaseMock) Health() error {
	args := m.Called()
	return args.Error(0)
}

func (m *DatabaseMock) CreateSchema() error {
	return nil
}
