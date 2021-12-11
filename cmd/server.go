package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/fadedpez/driver-tracker/internal/repositories/drivers"
	"github.com/fadedpez/driver-tracker/internal/repositories/teams"

	"github.com/fadedpez/driver-tracker/internal/handlers"
	"github.com/fadedpez/driver-tracker/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server")
		port := 5000
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		handler, err := handlers.NewAlpha(&handlers.AlphaConfig{
			DriverRepo: &drivers.MockRepo{},
			TeamRepo:   &teams.MockRepo{},
		})
		if err != nil {
			log.Fatal("err returned from handlers.NewAlpha()", err)
		}

		protos.RegisterDriverTrackerAPIServer(s, handler)

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(serverCommand)
}
