package cmd

import (
	"log"

	"github.com/KirkDiggler/go-projects/dynamo"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "driver-tracker",
}

func newDynamoClient() *dynamo.Client {
	endpoint := "http://localhost:8000"
	region := "us-east-1"

	awsClient := dynamodb.New(dynamodb.Options{
		Region:           region,
		EndpointResolver: dynamodb.EndpointResolverFromURL(endpoint),
	})

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
