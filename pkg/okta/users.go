package okta

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/okta/okta-sdk-golang/v5/okta"
	"github.com/radar07/mcp-server-okta/internal/oktamcp"
)

// CreateUserInput defines the input structure for creating a user
type CreateUserInput struct {
	FirstName string `json:"first_name" jsonschema:"required,description=First name of the user"`
	LastName  string `json:"last_name"  jsonschema:"required,description=Last name of the user"`
	Email     string `json:"email"      jsonschema:"required,description=Email of the user"`
	Login     string `json:"login"      jsonschema:"required,description=Login of the user"`
}

// CreateUserOutput defines the output structure for creating a user
type CreateUserOutput struct {
	User *okta.User `json:"user,omitempty" jsonschema:"description=The created user object"`
}

// CreateUser creates an Okta user
func CreateUser(
	log *slog.Logger,
	client *okta.APIClient,
) func(context.Context, *oktamcp.CallToolRequest, CreateUserInput) (*oktamcp.CallToolResult, CreateUserOutput, error) {
	return func(ctx context.Context, req *oktamcp.CallToolRequest, input CreateUserInput) (*oktamcp.CallToolResult, CreateUserOutput, error) {
		// Create the user profile
		profile := okta.NewUserProfile()
		profile.FirstName = *okta.NewNullableString(&input.FirstName)
		profile.LastName = *okta.NewNullableString(&input.LastName)
		profile.Email = &input.Email
		profile.Login = &input.Login

		createUserRequest := okta.CreateUserRequest{
			Profile: *profile,
		}

		// Create the user via Okta API
		user, _, err := client.UserAPI.CreateUser(ctx).
			Body(createUserRequest).
			Activate(true).
			Execute()
		if err != nil {
			log.Error("Failed to create user", "error", err)
			return &oktamcp.CallToolResult{
				Content: []oktamcp.Content{
					&oktamcp.TextContent{
						Text: fmt.Sprintf("Failed to create user: %s", err.Error()),
					},
				},
				IsError: true,
			}, CreateUserOutput{}, nil
		}

		log.Info("User created successfully", "login", input.Login)

		return &oktamcp.CallToolResult{
			Content: []oktamcp.Content{
				&oktamcp.TextContent{
					Text: fmt.Sprintf("User created successfully: %s", input.Login),
				},
			},
		}, CreateUserOutput{User: user}, nil
	}
}

// FetchUsersInput defines the input structure for fetching users (empty for now)
type FetchUsersInput struct {
	// No parameters needed for now
}

// FetchUsersOutput defines the output structure for fetching users
type FetchUsersOutput struct {
	Users []okta.User `json:"users" jsonschema:"description=List of users in the organization"`
}

// FetchUsers fetches users from the org
func FetchUsers(
	log *slog.Logger,
	client *okta.APIClient,
) func(context.Context, *oktamcp.CallToolRequest, FetchUsersInput) (*oktamcp.CallToolResult, FetchUsersOutput, error) {
	return func(ctx context.Context, req *oktamcp.CallToolRequest, input FetchUsersInput) (*oktamcp.CallToolResult, FetchUsersOutput, error) {
		users, _, err := client.UserAPI.ListUsers(ctx).Execute()
		if err != nil {
			log.Error("Failed to fetch users", "error", err)
			return &oktamcp.CallToolResult{
				Content: []oktamcp.Content{
					&oktamcp.TextContent{
						Text: fmt.Sprintf("Failed to fetch users: %s", err.Error()),
					},
				},
				IsError: true,
			}, FetchUsersOutput{}, nil
		}

		log.Info("Users fetched successfully", "count", len(users))

		return &oktamcp.CallToolResult{
			Content: []oktamcp.Content{
				&oktamcp.TextContent{
					Text: fmt.Sprintf("Successfully fetched %d users", len(users)),
				},
			},
		}, FetchUsersOutput{Users: users}, nil
	}
}
