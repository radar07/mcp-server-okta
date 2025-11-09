package oktamcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RunStdio runs the server using stdio transport
func RunStdio(ctx context.Context, server *Server) error {
	transport := &mcp.StdioTransport{}
	return server.GetMCPServer().Run(ctx, transport)
}
