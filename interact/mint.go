package interact

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/blackironj/interact-with-sc-go/api"
	"github.com/blackironj/interact-with-sc-go/auth"
	"github.com/blackironj/interact-with-sc-go/util"
)

var (
	amt          string
	to           string
	contractAddr string

	mintCmd = &cobra.Command{
		Use:   "mint",
		Short: "mint tokens",
		Run: func(cmd *cobra.Command, args []string) {
			if to == "" {
				er("to(address) must be needed")
			}

			client, err := ethclient.Dial(url)
			if err != nil {
				er(err)
			}

			conn, err := api.NewApi(common.HexToAddress(contractAddr), client)
			if err != nil {
				er(err)
			}

			txAuth, err := auth.GetAccountAuth(client, privateKey)
			if err != nil {
				er(err)
			}

			tx, err := conn.Mint(txAuth, common.HexToAddress(to), util.ToWei(amt, 18))
			if err != nil {
				er(err)
			}

			fmt.Println("tx hash hex: ", tx.Hash().Hex())
		},
	}
)

func init() {
	mintCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "account address to receive token after minting (required)")
	if err := mintCmd.MarkPersistentFlagRequired("to"); err != nil {
		er(err)
	}

	mintCmd.PersistentFlags().StringVarP(&contractAddr, "contract-addr", "c", "", "token contract address (required)")
	if err := mintCmd.MarkPersistentFlagRequired("contract-addr"); err != nil {
		er(err)
	}

	mintCmd.Flags().StringVarP(&amt, "amt", "a", "1000", "how many token do you want to mint (default: 1000)")

	rootCmd.AddCommand(mintCmd)
}
