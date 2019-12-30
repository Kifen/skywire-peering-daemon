package cmd

import (
	"fmt"
	"os"

	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
	"github.com/spf13/cobra"
)

var (
	pubKey, namedPipe string
)

var rootCmd = &cobra.Command{
	Use:   "daemon",
	Short: "A skywire-peering-daemon",
	Run: func(cmd *cobra.Command, args []string) {
		daemon := apd.NewApd(pubKey, namedPipe)
		daemon.Run()
	},
}

// Execute executes root CLI command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&pubKey, "pubkey", "", "visor's public key")
	rootCmd.Flags().StringVar(&namedPipe, "nm", "", "path to file `named pipe`")
}
