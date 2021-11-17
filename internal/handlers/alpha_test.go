package handlers

import (
	"context"
	"errors"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/fadedpez/driver-tracker/internal/repositories/drivers"
	"github.com/fadedpez/driver-tracker/internal/repositories/teams"
	"github.com/fadedpez/driver-tracker/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func setupFixture() *Alpha {
	return &Alpha{
		driverRepo: &drivers.MockRepo{},
		teamRepo:   &teams.MockRepo{},
	}
}

func TestNewAlpha(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewAlpha(nil)

		expError := errors.New(requiredConfig)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})

	t.Run("it requires a driverRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			DriverRepo: nil,
		})

		expError := errors.New(driverRequiredRepo)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})

	t.Run("it requires a teamRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			DriverRepo: &drivers.MockRepo{},
			TeamRepo:   nil,
		})

		expError := errors.New(teamRequiredRepo)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})

	t.Run("it returns an alpha", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			DriverRepo: &drivers.MockRepo{},
			TeamRepo:   &teams.MockRepo{},
		})

		assert.NotNil(t, actual)
		assert.Nil(t, err)
		assert.NotNil(t, actual.driverRepo)
	})
}

func TestAlpha_StoreDriver(t *testing.T) {
	t.Run("it calls the repository properly", func(t *testing.T) {
		handler := setupFixture()
		m := handler.driverRepo.(*drivers.MockRepo)

		expDriver := &entities.Driver{
			FirstName:         "kirk",
			LastName:          "diggler",
			DriverNumber:      "1",
			DriverNationality: "USA",
		}

		retDriver := &entities.Driver{
			FirstName:         "kirk",
			LastName:          "diggler",
			DriverNumber:      "1",
			DriverNationality: "USA",
			ID:                "0",
		}

		m.On("CreateDriver", expDriver).Return(retDriver, nil)

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst:         "kirk",
			NameLast:          "diggler",
			DriverNumber:      "1",
			DriverNationality: "USA",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, &protos.StoreDriverResponse{
			Driver: &protos.Driver{
				NameLast:          retDriver.LastName,
				NameFirst:         retDriver.FirstName,
				DriverNumber:      retDriver.DriverNumber,
				DriverNationality: retDriver.DriverNationality,
				Id:                retDriver.ID,
			},
		}, actual)
	})

	t.Run("it returns an error when the driverRepo errors", func(t *testing.T) {
		handler := setupFixture()
		m := handler.driverRepo.(*drivers.MockRepo)

		expErr := errors.New(mockDriverCreateFail)

		m.On("CreateDriver", mock.Anything).Return(nil, expErr)

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst:         "kirk",
			NameLast:          "diggler",
			DriverNumber:      "1",
			DriverNationality: "USA",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a first name", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(driverFirstName)

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a last name", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(driverLastName)
		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst: "Stan",
			NameLast:  "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a driver number", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(driverNumber)

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst:    "Stan",
			NameLast:     "Daniel",
			DriverNumber: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a driver nationality", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(driverNationality)

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst:         "Stan",
			NameLast:          "Daniel",
			DriverNumber:      "13",
			DriverNationality: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}

func TestAlpha_StoreTeam(t *testing.T) {
	t.Run("it calls the repository properly", func(t *testing.T) {
		handler := setupFixture()
		m := handler.teamRepo.(*teams.MockRepo)

		expTeam := &entities.Team{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
		}

		retTeam := &entities.Team{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
			ID:                  "0",
		}

		m.On("CreateTeam", expTeam).Return(retTeam, nil)

		actual, err := handler.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, &protos.StoreTeamResponse{
			Team: &protos.Team{
				TeamName:            retTeam.TeamName,
				TeamNationality:     retTeam.TeamNationality,
				TeamPrincipal:       retTeam.TeamPrincipal,
				TeamEstablishedYear: retTeam.TeamEstablishedYear,
				Id:                  retTeam.ID,
			},
		}, actual)
	})

	t.Run("it requires a team name", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(teamName)

		actual, err := handler.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a team nationality", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(teamNationality)

		actual, err := handler.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName:        "beer camp",
			TeamNationality: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a team principal", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(teamPrincipal)

		actual, err := handler.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName:        "beer camp",
			TeamNationality: "USA",
			TeamPrincipal:   "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a team established year", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New(teamEstablished)

		actual, err := handler.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

}
