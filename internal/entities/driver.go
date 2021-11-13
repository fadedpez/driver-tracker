package entities

import (
	"errors"
	"time"
)

type Driver struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DriverNumber string `json:"driver_number"`
	Nationality  string `json:"nationality"`
	dob          time.Time
}

type DriverConfig struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DriverNumber string `json:"driver_number"`
	Nationality  string `json:"nationality"`
	dob          time.Time
}

func NewDriver(cfg *DriverConfig) (*Driver, error) {
	if cfg == nil {
		return nil, errors.New("Missing DriverConfig")
	}

	return &Driver{}, nil
}
