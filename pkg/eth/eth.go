package eth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ctx = context.Background()

func GetTx(hash string) error {
	conn, err := ethclient.DialContext(ctx, "https://cloudflare-eth.com")
	if err != nil {
		return err
	}

	// wait the transaction not to be pending and practice goroutines/channels
	tx, isPending, err := conn.TransactionByHash(ctx, common.HexToHash(hash))
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
}
