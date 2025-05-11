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

	// flags will be available for all subcommands
	rootCmd.PersistentFlags().String("api_token", "", "your okta api key")
	rootCmd.PersistentFlags().String("org_url", "", "your okta org url")
	rootCmd.PersistentFlags().String("log_file", "", "path to the log file")
	rootCmd.PersistentFlags().StringSlice("toolsets", []string{}, "comma-separated list of toolsets to enable")
	rootCmd.PersistentFlags().Bool("read_only", false, "run server in read-only mode")

	// bind flags to viper
	_ = viper.BindPFlag("api_token", rootCmd.PersistentFlags().Lookup("api_token"))
	_ = viper.BindPFlag("org_url", rootCmd.PersistentFlags().Lookup("org_url"))
	_ = viper.BindPFlag("log_file", rootCmd.PersistentFlags().Lookup("log_file"))
	_ = viper.BindPFlag("toolsets", rootCmd.PersistentFlags().Lookup("toolsets"))
	_ = viper.BindPFlag("read_only", rootCmd.PersistentFlags().Lookup("read-only"))

	// Set environment variable mappings
	_ = viper.BindEnv("key", "OKTA_KEY_ID")        // Maps RAZORPAY_KEY_ID to key
	_ = viper.BindEnv("secret", "OKTA_KEY_SECRET") // Maps RAZORPAY_KEY_SECRET to secret

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
