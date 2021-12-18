package drivers

import (
	"context"
	"testing"

	"github.com/KirkDiggler/go-projects/dynamo/inputs/getitem"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/fadedpez/driver-tracker/internal/common"

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
		uuidGenerator: &common.MockUUIDGenerator{},
	}
}

func TestDynamo_GetDriver(t *testing.T) {
	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		ctx := context.Background()
		testID := "abcd"
		options := getitem.NewOptions(getitem.WithKey(map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: testID},
		}))

		m.On("GetItem", ctx, "test_table", options).Return(&getitem.Result{
			Item: map[string]types.AttributeValue{
				"id":          &types.AttributeValueMemberS{Value: testID},
				"FirstName":   &types.AttributeValueMemberS{Value: "kirk"},
				"LastName":    &types.AttributeValueMemberS{Value: "diggler"},
				"Number":      &types.AttributeValueMemberS{Value: "42"},
				"Nationality": &types.AttributeValueMemberS{Value: "USA"},
			},
		}, nil)

		actual, err := repo.GetDriver(ctx, testID)

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, testID, actual.ID)
		assert.Equal(t, "kirk", actual.FirstName)
		assert.Equal(t, "diggler", actual.LastName)
		assert.Equal(t, "42", actual.Number)
		assert.Equal(t, "USA", actual.Nationality)

	})
}

func TestDynamo_CreateDriver(t *testing.T) {
	t.Run("it calls the client properly", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*dynamo.Mock)

		muuid := repo.uuidGenerator.(*common.MockUUIDGenerator)

		inputDriver := &entities.Driver{
			FirstName:   "kirk",
			LastName:    "diggler",
			Number:      "12",
			Nationality: "USA",
			ID:          "abcd",
		}

		muuid.On("New").Return("abcd")

		options := putitem.NewOptions(putitem.WithEntity(inputDriver))

		m.On("PutItem", mock.Anything, "test_table", options).Return(&putitem.Result{}, nil)

		actual, err := repo.CreateDriver(&entities.Driver{
			FirstName:   "kirk",
			LastName:    "diggler",
			Number:      "12",
			Nationality: "USA",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)

		m.AssertExpectations(t)
		muuid.AssertExpectations(t)
	})
}
