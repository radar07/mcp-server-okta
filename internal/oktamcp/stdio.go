package oktamcp

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/mark3labs/mcp-go/server"
)

// NewStdioServer creates a new stdio transport server
func NewStdioServer(mcpServer Server) (*mark3labsStdioImpl, error) {
	sImpl, ok := mcpServer.(*mark3labsImpl)
	if !ok {
		return nil, fmt.Errorf("%w: expected *mark3labsImpl, got %T",
			errors.New("invalid server"), mcpServer)
	}

	return &mark3labsStdioImpl{
		mcpStdioServer: server.NewStdioServer(sImpl.mcpServer),
	}, nil
}

// mark3labsStdioImpl implements the TransportServer
// interface for stdio transport
type mark3labsStdioImpl struct {
	mcpStdioServer *server.StdioServer
}

func (s *mark3labsStdioImpl) Listen(ctx context.Context, in io.Reader, out io.Writer) error {
	return s.mcpStdioServer.Listen(ctx, in, out)
}
