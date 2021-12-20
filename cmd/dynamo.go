package cmd

import (
	"fmt"

	"github.com/fadedpez/driver-tracker/dyctl/decode"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cabraCmd = &cobra.Command{
	Use:   "cabra <yaml-file>",
	Short: "Creates tables from a Goatbot-Cabra config yaml file",
	Example: `
AWS_ENDPOINT=http://localhost:8000 AWS_REGION=us-east-1 DELETE_EXISTING=true dynctl cabra ../logged-actions/cabra-config/staging.yaml
`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := decode.DecodeCabraFile(args[0])

		if err != nil {
			panic(err)
		}

		fmt.Println(cfg)

		appName := viper.GetString("cabra.app-name")
		environment := viper.GetString("cabra.environment")

		inputs, err := cfg.CreateTableInputs(appName, environment)

		if err != nil {
			panic(err)
		}

		fmt.Println(inputs)

		endpoint := viper.GetString("cabra.endpoint")
		region := viper.GetString("cabra.region")
		deleteExisting := viper.GetBool("cabra.delete-existing")

		fmt.Println("delete existing", deleteExisting)

		awsConf := aws.NewConfig()

		if endpoint != "" {
			awsConf = awsConf.WithEndpoint(endpoint)
		}

		if region != "" {
			awsConf = awsConf.WithRegion(region)
		}

		client := dynamodb.New(
			session.Must(session.NewSession()),
			awsConf,
		)

		for _, input := range inputs {
			if deleteExisting {
				_, err := client.DescribeTable(&dynamodb.DescribeTableInput{
					TableName: input.TableName,
				})

				if err == nil {
					_, err := client.DeleteTable(&dynamodb.DeleteTableInput{
						TableName: input.TableName,
					})

					if err != nil {
						panic(err)
					}
				}
			}

			resp, err := client.CreateTable(input)

			if err != nil {
				panic(err)
			}

			fmt.Println(resp.TableDescription)
		}
	},
}

func init() {
	rootCommand.AddCommand(cabraCmd)

	cabraCmd.Flags().StringP("endpoint", "e", "", "Dynamo endpoint")
	_ = viper.BindPFlag("cabra.endpoint", cabraCmd.Flags().Lookup("endpoint"))
	_ = viper.BindEnv("cabra.endpoint", "AWS_ENDPOINT")

	cabraCmd.Flags().StringP("region", "r", "", "Dynamo region")
	_ = viper.BindPFlag("cabra.region", cabraCmd.Flags().Lookup("region"))
	_ = viper.BindEnv("cabra.region", "AWS_REGION")

	cabraCmd.Flags().StringP("app-name", "a", "", "App name")
	_ = viper.BindPFlag("cabra.app-name", cabraCmd.Flags().Lookup("app-name"))

	cabraCmd.Flags().StringP("environment", "v", "local", "Environment")
	_ = viper.BindPFlag("cabra.environment", cabraCmd.Flags().Lookup("environment"))

	cabraCmd.Flags().BoolP("delete-existing", "d", true, "Delete existing table if exists")
	_ = viper.BindPFlag("cabra.delete-existing", cabraCmd.Flags().Lookup("delete-existing"))
	_ = viper.BindEnv("cabra.delete-existing", "DELETE_EXISTING")
}
