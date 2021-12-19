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

	Run: func(comd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5000")
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

		response, err := client.StoreDriver(context.Background(), &protos.StoreDriverRequest{
			NameLast:          "kirk",
			NameFirst:         "diggler",
			DriverNumber:      "42",
			DriverNationality: "USA",
		})
		if err != nil {
			log.Println("error returned", err)
			return
		}

		log.Println("return successfully", response)
	},
}

func init() {
	rootCommand.AddCommand(testCommand)
}
