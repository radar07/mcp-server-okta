package okta

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/okta/okta-sdk-golang/v5/okta"
	"github.com/radar07/mcp-server-okta/internal/oktamcp"
)

// Create an Okta user
func CreateUser(log *slog.Logger, client *okta.APIClient) oktamcp.Tool {
	parameters := []oktamcp.ToolParameter{
		oktamcp.WithString(
			"first_name",
			oktamcp.Description("First name of the user"),
			oktamcp.Required(),
		),
		oktamcp.WithString(
			"last_name",
			oktamcp.Description("Last name of the user"),
			oktamcp.Required(),
		),
		oktamcp.WithString(
			"email",
			oktamcp.Description("Email of the user"),
			oktamcp.Required(),
		),
		oktamcp.WithString(
			"login",
			oktamcp.Description("Login of the user"),
			oktamcp.Required(),
		),
	}

	handler := func(ctx context.Context, r oktamcp.CallToolRequest) (*oktamcp.ToolResult, error) {
		params := make(map[string]any)

		validator := NewValidator(&r).
			ValidateAndAddRequiredString(params, "first_name").
			ValidateAndAddRequiredString(params, "last_name").
			ValidateAndAddRequiredString(params, "email").
			ValidateAndAddRequiredString(params, "login")

		if result, err := validator.HandleErrorsIfAny(); result != nil {
			return result, err
		}

		firstName := params["first_name"].(string)
		lastName := params["last_name"].(string)
		email := params["email"].(string)
		login := params["login"].(string)

		profile := okta.NewUserProfile()
		profile.FirstName = *okta.NewNullableString(&firstName)
		profile.LastName = *okta.NewNullableString(&lastName)
		profile.Email = &email
		profile.Login = &login
		okta.NewCreateUserRequest(*profile)

		createUserRequest := okta.CreateUserRequest{
			Profile: *profile,
		}

		user, _, err := client.UserAPI.CreateUser(ctx).Body(createUserRequest).Activate(true).Execute()
		if err != nil {
			fmt.Printf("Error Creating Users: %v\n", err)
		}

		if err != nil {
			return oktamcp.NewToolResultError(
				fmt.Sprintf("updating user failed: %s", err.Error())), nil
		}

		return oktamcp.NewToolResultJSON(user)
	}

	return oktamcp.NewTool(
		"create_user",
		"Use this tool to createupdate an user in the org",
		parameters,
		handler,
	)
}

// Fetch users from the org
func FetchUsers(log *slog.Logger, client *okta.APIClient) oktamcp.Tool {
	parameters := []oktamcp.ToolParameter{}

	handler := func(ctx context.Context, r oktamcp.CallToolRequest) (*oktamcp.ToolResult, error) {
		users, _, err := client.UserAPI.ListUsers(ctx).Execute()
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
