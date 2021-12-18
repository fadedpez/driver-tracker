package tracks

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateTrack(track *entities.Track) (*entities.Track, error)
}
