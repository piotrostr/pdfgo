package cmd

import (
	"errors"

	"github.com/piotrostr/pdfeth/pkg/eth"
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx [hash]",
	Short: "gets the tx by hash",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no hash given")
		}
		eth.GetTx(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
