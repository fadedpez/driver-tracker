package drivers

import "github.com/fadedpez/driver-tracker/internal/entities"

type Repository interface {
	CreateDriver(driver *entities.Driver) (*entities.Driver, error)
}
