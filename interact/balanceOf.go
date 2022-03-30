package interact

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/blackironj/interact-with-sc-go/api"
	"github.com/blackironj/interact-with-sc-go/util"
)

var (
	accountAddr string

	balanceCmd = &cobra.Command{
		Use:   "balance",
		Short: "account token balance",
		Run: func(cmd *cobra.Command, args []string) {
			ethCli, err := ethclient.Dial(url)
			if err != nil {
				er(err)
			}

			conn, err := api.NewApi(common.HexToAddress(contractAddr), ethCli)
			if err != nil {
				er(err)
			}

			balanace, err := conn.BalanceOf(&bind.CallOpts{}, common.HexToAddress(accountAddr))
			if err != nil {
				er(err)
			}

			fmt.Println("MHT balance: ", util.ToDecimal(balanace.String(), 18))
		},
	}
)

func init() {
	balanceCmd.PersistentFlags().StringVarP(&accountAddr, "account-addr", "a", "", "(required)")
	if err := balanceCmd.MarkPersistentFlagRequired("account-addr"); err != nil {
		er(err)
	}

	balanceCmd.PersistentFlags().StringVarP(&contractAddr, "contract-addr", "c", "", "token contract address (required)")
	if err := balanceCmd.MarkPersistentFlagRequired("contract-addr"); err != nil {
		er(err)
	}

	rootCmd.AddCommand(balanceCmd)
}
