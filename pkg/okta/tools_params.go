package okta

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/radar07/mcp-server-okta/internal/oktamcp"
)

// Validator provides a fluent interface for validating parameters
// and collecting errors
type Validator struct {
	request *oktamcp.CallToolRequest
	errors  []error
}

// NewValidator creates a new validator for the given request
func NewValidator(r *oktamcp.CallToolRequest) *Validator {
	return &Validator{
		request: r,
		errors:  []error{},
	}
}

// HasErrors returns true if there are any validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// HandleErrorsIfAny formats all errors and returns an appropriate tool result
func (v *Validator) HandleErrorsIfAny() (*oktamcp.ToolResult, error) {
	if v.HasErrors() {
		messages := make([]string, 0, len(v.errors))
		for _, err := range v.errors {
			messages = append(messages, err.Error())
		}
		errorMsg := "Validation errors:\n- " + strings.Join(messages, "\n- ")
		return oktamcp.NewToolResultError(errorMsg), nil
	}
	return nil, nil
}

// addError adds a non-nil error to the collection
func (v *Validator) addError(err error) *Validator {
	if err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}

// extractValueGeneric is a standalone generic function to extract a parameter
// of type T
func extractValueGeneric[T any](
	request *oktamcp.CallToolRequest,
	name string,
	required bool,
) (*T, error) {
	val, ok := request.Arguments[name]
	if !ok || val == nil {
		if required {
			return nil, errors.New("missing required parameter: " + name)
		}
		return nil, nil // Not an error for optional params
	}

	var result T
	data, err := json.Marshal(val)
	if err != nil {
		return nil, errors.New("invalid parameter type: " + name)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("invalid parameter type: " + name)
	}

	return &result, nil
}

// ValidateAndAddRequiredString validates and adds a required string parameter
func (v *Validator) ValidateAndAddRequiredString(
	params map[string]interface{},
	name string,
) *Validator {
	return validateAndAddRequired[string](v, params, name)
}

// validateAndAddRequired validates and adds a required parameter of any type
func validateAndAddRequired[T any](
	v *Validator,
	params map[string]interface{},
	name string,
) *Validator {
	value, err := extractValueGeneric[T](v.request, name, true)
	if err != nil {
		return v.addError(err)
	}

	if value == nil {
		return v
	}

	params[name] = *value
	return v
}
