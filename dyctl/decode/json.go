package decode

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ParameterPair struct {
	Key   string `json:"ParameterKey"`
	Value string `json:"ParameterValue"`
}

type ParameterSet struct {
	Parameters []ParameterPair
}

func (ps *ParameterSet) GetString(key string) string {
	for _, p := range ps.Parameters {
		if p.Key == key {
			return p.Value
		}
	}

	return ""
}

func (ps *ParameterSet) indexName(schemas []*dynamodb.KeySchemaElement) (string, error) {
	if len(schemas) == 0 {
		return "", fmt.Errorf("invalid schema list")
	}

	indexName := fmt.Sprintf("%s-index", *schemas[0].AttributeName)

	if len(schemas) > 1 {
		indexName = fmt.Sprintf("%s-%s-index", *schemas[0].AttributeName, *schemas[1].AttributeName)
	}

	return indexName, nil
}

func (ps *ParameterSet) tableName(app, env, table string) string {
	return fmt.Sprintf("%s-%s-%s", app, env, table)
}

func (ps *ParameterSet) CreateTableInput(app, env, table string) (*dynamodb.CreateTableInput, error) {
	if app == "" {
		return nil, fmt.Errorf("invalid: AppName")
	}

	if env == "" {
		return nil, fmt.Errorf("invalid: AppEnvironment")
	}

	if table == "" {
		return nil, fmt.Errorf("invalid: TableName")
	}

	tableHashKey := ps.GetString("TableHashKeyName")

	if tableHashKey == "" {
		return nil, fmt.Errorf("invalid: TableHashKeyName")
	}

	var attrDefs []*dynamodb.AttributeDefinition

	attrDefMap := make(map[string]*dynamodb.AttributeDefinition)

	for {
		i := len(attrDefs) + 1

		attrDefName := ps.GetString(fmt.Sprintf("Attribute%dName", i))

		if attrDefName == "" {
			break
		}

		attrDefType := ps.GetString(fmt.Sprintf("Attribute%dType", i))

		if attrDefType == "" {
			return nil, fmt.Errorf("invalid: %s", fmt.Sprintf("Attribute%dType", i))
		}

		attrDef := &dynamodb.AttributeDefinition{
			AttributeName: aws.String(attrDefName),
			AttributeType: aws.String(attrDefType),
		}

		attrDefs = append(attrDefs, attrDef)
		attrDefMap[attrDefName] = attrDef
	}

	if _, found := attrDefMap[tableHashKey]; !found {
		return nil, fmt.Errorf("invalid: no associated attribute for %s", tableHashKey)
	}

	tableSchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String(tableHashKey),
			KeyType:       aws.String("HASH"),
		},
	}

	if tableSortKey := ps.GetString("TableSortKeyName"); tableSortKey != "" {
		if _, found := attrDefMap[tableSortKey]; !found {
			return nil, fmt.Errorf("invalid: no associated attribute for %s", tableSortKey)
		}

		tableSchema = append(tableSchema, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(tableSortKey),
			KeyType:       aws.String("RANGE"),
		})
	}

	var gsiIndices []*dynamodb.GlobalSecondaryIndex

	for {
		i := len(gsiIndices) + 1

		gsiHashKey := ps.GetString(fmt.Sprintf("GSI%dHashKeyName", i))

		if gsiHashKey == "" {
			break
		}

		if _, found := attrDefMap[gsiHashKey]; !found {
			return nil, fmt.Errorf("invalid: no associated attribute for %s", gsiHashKey)
		}

		gsiSchema := []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(gsiHashKey),
				KeyType:       aws.String("HASH"),
			},
		}

		if sortKey := ps.GetString(fmt.Sprintf("GSI%dSortKeyName", i)); sortKey != "" {
			if _, found := attrDefMap[sortKey]; !found {
				return nil, fmt.Errorf("invalid: no associated attribute for %s", sortKey)
			}

			gsiSchema = append(gsiSchema, &dynamodb.KeySchemaElement{
				AttributeName: aws.String(sortKey),
				KeyType:       aws.String("RANGE"),
			})
		}

		indexName, err := ps.indexName(gsiSchema)

		if err != nil {
			return nil, err
		}

		gsiIndices = append(gsiIndices, &dynamodb.GlobalSecondaryIndex{
			IndexName: aws.String(indexName),
			KeySchema: gsiSchema,

			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},

			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(defaultReadProvision),
				WriteCapacityUnits: aws.Int64(defaultWriteProvision),
			},
		})
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attrDefs,
		KeySchema:            tableSchema,

		GlobalSecondaryIndexes: gsiIndices,

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(defaultReadProvision),
			WriteCapacityUnits: aws.Int64(defaultWriteProvision),
		},

		TableName: aws.String(ps.tableName(app, env, table)),
	}

	return input, nil
}

func DecodeJsonParams(r io.Reader) (*ParameterSet, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	params := make([]ParameterPair, 0)

	err = json.Unmarshal(b, &params)

	if err != nil {
		return nil, err
	}

	return &ParameterSet{
		Parameters: params,
	}, nil
}
