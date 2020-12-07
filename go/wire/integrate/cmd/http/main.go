package http

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		App().Run(cmd.Context())
	},
}
