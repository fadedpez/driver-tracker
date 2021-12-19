package decode

import (
	"strings"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDecodeJsonParms(t *testing.T) {
	t.Run("errors on empty string", func(t *testing.T) {
		_, err := DecodeJsonParams(strings.NewReader(""))

		assert.Error(t, err)
	})

	t.Run("errors on invalid json", func(t *testing.T) {
		_, err := DecodeJsonParams(strings.NewReader("<<<>>>>"))

		assert.Error(t, err)
	})

	t.Run("errors on invalid format for parameter set", func(t *testing.T) {
		_, err := DecodeJsonParams(strings.NewReader(`{"a":1}`))

		assert.Error(t, err)
	})

	t.Run("return empty parameter set", func(t *testing.T) {
		ps, err := DecodeJsonParams(strings.NewReader(`[]`))

		exp := ParameterSet{
			Parameters: make([]ParameterPair, 0),
		}

		assert.Nil(t, err)
		assert.Equal(t, exp, *ps)
	})

	t.Run("returns a parameter set", func(t *testing.T) {
		str := `[{
    "ParameterKey": "AppName",
    "ParameterValue": "dynctl"
  },
  {
    "ParameterKey": "AppEnvironment",
    "ParameterValue": "dev"
  }]`

		ps, err := DecodeJsonParams(strings.NewReader(str))

		exp := ParameterSet{
			Parameters: []ParameterPair{
				{
					Key:   "AppName",
					Value: "dynctl",
				},
				{
					Key:   "AppEnvironment",
					Value: "dev",
				},
			},
		}

		assert.Nil(t, err)
		assert.Equal(t, exp, *ps)
	})
}

func TestParameterSet(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		t.Run("returns empty string with key not found", func(t *testing.T) {
			ps := ParameterSet{}

			assert.Equal(t, "", ps.GetString("foo"))
		})

		t.Run("returns value when key is found", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			}

			assert.Equal(t, "bar", ps.GetString("foo"))
		})
	})

	t.Run("CreateTableInput", func(t *testing.T) {
		ppAppName := ParameterPair{
			Key:   "AppName",
			Value: "App-1",
		}

		ppAppEnv := ParameterPair{
			Key:   "AppEnvironment",
			Value: "dev",
		}

		ppTableName := ParameterPair{
			Key:   "TableName",
			Value: "users",
		}

		ppAttr1Name := ParameterPair{
			Key:   "Attribute1Name",
			Value: "id",
		}

		ppAttr1Type := ParameterPair{
			Key:   "Attribute1Type",
			Value: "N",
		}

		ppAttr2Name := ParameterPair{
			Key:   "Attribute2Name",
			Value: "updated_at",
		}

		ppAttr2Type := ParameterPair{
			Key:   "Attribute2Type",
			Value: "S",
		}

		ppAttr3Name := ParameterPair{
			Key:   "Attribute3Name",
			Value: "foo",
		}

		ppAttr3Type := ParameterPair{
			Key:   "Attribute3Type",
			Value: "S",
		}

		ppTableHashKey := ParameterPair{
			Key:   "TableHashKeyName",
			Value: "id",
		}

		ppTableSortKey := ParameterPair{
			Key:   "TableSortKeyName",
			Value: "updated_at",
		}

		ppGsi1HashKey := ParameterPair{
			Key:   "GSI1HashKeyName",
			Value: "foo",
		}

		ppGsi1SortKey := ParameterPair{
			Key:   "GSI1SortKeyName",
			Value: "updated_at",
		}

		t.Run("throws error with empty ps (no AppName)", func(t *testing.T) {
			ps := ParameterSet{}

			_, err := ps.CreateTableInput("", "staging", "brands")

			assert.Error(t, err, "invalid: AppName")
		})

		t.Run("throws error with no AppEnvironment", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "", "brands")

			assert.Error(t, err, "invalid: AppEnvironment")
		})

		t.Run("throws error with no TableName", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "staging", "")

			assert.Error(t, err, "invalid: TableName")
		})

		t.Run("throws error with no TableHashKeyName", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Error(t, err, "invalid: TableHashKeyName")
		})

		t.Run("throws an error with no associated Attribute for TableHashKey", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppTableHashKey,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Error(t, err, fmt.Sprintf("invalid: no associated attribute for %s", ppTableHashKey))
		})

		t.Run("throws an error with no TableHashKeyType", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppTableHashKey,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Error(t, err, "invalid: TableHashKeyName")
		})

		t.Run("throws an error with no TableSortKeyType", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppAttr1Name,
					ppAttr1Type,
					ppTableHashKey,
					ppTableSortKey,
				},
			}

			_, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Error(t, err, "type is required for sort key updated_at")
		})

		t.Run("returns a CreateTableOutput (no indices)", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppAttr1Name,
					ppAttr1Type,
					ppAttr2Name,
					ppAttr2Type,
					ppTableHashKey,
					ppTableSortKey,
				},
			}

			result, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Nil(t, err)
			assert.Equal(t, 2, len(result.AttributeDefinitions))
			assert.Equal(t, 2, len(result.KeySchema))
			assert.Nil(t, result.GlobalSecondaryIndexes)
		})

		t.Run("returns an error when gsi is missing associated hash key", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppAttr1Name,
					ppAttr1Type,
					ppAttr2Name,
					ppAttr2Type,
					ppTableHashKey,
					ppTableSortKey,
					ppGsi1HashKey,
					ppGsi1SortKey,
				},
			}

			result, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Error(t, err, fmt.Sprintf("invalid: no associated attribute for %s", ppAttr3Name))
			assert.Nil(t, result)
		})

		t.Run("returns a gsi with RANGE key", func(t *testing.T) {
			ps := ParameterSet{
				Parameters: []ParameterPair{
					ppAppName,
					ppAppEnv,
					ppTableName,
					ppAttr1Name,
					ppAttr1Type,
					ppAttr2Name,
					ppAttr2Type,
					ppAttr3Name,
					ppAttr3Type,
					ppTableHashKey,
					ppTableSortKey,
					ppGsi1HashKey,
					ppGsi1SortKey,
				},
			}

			result, err := ps.CreateTableInput("sneakers", "staging", "brands")

			assert.Nil(t, err)
			assert.Equal(t, 3, len(result.AttributeDefinitions))
			assert.Equal(t, 2, len(result.KeySchema))
			assert.Equal(t, 1, len(result.GlobalSecondaryIndexes))
			assert.Equal(t, "foo-updated_at-index", *result.GlobalSecondaryIndexes[0].IndexName)
			assert.NotNil(t, result.GlobalSecondaryIndexes[0].Projection)
		})
	})
}
