package cmd

import (
	"context"
	"log"

	"github.com/fadedpez/driver-tracker/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "Run the gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:5000",
			grpc.WithInsecure())
		if err != nil {
			log.Fatal("err returned from grpc.Dial")
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Println("connection not closed")
			}
		}(conn)
		client := protos.NewDriverTrackerAPIClient(conn)

		teamResponse, err := client.StoreTeam(context.Background(), &protos.StoreTeamRequest{
			TeamName:            "beer camp",
			TeamNationality:     "USA",
			TeamPrincipal:       "mongo",
			TeamEstablishedYear: "2015",
			TeamYearsActive:     "6",
		})
		if err != nil {
			log.Println("store team returned error. ", err)
			return
		}

		log.Println("store team was successful. ", teamResponse)

		response, err := client.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameLast:          "kirk",
			NameFirst:         "diggler",
			DriverNumber:      "42",
			DriverNationality: "USA",
		})
		if err != nil {
			log.Println("store driver returned error. ", err)
			return
		}

		log.Println("create driver returned successfully. ", response)

		getResponse, err := client.GetDriver(context.Background(), &protos.GetDriverRequest{Id: response.Driver.Id})
		if err != nil {
			log.Println("get driver returned an error. ", err)
			return
		}

		log.Print("get driver was successful. ", getResponse)
	},
}

func init() {
	rootCommand.AddCommand(testCommand)
}
