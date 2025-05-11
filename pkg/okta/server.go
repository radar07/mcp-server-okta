package okta

import (
	"fmt"
	"log/slog"

	"github.com/okta/okta-sdk-golang/v5/okta"

	"github.com/radar07/mcp-server-okta/internal/oktamcp"
	"github.com/radar07/mcp-server-okta/pkg/toolset"
)

type Server struct {
	log      *slog.Logger
	client   *okta.APIClient
	server   oktamcp.Server
	toolsets *toolset.ToolsetGroup
}

func NewOktaClient(orgURL, token string) *okta.APIClient {
	cfg, err := okta.NewConfiguration(okta.WithOrgUrl(orgURL), okta.WithToken(token))
	if err != nil {
		fmt.Println("Error")
	}

	return okta.NewAPIClient(cfg)
}

func NewServer(log *slog.Logger, client *okta.APIClient, enabledToolsets []string, readOnly bool) (*Server, error) {
	// Create default options
	opts := []oktamcp.ServerOption{
		oktamcp.WithLogging(),
		oktamcp.WithResourceCapabilities(true, true),
		oktamcp.WithToolCapabilities(true),
	}

	server := oktamcp.NewServer("mcp-server-okta", "0.0.1", opts...)

	toolsets, err := NewToolSets(log, client, enabledToolsets, readOnly)
	if err != nil {
		return nil, fmt.Errorf("failed to create toolsets: %w", err)
	}

	srv := &Server{
		log:      log,
		client:   client,
		server:   server,
		toolsets: toolsets,
	}

	return srv, nil
}

func (s *Server) GetMCPServer() oktamcp.Server {
	return s.server
}
