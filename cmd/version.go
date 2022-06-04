package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the version",
	Long:  "prints the version of the package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
