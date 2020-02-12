package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	spd "github.com/SkycoinProject/skywire-peering-daemon/pkg/daemon"
)

var (
	pubKey, namedPipe, lAddr string
)

var rootCmd = &cobra.Command{
	Use:   "skywire-peering-daemon",
	Short: "A skywire-peering-skywire-peering-daemon",
	Run: func(cmd *cobra.Command, args []string) {
		d := spd.NewDaemon(pubKey, lAddr, namedPipe)
		d.Run()
	},
}

// Execute executes root CLI command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&pubKey, "pubkey", "", "none", "visor's public key")
	rootCmd.Flags().StringVarP(&namedPipe, "named-pipe", "", "none", "path to file `named pipe`")
	rootCmd.Flags().StringVarP(&lAddr, "laddr", "", "none", "address of visor")
}
