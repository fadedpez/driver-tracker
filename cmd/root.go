package cmd

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "driver-tracker",
}

func newDynamoClient() *dynamo.Client {

	awsRegion := "us-east-1"
	awsEndpoint := "http://localhost:8000"

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	awsClient := dynamodb.NewFromConfig(awsCfg)

	client, err := dynamo.NewClient(&dynamo.ClientConfig{AWSClient: awsClient})
	if err != nil {
		log.Fatal("err returned from dynamo.NewClient. ", err)
	}

	return client
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("error in rootCommand.Execute", err)
	}
}
