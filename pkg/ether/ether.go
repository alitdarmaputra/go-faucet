package ether

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Ether struct {
	client *ethclient.Client
}

func NewEther(client *ethclient.Client) *Ether {
	return &Ether{
		client: client,
	}
}

func (e *Ether) GetBalance(ctx context.Context, walletPrivateKey string) (*big.Int, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(walletPrivateKey)
	if err != nil {
		return nil, err
	}

	account, err := GetAccountFromECDSAPrivateKey(privateKeyECDSA)
	if err != nil {
		return nil, err
	}
	balance, err := e.client.BalanceAt(ctx, *account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (e *Ether) GetPendingNonce(ctx context.Context, account common.Address) (uint64, error) {
	nonce, err := e.client.PendingNonceAt(ctx, account)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (e *Ether) Transfer(ctx context.Context, toPublicKey string, walletPrivateKey string, value *big.Int) (string, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(walletPrivateKey)
	if err != nil {
		return "", err
	}

	account, err := GetAccountFromECDSAPrivateKey(privateKeyECDSA)
	if err != nil {
		return "", err
	}

	nonce, err := e.GetPendingNonce(ctx, *account)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000)
	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	toAccount := common.HexToAddress(toPublicKey)
	var data []byte
	tx := types.NewTransaction(nonce, toAccount, value, gasLimit, gasPrice, data)

	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return "", err
	}

	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
