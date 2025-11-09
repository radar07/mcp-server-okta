package okta

import (
	"log/slog"

	okta "github.com/okta/okta-sdk-golang/v5/okta"

	"github.com/radar07/mcp-server-okta/internal/oktamcp"
	"github.com/radar07/mcp-server-okta/pkg/toolset"
)

func NewToolSets(
	log *slog.Logger,
	client *okta.APIClient,
	enabledToolsets []string,
	readOnly bool,
) (*toolset.ToolsetGroup, error) {
	// Create a new toolset group
	toolsetGroup := toolset.NewToolsetGroup(readOnly)

	// Create toolsets
	users := toolset.NewToolset("users", "Okta Users").
		AddWriteTools(
			toolset.ToolDefinition{
				Name:        "create_user",
				Description: "Use this tool to create a user in the org",
				Register: func(s *oktamcp.Server) {
					oktamcp.AddTool(s, &oktamcp.Tool{
						Name:        "create_user",
						Description: "Use this tool to create a user in the org",
					}, CreateUser(log, client))
				},
			},
		).
		AddReadTools(
			toolset.ToolDefinition{
				Name:        "fetch_users",
				Description: "Use this tool to retrieve the users of an org",
				Register: func(s *oktamcp.Server) {
					oktamcp.AddTool(s, &oktamcp.Tool{
						Name:        "fetch_users",
						Description: "Use this tool to retrieve the users of an org",
					}, FetchUsers(log, client))
				},
			},
		)

	// Add toolsets to the group
	toolsetGroup.AddToolset(users)

	// Enable the requested features
	if err := toolsetGroup.EnableToolsets(enabledToolsets); err != nil {
		return nil, err
	}

	return toolsetGroup, nil
}
