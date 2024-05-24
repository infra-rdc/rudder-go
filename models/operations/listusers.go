package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListUsersResult - Result of the request
type ListUsersResult string

const (
	ListUsersResultSuccess ListUsersResult = "success"
	ListUsersResultError   ListUsersResult = "error"
)

func (e ListUsersResult) ToPointer() *ListUsersResult {
	return &e
}
func (e *ListUsersResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListUsersResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersResult: %v", v)
	}
}

// ListUsersAction - The id of the action
type ListUsersAction string

const (
	ListUsersActionListUsers ListUsersAction = "listUsers"
)

func (e ListUsersAction) ToPointer() *ListUsersAction {
	return &e
}
func (e *ListUsersAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listUsers":
		*e = ListUsersAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersAction: %v", v)
	}
}

// ListUsersResponseBody - List users
type ListUsersResponseBody struct {
	// Result of the request
	Result ListUsersResult `json:"result"`
	// The id of the action
	Action ListUsersAction            `json:"action"`
	Data   []components.ValidatedUser `json:"data"`
}

func (o *ListUsersResponseBody) GetResult() ListUsersResult {
	if o == nil {
		return ListUsersResult("")
	}
	return o.Result
}

func (o *ListUsersResponseBody) GetAction() ListUsersAction {
	if o == nil {
		return ListUsersAction("")
	}
	return o.Action
}

func (o *ListUsersResponseBody) GetData() []components.ValidatedUser {
	if o == nil {
		return []components.ValidatedUser{}
	}
	return o.Data
}

type ListUsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List users
	Object *ListUsersResponseBody
}

func (o *ListUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListUsersResponse) GetObject() *ListUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
