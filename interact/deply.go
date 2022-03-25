package interact

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/blackironj/interact-with-sc-go/api"
	"github.com/blackironj/interact-with-sc-go/auth"
)

var (
	url        string
	privateKey string

	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deploy a contract to network",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ethclient.Dial(url)
			if err != nil {
				er(err)
			}

			txAuth, err := auth.GetAccountAuth(client, privateKey)
			if err != nil {
				er(err)
			}

			deloyedContractAddress, tx, _, err := api.DeployApi(txAuth, client)
			if err != nil {
				er(err)
			}

			_, err = bind.WaitDeployed(context.Background(), client, tx)
			if err != nil {
				er(err)
			}

			fmt.Println("success to deploy contract!")
			fmt.Println("contract address: ", deloyedContractAddress.Hex())
		},
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&url, "url", "u", "", "network url (required)")
	deployCmd.Flags().StringVarP(&privateKey, "privatekey", "p", "", "private-key to deploy a your contract (required)")
}
