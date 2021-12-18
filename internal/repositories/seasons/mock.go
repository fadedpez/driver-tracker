package seasons

import (
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateSeason(season *entities.Season) (*entities.Season, error) {
	args := m.Called(season)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Season), nil
}
