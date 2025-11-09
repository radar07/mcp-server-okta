package oktamcp

// This file re-exports types from the official MCP SDK for convenience
// and to minimize changes to the rest of the codebase.

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Re-export commonly used types from the official SDK
type (
	CallToolRequest = mcp.CallToolRequest
	CallToolResult  = mcp.CallToolResult
	Tool            = mcp.Tool
	TextContent     = mcp.TextContent
	Content         = mcp.Content
)

// AddTool is a convenience wrapper around mcp.AddTool
// It adds a tool with typed input and output to the server
func AddTool[In, Out any](
	server *Server,
	tool *Tool,
	handler func(ctx context.Context, req *CallToolRequest, input In) (*CallToolResult, Out, error),
) {
	mcp.AddTool(server.GetMCPServer(), tool, handler)
}
