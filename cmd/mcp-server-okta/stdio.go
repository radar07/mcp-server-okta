package main

import (
	"context"
	"fmt"
	stdlog "log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	oktasdk "github.com/okta/okta-sdk-golang/v5/okta"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/radar07/mcp-server-okta/internal/oktamcp"
	"github.com/radar07/mcp-server-okta/pkg/log"
	"github.com/radar07/mcp-server-okta/pkg/okta"
)

var stdioCmd = &cobra.Command{
	Use:   "stdio",
	Short: "Standard I/O",
	Long:  "A command that reads from standard input and writes to standard output.",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiToken := viper.GetString("api_token")
		if apiToken == "" {
			return fmt.Errorf("`api_token` is required")
		}

		orgURL := viper.GetString("org_url")
		if orgURL == "" {
			return fmt.Errorf("`org_url` is required")
		}

		logPath := viper.GetString("log_file")
		log, close, err := log.New(logPath)
		if err != nil {
			stdlog.Fatalf("create logger: %v", err)
		}
		defer close()

		// Get toolsets to enable from config
		enabledToolsets := viper.GetStringSlice("toolsets")

		client := okta.NewOktaClient(orgURL, apiToken)

		readOnly := viper.GetBool("read_only")
		return runStdioServer(log, client, enabledToolsets, readOnly)
	},
}

func runStdioServer(
	log *slog.Logger,
	client *oktasdk.APIClient,
	enabledToolsets []string,
	readOnly bool,
) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srvr, err := okta.NewServer(log, client, enabledToolsets, readOnly)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	_, _ = fmt.Fprintf(
		os.Stderr,
		"Okta MCP Server running on stdio\n",
	)

	log.Info("starting server")

	// Run the server using the stdio transport
	return oktamcp.RunStdio(ctx, srvr.GetMCPServer())
}
