package oktamcp

import (
	"github.com/mark3labs/mcp-go/server"
)

type Server interface{}

func NewServer(name, version string, opts ...ServerOption) *mark3labsImpl {
	// Create option setter to collect mcp options
	optSetter := &mark3labsOptionSetter{
		mcpOptions: []server.ServerOption{},
	}

	mcpServer := server.NewMCPServer(name, version, optSetter.mcpOptions...)

	return &mark3labsImpl{
		mcpServer: mcpServer,
		name:      name,
	}
}

// mark3labsImpl implements the Server interface using mark3labs/mcp-go
type mark3labsImpl struct {
	mcpServer *server.MCPServer
	name      string
}

func (s *mark3labsOptionSetter) SetOption(option any) error {
	if opt, ok := option.(server.ServerOption); ok {
		s.mcpOptions = append(s.mcpOptions, opt)
	}
	return nil
}

// mark3labsOptionSetter is used to apply options to the server
type mark3labsOptionSetter struct {
	mcpOptions []server.ServerOption
}

// OptionSetter is an interface for setting options on a configurable object
type OptionSetter interface {
	SetOption(option any) error
}

type ServerOption func(OptionSetter) error

// WithLogging returns a server option that enables logging
func WithLogging() ServerOption {
	return func(s OptionSetter) error {
		return s.SetOption(server.WithLogging())
	}
}

// WithResourceCapabilities returns a server option
// that enables resource capabilities
func WithResourceCapabilities(read, list bool) ServerOption {
	return func(s OptionSetter) error {
		return s.SetOption(server.WithResourceCapabilities(read, list))
	}
}

// WithToolCapabilities returns a server option that enables tool capabilities
func WithToolCapabilities(enabled bool) ServerOption {
	return func(s OptionSetter) error {
		return s.SetOption(server.WithToolCapabilities(enabled))
	}
}
