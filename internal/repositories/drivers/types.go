package drivers

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UUIDGenerator struct{}

func (u *UUIDGenerator) New() string {
	return uuid.New().String()
}

type mockUUIDGenerator struct {
	mock.Mock
}

func (m *mockUUIDGenerator) New() string {
	args := m.Called()

	return args.Get(0).(string)
}

type UUIDInterface interface {
	New() string
}
