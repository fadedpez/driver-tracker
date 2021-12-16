package teams

import (
	"context"
	"errors"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/getitem"
	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/fadedpez/driver-tracker/internal/common"
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

func (r *Dynamo) CreateTeam(team *entities.Team) (*entities.Team, error) {
	if team == nil {
		return nil, errors.New("team repository requires a team")
	}

	team.ID = r.uuidGenerator.New()

	_, err := r.client.PutItem(context.Background(), r.tableName,
		putitem.WithEntity(team))
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *Dynamo) GetTeam(ctx context.Context, teamID string) (*entities.Team, error) {
	if teamID == "" {
		return nil, errors.New("retrieved team ID is empty")
	}

	result, err := r.client.GetItem(ctx, r.tableName,
		getitem.WithKey(map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: teamID},
		}))
	if err != nil {
		return nil, err
	}

	team := &entities.Team{}

	err = attributevalue.UnmarshalMap(result.Item, team)
	if err != nil {
		return nil, err
	}

	return team, nil
}