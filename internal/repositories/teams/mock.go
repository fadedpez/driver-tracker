package teams

import (
	"context"

	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateTeam(ctx context.Context, team *entities.Team) (*entities.Team, error) {
	args := m.Called(ctx, team)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Team), nil
}

func (m *MockRepo) SearchTeamByName(ctx context.Context, name string) ([]*entities.Team, error) {
	args := m.Called(ctx, name)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*entities.Team), nil
}

func (m *MockRepo) GetTeam(ctx context.Context, name string) (*entities.Team, error) {
	args := m.Called(ctx, name)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Team), nil
}
