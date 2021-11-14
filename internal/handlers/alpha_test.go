package handlers

import (
	"context"
	"errors"
	"github.com/fadedpez/driver-tracker/internal/repositories/drivers"
	"github.com/fadedpez/driver-tracker/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func setupFixture() *Alpha {
	return &Alpha{repo: &drivers.MockRepo{}}
}

func TestNewAlpha(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewAlpha(nil)

		expError := errors.New(requiredConfig)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})

	t.Run("it requires a repo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			Repo: nil,
		})

		expError := errors.New(requiredRepo)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})

	t.Run("it returns an alpha", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{Repo: &drivers.MockRepo{}})

		assert.NotNil(t, actual)
		assert.Nil(t, err)
		assert.NotNil(t, actual.repo)
	})
}

func TestAlpha_StoreDriver(t *testing.T) {
	t.Run("it calls the repository properly", func(t *testing.T) {

	})

	t.Run("it returns an error when the repo errors", func(t *testing.T) {
		handler := setupFixture()
		m := handler.repo.(*drivers.MockRepo)

		expErr := errors.New("mock create driver fail")

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

		expErr := errors.New("first name is required")

		actual, err := handler.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameFirst: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}
