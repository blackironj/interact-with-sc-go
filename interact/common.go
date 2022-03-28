package interact

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/blackironj/interact-with-sc-go/api"
	"github.com/blackironj/interact-with-sc-go/auth"
)

var (
	amt          string
	to           string
	contractAddr string
)

func getContractInteractor(url, contractAddr, privateKey string) (*api.Api, *bind.TransactOpts) {
	ethCli, err := ethclient.Dial(url)
	if err != nil {
		er(err)
	}

	conn, err := api.NewApi(common.HexToAddress(contractAddr), ethCli)
	if err != nil {
		er(err)
	}

	txAuth, err := auth.GetAccountAuth(ethCli, privateKey)
	if err != nil {
		er(err)
	}

	return conn, txAuth
}
