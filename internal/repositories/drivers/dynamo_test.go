package drivers

import (
	"testing"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client:        &dynamo.Mock{},
		tableName:     "test_table", //TODO: move to constant later
		uuidGenerator: &mockUUIDGenerator{},
	}
}

func TestDynamo_CreateDriver(t *testing.T) {
	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		muuid := repo.uuidGenerator.(*mockUUIDGenerator)

		inputDriver := &entities.Driver{
			FirstName:         "kirk",
			LastName:          "diggler",
			DriverNumber:      "12",
			DriverNationality: "USA",
			ID:                "abcd",
		}

		muuid.On("New").Return("abcd")

		options := putitem.NewOptions(putitem.WithEntity(inputDriver))

		m.On("PutItem", mock.Anything, "test_table", options).Return(&putitem.Result{}, nil)

		actual, err := repo.CreateDriver(&entities.Driver{
			FirstName:         "kirk",
			LastName:          "diggler",
			DriverNumber:      "12",
			DriverNationality: "USA",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)

		m.AssertExpectations(t)
		muuid.AssertExpectations(t)
	})
}
