package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	spd "github.com/SkycoinProject/skywire-peering-daemon/pkg/daemon"
)

var cfg *spd.Config

var rootCmd = &cobra.Command{
	Use:   "skywire-peering-daemon",
	Short: "A skywire-peering-skywire-peering-daemon",
	Run: func(cmd *cobra.Command, args []string) {
		config := readEnv()
		d := spd.NewDaemon(config)
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

func readEnv() *spd.Config {
	cfg = &spd.Config{}
	viper.SetEnvPrefix("spd")
	err := bindEnv("pubkey", "laddr", "named-pipe")
	if err != nil {
		log.Fatal(err)
	}

	cfg.PubKey = viper.GetString("pubkey")
	cfg.NamedPipe = viper.GetString("named-pipe")
	cfg.LocalAddr = viper.GetString("laddr")

	return cfg
}

func bindEnv(args ...string) error {
	for _, arg := range args {
		if err := viper.BindEnv(arg); err != nil {
			return fmt.Errorf("Could not bind env: %s", err)
		}
	}

	return nil
}
