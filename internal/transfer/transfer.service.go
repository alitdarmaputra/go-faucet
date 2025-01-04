package transfer

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/alitdarmaputra/go-faucet/config"
	"github.com/alitdarmaputra/go-faucet/pkg/ether"
)

type TransferService struct {
	config *config.Config
	ether  *ether.Ether
}

func NewTransferService(config *config.Config, ether *ether.Ether) *TransferService {
	return &TransferService{
		config: config,
		ether:  ether,
	}
}

func (t *TransferService) TransferService(ctx context.Context, publicKey string) string {
	// validate balance
	balance, err := t.ether.GetBalance(ctx, t.config.WalletPrivateKey)
	if err != nil {
		log.Println("error: Getting faucet account balance\n", err)
		return "Error sending token"
	}

	if balance.Cmp(big.NewInt(0)) == -1 { // balance < 0
		return "Faucet account out of balance, Please contact the customer support"
	}

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)

	transactionHash, err := t.ether.Transfer(ctx, publicKey, t.config.WalletPrivateKey, value)
	if err != nil {
		log.Println("error: Sending token", err)
		return "Error sending token"
	}

	return fmt.Sprintf("Success sending token to `%s` with transaction hash: `%s`", publicKey, transactionHash)
}
