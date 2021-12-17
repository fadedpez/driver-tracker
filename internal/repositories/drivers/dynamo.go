package drivers

import (
	"context"
	"errors"

	"github.com/KirkDiggler/go-projects/dynamo/inputs/getitem"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/fadedpez/driver-tracker/internal/common"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/fadedpez/driver-tracker/internal/entities"
)

const (
	configRequired     = "a config is required"
	clientRequired     = "a client is required"
	tableNameRequired  = "a valid table name is required"
	driverRepoRequired = "a driver repository is required"
	driverIDEmpty      = "retrieved driver ID is empty"
)

type Dynamo struct {
	client        dynamo.Interface
	tableName     string
	uuidGenerator common.UUIDInterface
}

type DynamoConfig struct {
	Client    dynamo.Interface
	TableName string
}

func NewDynamo(cfg *DynamoConfig) (*Dynamo, error) {
	if cfg == nil {
		return nil, errors.New(configRequired)
	}

	if cfg.Client == nil {
		return nil, errors.New(clientRequired)
	}

	if len(cfg.TableName) < 3 {
		return nil, errors.New(tableNameRequired)
	}

	return &Dynamo{
		client:        cfg.Client,
		tableName:     cfg.TableName,
		uuidGenerator: &common.UUIDGenerator{},
	}, nil
}

func (r *Dynamo) CreateDriver(driver *entities.Driver) (*entities.Driver, error) {
	if driver == nil {
		return nil, errors.New(driverRepoRequired)
	}

	driver.ID = r.uuidGenerator.New()

	_, err := r.client.PutItem(context.Background(), r.tableName,
		putitem.WithEntity(driver))
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (r *Dynamo) GetDriver(ctx context.Context, driverID string) (*entities.Driver, error) {
	if driverID == "" {
		return nil, errors.New(driverIDEmpty)
	}

	result, err := r.client.GetItem(ctx, r.tableName,
		getitem.WithKey(map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: driverID},
		}))
	if err != nil {
		return nil, err
	}

	driver := &entities.Driver{}

	err = attributevalue.UnmarshalMap(result.Item, driver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
