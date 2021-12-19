package decode

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	yaml "gopkg.in/yaml.v2"
)

const (
	cabraDynTableResourceType = "dynamo"
)

type CabraParamPair struct {
	Name           string `yaml:"name"`
	Value          string `yaml:"value"`
	ParameterStore bool   `yaml:"parameter_store"`
}

type CabraResource struct {
	Type   string            `yaml:"type"`
	Params map[string]string `yaml:"parameters"`
}

type CabraConfig struct {
	Resources map[string]CabraResource `yaml:"resources"`
}

func (cfg *CabraConfig) CreateTableInputs(appName, environment string) ([]*dynamodb.CreateTableInput, error) {
	results := make([]*dynamodb.CreateTableInput, 0)

	for tableName, res := range cfg.Resources {
		if res.Type != cabraDynTableResourceType {
			continue
		}

		paramPairs := make([]ParameterPair, 0)

		for name, value := range res.Params {
			paramPairs = append(paramPairs, ParameterPair{
				Key:   name,
				Value: value,
			})
		}

		ps := ParameterSet{
			Parameters: paramPairs,
		}

		input, err := ps.CreateTableInput(appName, environment, tableName)

		if err != nil {
			return nil, err
		}

		results = append(results, input)
	}

	return results, nil
}

func DecodeCabraConfig(r io.Reader) (*CabraConfig, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	cfg := &CabraConfig{}

	err = yaml.Unmarshal(b, cfg)

	return cfg, err
}

func DecodeCabraFile(fp string) (*CabraConfig, error) {
	if fp == "" {
		return nil, fmt.Errorf("invalid file path")
	}

	yamlFile, err := os.Open(fp)

	if err != nil {
		return nil, err
	}

	return DecodeCabraConfig(yamlFile)
}
