package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var ctx = context.Background()

var txCmd = &cobra.Command{
	Use:   "tx [hash]",
	Short: "gets the tx by hash",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no hash given")
		}
		conn, err := ethclient.DialContext(ctx, "https://cloudflare-eth.com")
		if err != nil {
			return err
		}
		tx, isPending, err := conn.TransactionByHash(ctx, common.HexToHash(args[0]))
		if err != nil {
			return err
		}
		fmt.Println(tx)
		fmt.Println(isPending)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
