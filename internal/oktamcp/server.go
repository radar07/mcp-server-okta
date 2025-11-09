package oktamcp

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server wraps the official MCP SDK server
type Server struct {
	mcp *mcp.Server
}

// NewServer creates a new MCP server with the given name and version
func NewServer(name, version string) *Server {
	impl := &mcp.Implementation{
		Name:    name,
		Version: version,
	}

	mcpServer := mcp.NewServer(impl, nil)

	return &Server{
		mcp: mcpServer,
	}
}

// GetMCPServer returns the underlying mcp.Server for direct access
// This is needed for operations like Run() and adding tools
func (s *Server) GetMCPServer() *mcp.Server {
	return s.mcp
}

// AddTools is a placeholder method for compatibility
// Tools should now be added directly using oktamcp.AddTool or mcp.AddTool
func (s *Server) AddTools(tools ...any) {
	// This method exists for compatibility with the toolset package
	// but tools are now added using the generic AddTool function
}
