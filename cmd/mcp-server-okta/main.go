package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Okta MCP Server",
	Long:  "An Okta MCP Server that provides a command line interface to manage Okta resources.",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(stdioCmd)
}

func initConfig() {
	viper.SetEnvPrefix("okta")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
