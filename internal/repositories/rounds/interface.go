package rounds

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateRound(round *entities.Round) (*entities.Round, error)
}
