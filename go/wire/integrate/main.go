package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wenzong/demo/cmd/serve"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "This is Demo",
}

func init() {
	rootCmd.AddCommand(serve.Command)
}
