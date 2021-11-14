package entities

import (
	"errors"
	"time"
)

type Driver struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	dob       *time.Time
}

type DriverConfig struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	dob       *time.Time
}

func NewDriver(cfg *DriverConfig) (*Driver, error) {
	if cfg == nil {
		return nil, errors.New("missing DriverConfig")
	}

	if cfg.FirstName == "" {
		return nil, errors.New("missing FirstName")
	}

	if cfg.LastName == "" {
		return nil, errors.New("missing LastName")
	}

	if cfg.dob == nil {
		return nil, errors.New("missing DOB")
	}

	return &Driver{
		FirstName: cfg.FirstName,
		LastName:  cfg.LastName,
		dob:       cfg.dob,
	}, nil

}
