package cmd

import (
	"context"
	"encoding/json"
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

		// wait the transaction not to be pending and practice goroutines/channels
		tx, isPending, err := conn.TransactionByHash(ctx, common.HexToHash(args[0]))
		if err != nil {
			return err
		}

		num, err := conn.PendingTransactionCount(ctx)
		if err != nil {
			return err
		}

		jsonTx, err := json.MarshalIndent(tx, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", jsonTx)
		fmt.Printf("pending: %v\n", isPending)
		fmt.Printf("total pending in this block: %v\n", num)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
}
