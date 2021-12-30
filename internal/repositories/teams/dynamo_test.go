package teams

import (
	"context"
	"testing"

	"github.com/KirkDiggler/go-projects/dynamo/inputs/getitem"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/fadedpez/driver-tracker/internal/common"
	"github.com/fadedpez/driver-tracker/internal/entities"
	"github.com/stretchr/testify/assert"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client:        &dynamo.Mock{},
		tableName:     "test_table",
		uuidGenerator: &common.MockUUIDGenerator{},
	}
}

func TestDynamo_GetTeam(t *testing.T) {
	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		ctx := context.Background()
		testID := "efgh"
		options := getitem.NewOptions(getitem.WithKey(map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: testID},
		}))

		m.On("GetItem", ctx, "test_table", options).Return(&getitem.Result{
			Item: map[string]types.AttributeValue{
				"id":              &types.AttributeValueMemberS{Value: testID},      //the keys need to match the struct in team.go
				"Name":            &types.AttributeValueMemberS{Value: "beer camp"}, //key matching struct is not case sensitive, I matched for consistency
				"Nationality":     &types.AttributeValueMemberS{Value: "USA"},
				"Principal":       &types.AttributeValueMemberS{Value: "mongo"},
				"EstablishedYear": &types.AttributeValueMemberS{Value: "2015"},
			},
		}, nil)

		actual, err := repo.GetTeam(ctx, testID)

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, testID, actual.ID)
		assert.Equal(t, "beer camp", actual.Name)
		assert.Equal(t, "USA", actual.Nationality)
		assert.Equal(t, "mongo", actual.Principal)
		assert.Equal(t, "2015", actual.EstablishedYear)

	})
}

func TestDynamo_CreateTeam(t *testing.T) {
	ctx := context.Background()

	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		muuid := repo.uuidGenerator.(*common.MockUUIDGenerator)

		inputTeam := &entities.Team{
			Name:            "beer camp",
			Nationality:     "USA",
			Principal:       "mongo",
			EstablishedYear: "2015",
			ID:              "efgh",
		}

		muuid.On("New").Return("efgh")

		options := putitem.NewOptions(putitem.WithEntity(inputTeam))

		m.On("PutItem", ctx, "test_table", options).Return(&putitem.Result{}, nil)

		actual, err := repo.CreateTeam(ctx, &entities.Team{
			Name:            "beer camp",
			Nationality:     "USA",
			Principal:       "mongo",
			EstablishedYear: "2015",
			ID:              "efgh",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		m.AssertExpectations(t)
		muuid.AssertExpectations(t)

	})
}
