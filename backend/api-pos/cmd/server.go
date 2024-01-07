package cmd

import (
	"api-pos/internal/app/server"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	var cmd = cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			if err := server.StartApp(); err != nil {
				slog.Error("Error to running the server", "error", err)
				os.Exit(1)
			}
		},
	}
	return &cmd
}
