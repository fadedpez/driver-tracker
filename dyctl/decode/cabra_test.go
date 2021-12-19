package decode

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCabraConfig(t *testing.T) {
	t.Run("errors on empty string", func(t *testing.T) {
		_, err := DecodeJsonParams(strings.NewReader(""))

		assert.Error(t, err)
	})

	t.Run("errors on invalid yaml", func(t *testing.T) {
		_, err := DecodeJsonParams(strings.NewReader("<<<>>>>"))

		assert.Error(t, err)
	})

	t.Run("decodes valid yaml with no dynamo table resources", func(t *testing.T) {
		cfg, err := DecodeCabraConfig(strings.NewReader(`
schema: '20180101'

environment:
  - name: FOO
    value: thing

  - name: BAR
    parameter_store: true

  - name: BAZ
    value: 1111
`))

		assert.Nil(t, err)

		expectedTables := make([]*dynamodb.CreateTableInput, 0)

		result, err := cfg.CreateTableInputs("sneakers-stream", "staging")

		assert.Nil(t, err)
		assert.Equal(t, expectedTables, result)
	})

	t.Run("returns a valid CreateTableInput", func(t *testing.T) {
		cfg, err := DecodeCabraConfig(strings.NewReader(`
resources:
  brands:
    type: dynamo
    parameters:
      Attribute1Name: id
      Attribute1Type: 'N'
      TableHashKeyName: id
      MaxReadCapacity: 10
      MinReadCapacity: 1
      MaxWriteCapacity: 10
      MinWriteCapacity: 1
`))

		assert.Nil(t, err)

		brandsTable := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("N"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			TableName: aws.String("sneakers-stream-staging-brands"),
		}

		expectedTables := []*dynamodb.CreateTableInput{
			brandsTable,
		}

		result, err := cfg.CreateTableInputs("sneakers-stream", "staging")

		assert.Nil(t, err)
		assert.EqualValues(t, expectedTables, result)
	})
}
