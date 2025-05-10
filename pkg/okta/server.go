package okta

import (
	"fmt"
	"log/slog"

	"github.com/okta/okta-sdk-golang/v5/okta"

	"github.com/radar07/mcp-server-okta/internal/oktamcp"
)

type Server struct {
	log    *slog.Logger
	client *okta.APIClient
	server oktamcp.Server
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

	srv := &Server{
		log:    log,
		client: client,
		server: server,
	}

	return srv, nil
}

func (s *Server) GetMCPServer() oktamcp.Server {
	return s.server
}
