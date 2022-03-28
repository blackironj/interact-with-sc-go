package interact

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	url        string
	privateKey string

	rootCmd = &cobra.Command{
		Use:   "sc-interactor",
		Short: "A helper for managing TOKEN contract",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&privateKey, "privatekey", "p", "", "private-key to deploy a your contract or generate a transaction (required)")
	if err := rootCmd.MarkPersistentFlagRequired("privatekey"); err != nil {
		er(err)
	}

	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "network url (required)")
	if err := rootCmd.MarkPersistentFlagRequired("url"); err != nil {
		er(err)
	}
}

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}

	if privateKey == "" || url == "" {
		er(errors.New("private key & url must be needed"))
	}
}
