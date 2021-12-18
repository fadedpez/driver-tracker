package tracks

import (
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateTrack(track *entities.Track) (*entities.Track, error) {
	args := m.Called(track)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Track), nil
}
