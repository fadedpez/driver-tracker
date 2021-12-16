package teams

import (
	"testing"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/fadedpez/driver-tracker/internal/common"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client:        &dynamo.Mock{},
		tableName:     "test_table",
		uuidGenerator: &common.MockUUIDGenerator{},
	}
}

func TestDynamo_CreateTeam(t *testing.T) {
	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		muuid := repo.uuidGenerator.(*common.MockUUIDGenerator)

		inputTeam := &entities.Team{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
			ID:                  "efgh",
		}

		muuid.On("New").Return("efgh")

		options := putitem.NewOptions(putitem.WithEntity(inputTeam))

		m.On("PutItem", mock.Anything, "test_table", options).Return(&putitem.Result{}, nil)

		actual, err := repo.CreateTeam(&entities.Team{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
			ID:                  "efgh",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		m.AssertExpectations(t)
		muuid.AssertExpectations(t)

	})
}
