package teams

import (
	"context"

	"github.com/fadedpez/driver-tracker/internal/entities"
)

type Repository interface {
	CreateTeam(ctx context.Context, driver *entities.Team) (*entities.Team, error)
	SearchTeamByName(ctx context.Context, name string) ([]*entities.Team, error)
	GetTeam(ctx context.Context, teamID string) (*entities.Team, error)
}
