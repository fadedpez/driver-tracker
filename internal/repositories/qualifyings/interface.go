package qualifyings

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateQualifying(qualifying *entities.Qualifying) (*entities.Qualifying, error)
}
