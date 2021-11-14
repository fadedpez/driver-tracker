package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "driver-tracker",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("error in rootCommand.Execute")
	}
}
