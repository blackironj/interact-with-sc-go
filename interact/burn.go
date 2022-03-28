package interact

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/blackironj/interact-with-sc-go/util"
)

var (
	burnCmd = &cobra.Command{
		Use:   "burn",
		Short: "burn tokens",
		Run: func(cmd *cobra.Command, args []string) {
			if to == "" {
				er("to(address) must be needed")
			}

			conn, txAuth := getContractInteractor(url, contractAddr, privateKey)

			tx, err := conn.Burn(txAuth, common.HexToAddress(to), util.ToWei(amt, 18))
			if err != nil {
				er(err)
			}

			fmt.Println("tx hash hex: ", tx.Hash().Hex())
		},
	}
)

func init() {
	burnCmd.PersistentFlags().StringVarP(&to, "to", "t", "account address to burn tokens", " (required)")
	if err := burnCmd.MarkPersistentFlagRequired("to"); err != nil {
		er(err)
	}

	burnCmd.PersistentFlags().StringVarP(&contractAddr, "contract-addr", "c", "", "token contract address (required)")
	if err := burnCmd.MarkPersistentFlagRequired("contract-addr"); err != nil {
		er(err)
	}

	burnCmd.Flags().StringVarP(&amt, "amt", "a", "1000", "how many token do you want to burn (default: 1000)")

	rootCmd.AddCommand(burnCmd)
}
