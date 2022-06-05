package cmd

import (
	"fmt"
	"log"

	"github.com/piotrostr/pdfgo/pkg/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve [port]",
	Short: "run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port := "8080"
		if len(args) > 0 {
			port = args[0]
		}
		router, err := server.SetupRouter()
		if err != nil {
			return err
		}
		port = fmt.Sprintf(":%s", port)
		log.Printf("starting server on %s\n", port)
		if err := router.Run(port); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
