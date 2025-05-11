package okta

import (
	"log/slog"

	okta "github.com/okta/okta-sdk-golang/v5/okta"

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
		AddReadTools(
			FetchUsers(log, client),
		)

	// Add toolsets to the group
	toolsetGroup.AddToolset(users)

	// Enable the requested features
	if err := toolsetGroup.EnableToolsets(enabledToolsets); err != nil {
		return nil, err
	}

	return toolsetGroup, nil
}
