package teams

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateTeam(driver *entities.Team) (*entities.Team, error)
}
