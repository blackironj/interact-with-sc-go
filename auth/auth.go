package auth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountAuth(client *ethclient.Client, accountAddress string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		return nil, err
	}

	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3_000_000)
	auth.GasPrice = big.NewInt(1_000_000)

	return auth, nil
}
