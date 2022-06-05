package cmd

import (
	"fmt"

	"github.com/piotrostr/pdfgo/pkg/pdf"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge [directory]",
	Short: "merges the pdfs in the given directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println("please provide a directory")
			return cmd.Usage()
		}
		err := pdf.Merge(args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
