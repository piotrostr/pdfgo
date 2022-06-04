package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfeth",
	Short: "merge pdf's for ethereum",
	Long:  `Merge pdf's and pay for the computation in ethereum, its going to be on arbitrum and the cost will be equivalent of calling the lambda endpoint.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
