package handlers

import (
	"context"
	"errors"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/fadedpez/driver-tracker/internal/repositories/drivers"
	"github.com/fadedpez/driver-tracker/internal/repositories/teams"
	"github.com/fadedpez/driver-tracker/protos"
)

const (
	requiredConfig       = "config is required"
	driverRequiredRepo   = "driverRepo is required"
	mockDriverCreateFail = "mock driver create failed"
	driverFirstName      = "driver first name is required"
	driverLastName       = "driver last name is required"
	driverNumber         = "driver number is required"
	driverNationality    = "driver nationality is required"
	requiredReq          = "req is required"
	teamRequiredRepo     = "teamRepo is required"
	teamName             = "team name is required"
	teamNationality      = "team nationality is required"
	teamPrincipal        = "team principal is required"
	teamEstablished      = "team established year is required"
)

type Alpha struct {
	driverRepo drivers.Repository
	teamRepo   teams.Repository
}

func (a *Alpha) GetTeamByName(ctx context.Context, request *protos.GetTeamByNameRequest) (*protos.GetTeamByNameResponse, error) {
	return nil, errors.New("implement me")
}

type AlphaConfig struct {
	DriverRepo drivers.Repository
	TeamRepo   teams.Repository
}

func NewAlpha(cfg *AlphaConfig) (*Alpha, error) {
	if cfg == nil {
		return nil, errors.New(requiredConfig)
	}

	if cfg.DriverRepo == nil {
		return nil, errors.New(driverRequiredRepo)
	}

	if cfg.TeamRepo == nil {
		return nil, errors.New(teamRequiredRepo)
	}
	return &Alpha{
		driverRepo: cfg.DriverRepo,
		teamRepo:   cfg.TeamRepo,
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

	driver, err := a.driverRepo.CreateDriver(&entities.Driver{
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
			Id:                driver.ID,
		},
	}, nil
}

func (a *Alpha) StoreTeam(ctx context.Context, req *protos.StoreTeamRequest) (*protos.StoreTeamResponse, error) {
	if req == nil {
		return nil, errors.New(requiredReq)
	}

	if req.TeamName == "" {
		return nil, errors.New(teamName)
	}

	if req.TeamNationality == "" {
		return nil, errors.New(teamNationality)
	}

	if req.TeamPrincipal == "" {
		return nil, errors.New(teamPrincipal)
	}

	if req.TeamEstablishedYear == "" {
		return nil, errors.New(teamEstablished)
	}

	team, err := a.teamRepo.CreateTeam(&entities.Team{
		TeamName:            req.TeamName,
		TeamNationality:     req.TeamNationality,
		TeamPrincipal:       req.TeamPrincipal,
		TeamEstablishedYear: req.TeamEstablishedYear,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreTeamResponse{
		Team: &protos.Team{
			TeamName:            team.TeamName,
			TeamNationality:     team.TeamNationality,
			TeamPrincipal:       team.TeamPrincipal,
			TeamEstablishedYear: team.TeamEstablishedYear,
			Id:                  team.ID,
		},
	}, nil
}
