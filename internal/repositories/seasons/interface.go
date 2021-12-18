package seasons

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateSeason(season *entities.Season) (*entities.Season, error)
}
