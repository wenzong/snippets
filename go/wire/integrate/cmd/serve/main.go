package serve

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, _ := App()
		defer cleanup()

		app.Run(cmd.Context())
	},
}
