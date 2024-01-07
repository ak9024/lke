package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cmd = cobra.Command{}

func init() {
	cmd.AddCommand(Server())
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
