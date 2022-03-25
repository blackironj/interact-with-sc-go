package interact

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sc-interactor",
		Short: "A helper for managing TOKEN contract",
	}
)

func init() {
}

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}
