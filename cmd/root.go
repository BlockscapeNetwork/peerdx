package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/BlockscapeLab/peerdx/config"
)

var (
	cfgFile string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "optional config file")

	rootCmd.AddCommand(createAddrbookCmd())
	rootCmd.AddCommand(createRPCCmd())
}

var rootCmd = &cobra.Command{
	Use:   "peer-diagnostics",
	Short: "diagnose peers of your nodes",
	Long:  `diagnose peers of your nodes either by analyzing address books or rpc connections to your nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts command line app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfig() {
	if cfgFile == "" {
		return
	}
	err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Println("Couldn't load config file from", cfgFile, "so will use default instead:", err)
	}
}
