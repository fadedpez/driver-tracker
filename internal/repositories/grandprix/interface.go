package grandprix

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateGrandPrix(grandPrix *entities.GrandPrix) (*entities.GrandPrix, error)
}
