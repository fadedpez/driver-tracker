package handlers

import (
	"context"
	"errors"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/fadedpez/driver-tracker/internal/repositories/drivers"
	"github.com/fadedpez/driver-tracker/protos"
)

const (
	requiredConfig       = "config is required"
	requiredRepo         = "repo is required"
	mockDriverCreateFail = "mock driver create failed"
	driverFirstName      = "driver first name is required"
	driverLastName       = "driver last name is required"
	driverNumber         = "driver number is required"
	driverNationality    = "driver nationality is required"
	requiredReq          = "req is required"
)

type Alpha struct {
	repo drivers.Repository
}

type AlphaConfig struct {
	Repo drivers.Repository
}

func NewAlpha(cfg *AlphaConfig) (*Alpha, error) {
	if cfg == nil {
		return nil, errors.New(requiredConfig)
	}

	if cfg.Repo == nil {
		return nil, errors.New(requiredRepo)
	}
	return &Alpha{
		repo: cfg.Repo,
	}, nil
}

func (a *Alpha) StoreDriver(ctx context.Context, req *protos.StoreDriverRequest) (*protos.StoreDriverResponse, error) {
	if req == nil {
		return nil, errors.New(requiredReq)
	}

	if req.NameFirst == "" {
		return nil, errors.New(driverFirstName)
	}

	if req.NameLast == "" {
		return nil, errors.New(driverLastName)
	}

	if req.DriverNumber == "" {
		return nil, errors.New(driverNumber)
	}

	if req.DriverNationality == "" {
		return nil, errors.New(driverNationality)
	}

	driver, err := a.repo.CreateDriver(&entities.Driver{
		FirstName:         req.NameFirst,
		LastName:          req.NameLast,
		DriverNumber:      req.DriverNumber,
		DriverNationality: req.DriverNationality,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreDriverResponse{
		Driver: &protos.Driver{
			NameFirst:         driver.FirstName,
			NameLast:          driver.LastName,
			DriverNumber:      driver.DriverNumber,
			DriverNationality: driver.DriverNationality,
		},
	}, nil
}