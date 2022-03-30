package interact

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/blackironj/interact-with-sc-go/util"
)

var (
	transferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "transfer tokens to others",
		Run: func(cmd *cobra.Command, args []string) {
			conn, txAuth := getContractInteractor(url, contractAddr, privateKey)

			tx, err := conn.Transfer(txAuth, common.HexToAddress(to), util.ToWei(amt, 18))
			if err != nil {
				er(err)
			}

			fmt.Println("tx hash hex: ", tx.Hash().Hex())
		},
	}
)

func init() {
	transferCmd.PersistentFlags().StringVarP(&privateKey, "privatekey", "p", "", "private-key to deploy a your contract (required)")
	if err := transferCmd.MarkPersistentFlagRequired("privatekey"); err != nil {
		er(err)
	}

	transferCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "account address to receive token after minting (required)")
	if err := transferCmd.MarkPersistentFlagRequired("to"); err != nil {
		er(err)
	}

	transferCmd.PersistentFlags().StringVarP(&contractAddr, "contract-addr", "c", "", "token contract address (required)")
	if err := transferCmd.MarkPersistentFlagRequired("contract-addr"); err != nil {
		er(err)
	}

	transferCmd.Flags().StringVarP(&amt, "amt", "a", "1000", "how many token do you want to mint (default: 1000)")

	rootCmd.AddCommand(transferCmd)
}
