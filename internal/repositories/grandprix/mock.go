package grandprix

import (
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateGrandPrix(grandPrix *entities.GrandPrix) (*entities.GrandPrix, error) {
	args := m.Called(grandPrix)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.GrandPrix), nil
}
