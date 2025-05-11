package okta

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/okta/okta-sdk-golang/v5/okta"
	"github.com/radar07/mcp-server-okta/internal/oktamcp"
)

func FetchUsers(log *slog.Logger, client *okta.APIClient) oktamcp.Tool {
	parameters := []oktamcp.ToolParameter{}

	handler := func(ctx context.Context, r oktamcp.CallToolRequest) (*oktamcp.ToolResult, error) {
		users, _, err := client.UserAPI.ListUsers(client.GetConfig().Context).Execute()
		if err != nil {
			return oktamcp.NewToolResultError(
				fmt.Sprintf("fetching users failed: %s", err.Error())), nil
		}

		return oktamcp.NewToolResultJSON(users)
	}

	return oktamcp.NewTool(
		"fetch_users",
		"Use this tool to retrieve the users of an org",
		parameters,
		handler,
	)
}
