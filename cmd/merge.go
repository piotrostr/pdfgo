package cmd

import (
	"fmt"

	"github.com/piotrostr/pdfeth/pkg/pdf"
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
	// TODO how to make it so that if the package sits in bin for instance
	// it will get the path to the given directory if given arg "." and not
	// read /usr/local/bin
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
