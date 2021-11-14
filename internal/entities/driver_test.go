package entities

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDriver(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewDriver(nil)

		expError := errors.New("Missing DriverConfig")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})
	t.Run("it requires a first name", func(t *testing.T) {
		actual, err := NewDriver(&DriverConfig{
			FirstName: "",
			LastName: "ferguson",
		})

		expError := errors.New("missing FirstName")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)
	})
}
