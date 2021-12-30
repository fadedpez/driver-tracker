package drivers

import (
	"context"

	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetDriver(ctx context.Context, driverID string) (*entities.Driver, error) {
	panic("implement me")
}

func (m *MockRepo) CreateDriver(driver *entities.Driver) (*entities.Driver, error) {
	args := m.Called(driver)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Driver), nil
}
