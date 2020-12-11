package serve

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		App().Run(cmd.Context())
	},
}
