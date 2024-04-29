package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/alcb1310/bca-json/internal/types"
)

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

func (m *DatabaseMock) Register(reg types.RegisterInformation) (uuid.UUID, error) {
	args := m.Called(reg)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
