package drivers

import (
	"context"

	"github.com/fadedpez/driver-tracker/internal/entities"
)

type Repository interface {
	CreateDriver(driver *entities.Driver) (*entities.Driver, error)
	GetDriver(ctx context.Context, driverID string) (*entities.Driver, error)
}
