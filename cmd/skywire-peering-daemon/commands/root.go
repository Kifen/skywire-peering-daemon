package cmd

import (
	"os"
	"log"

	spd "github.com/SkycoinProject/skywire-peering-daemon/pkg/daemon"
	"github.com/spf13/cobra"
)

var (
	pubKey, namedPipe, lAddr string
)

var rootCmd = &cobra.Command{
	Use:   "skywire-peering-daemon",
	Short: "A skywire-peering-skywire-peering-daemon",
	Run: func(cmd *cobra.Command, args []string) {
		d := spd.NewDaemon(pubKey, lAddr, pubKey)
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
	rootCmd.Flags().StringVar(&pubKey, "pubkey", "", "visor's public key")
	rootCmd.Flags().StringVar(&namedPipe, "nm", "", "path to file `named pipe`")
}
