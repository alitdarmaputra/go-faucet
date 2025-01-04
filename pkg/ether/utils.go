package ether

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetAccountFromECDSAPrivateKey(privateKeyECDSA *ecdsa.PrivateKey) (*common.Address, error) {
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	account := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &account, nil
}
