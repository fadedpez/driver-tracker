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
	teamNotFound         = "team not found"
	teamIDNotFound       = "team ID not found"
)

type Alpha struct {
	driverRepo drivers.Repository
	teamRepo   teams.Repository
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
		FirstName:   req.NameFirst,
		LastName:    req.NameLast,
		Number:      req.DriverNumber,
		Nationality: req.DriverNationality,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreDriverResponse{
		Driver: &protos.Driver{
			NameFirst:         driver.FirstName,
			NameLast:          driver.LastName,
			DriverNumber:      driver.Number,
			DriverNationality: driver.Nationality,
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

	team, err := a.teamRepo.CreateTeam(ctx, &entities.Team{
		Name:            req.TeamName,
		Nationality:     req.TeamNationality,
		Principal:       req.TeamPrincipal,
		EstablishedYear: req.TeamEstablishedYear,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreTeamResponse{
		Team: teamToProto(team),
	}, nil
}

func (a *Alpha) SearchTeamByName(ctx context.Context, req *protos.SearchTeamByNameRequest) (*protos.SearchTeamByNameResponse, error) {
	if req == nil {
		return nil, errors.New(requiredReq)
	}

	if req.TeamName == "" {
		return nil, errors.New(teamName)
	}

	team, err := a.teamRepo.SearchTeamByName(ctx, req.TeamName)
	if err != nil {
		return nil, err
	}

	return &protos.SearchTeamByNameResponse{
		Teams: teamsToProtos(team),
	}, nil
}

func teamsToProtos(teams []*entities.Team) []*protos.Team {
	if teams == nil {
		return nil
	}

	output := make([]*protos.Team, 0)

	for _, team := range teams {
		output = append(output, teamToProto(team))
	}

	return output
}

func (a *Alpha) GetTeam(ctx context.Context, req *protos.GetTeamRequest) (*protos.GetTeamResponse, error) {
	if req == nil {
		return nil, errors.New(requiredReq)
	}

	if req.TeamID == "" {
		return nil, errors.New(teamIDNotFound)
	}

	team, err := a.teamRepo.GetTeam(ctx, req.TeamID)
	if err != nil {
		return nil, err
	}

	return &protos.GetTeamResponse{
		Team: teamToProto(team),
	}, nil

}

func teamToProto(team *entities.Team) *protos.Team {
	if team == nil {
		return nil
	}

	return &protos.Team{
		Id:                  team.ID,
		TeamName:            team.Name,
		TeamNationality:     team.Nationality,
		TeamPrincipal:       team.Principal,
		TeamEstablishedYear: team.EstablishedYear,
	}
}

//TODO: fill out the functions below this comment!

//TODO: Probably have to create (x)ToProto / (x)ToProtos for Driver, GrandPrix, Quali, Round, Season, Track

func (a *Alpha) GetDriver(ctx context.Context, req *protos.GetDriverRequest) (*protos.GetDriverResponse, error) {
	if req == nil {
		return nil, errors.New(requiredReq)
	}

	if req.Id == "" {
		return nil, errors.New("GetDriverRequest.ID is empty")
	}

	driver, err := a.driverRepo.GetDriver(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &protos.GetDriverResponse{ //TODO: turn this into a function like teams
		Driver: &protos.Driver{
			Id:                driver.ID,
			NameLast:          driver.LastName,
			NameFirst:         driver.FirstName,
			DriverNumber:      driver.Number,
			DriverNationality: driver.Nationality,
		},
	}, nil
}

func (a *Alpha) SearchDriver(ctx context.Context, req *protos.SearchDriverRequest) (*protos.SearchDriverResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) StoreGrandPrix(ctx context.Context, req *protos.StoreGrandPrixRequest) (*protos.StoreGrandPrixResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) GetGrandPrix(ctx context.Context, req *protos.GetGrandPrixRequest) (*protos.GetGrandPrixResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) SearchGrandPrix(ctx context.Context, req *protos.SearchGrandPrixRequest) (*protos.SearchGrandPrixResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) StoreQualifying(ctx context.Context, req *protos.StoreQualifyingRequest) (*protos.StoreQualifyingResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) GetQualifying(ctx context.Context, req *protos.GetQualifyingRequest) (*protos.GetQualifyingResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) SearchQualifying(ctx context.Context, req *protos.SearchQualifyingRequest) (*protos.SearchQualifyingResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) StoreRound(ctx context.Context, req *protos.StoreRoundRequest) (*protos.StoreRoundResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) GetRound(ctx context.Context, req *protos.GetRoundRequest) (*protos.GetRoundResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) SearchRound(ctx context.Context, req *protos.SearchRoundRequest) (*protos.SearchRoundResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) StoreSeason(ctx context.Context, req *protos.StoreSeasonRequest) (*protos.StoreSeasonResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) GetSeason(ctx context.Context, req *protos.GetSeasonRequest) (*protos.GetSeasonResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) SearchSeason(ctx context.Context, req *protos.SearchSeasonRequest) (*protos.SearchSeasonResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) StoreTrack(ctx context.Context, req *protos.StoreTrackRequest) (*protos.StoreTrackResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) GetTrack(ctx context.Context, req *protos.GetTrackRequest) (*protos.GetTrackResponse, error) {
	return nil, errors.New("not yet implemented")
}

func (a *Alpha) SearchTrack(ctx context.Context, req *protos.SearchTrackRequest) (*protos.SearchTrackResponse, error) {
	return nil, errors.New("not yet implemented")
}
