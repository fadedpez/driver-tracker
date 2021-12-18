package qualifyings

import (
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateQualifying(qualifying *entities.Qualifying) (*entities.Qualifying, error) {
	args := m.Called(qualifying)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Qualifying), nil
}
