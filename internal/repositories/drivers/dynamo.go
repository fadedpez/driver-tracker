package drivers

import (
	"context"
	"errors"

	"github.com/fadedpez/driver-tracker/internal/common"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/fadedpez/driver-tracker/internal/entities"
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
		return nil, errors.New("a config is required") //TODO: turn the strings into variables
	}

	if cfg.Client == nil {
		return nil, errors.New("a client is required")
	}

	if len(cfg.TableName) < 3 {
		return nil, errors.New("a valid table name is required")
	}

	return &Dynamo{
		client:        cfg.Client,
		tableName:     cfg.TableName,
		uuidGenerator: &common.UUIDGenerator{},
	}, nil
}

func (r *Dynamo) CreateDriver(driver *entities.Driver) (*entities.Driver, error) {
	if driver == nil {
		return nil, errors.New("driver repository requires a driver")
	}

	driver.ID = r.uuidGenerator.New()

	_, err := r.client.PutItem(context.Background(), r.tableName,
		putitem.WithEntity(driver))
	if err != nil {
		return nil, err
	}

	return driver, nil
}