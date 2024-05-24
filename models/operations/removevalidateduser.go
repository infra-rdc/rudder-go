package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type RemoveValidatedUserRequest struct {
	// Username of an user (unique)
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *RemoveValidatedUserRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// RemoveValidatedUserResult - Result of the request
type RemoveValidatedUserResult string

const (
	RemoveValidatedUserResultSuccess RemoveValidatedUserResult = "success"
	RemoveValidatedUserResultError   RemoveValidatedUserResult = "error"
)

func (e RemoveValidatedUserResult) ToPointer() *RemoveValidatedUserResult {
	return &e
}
func (e *RemoveValidatedUserResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = RemoveValidatedUserResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RemoveValidatedUserResult: %v", v)
	}
}

// RemoveValidatedUserAction - The id of the action
type RemoveValidatedUserAction string

const (
	RemoveValidatedUserActionListUsers RemoveValidatedUserAction = "listUsers"
)

func (e RemoveValidatedUserAction) ToPointer() *RemoveValidatedUserAction {
	return &e
}
func (e *RemoveValidatedUserAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listUsers":
		*e = RemoveValidatedUserAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RemoveValidatedUserAction: %v", v)
	}
}

// RemoveValidatedUserResponseBody - Removed user
type RemoveValidatedUserResponseBody struct {
	// Result of the request
	Result RemoveValidatedUserResult `json:"result"`
	// The id of the action
	Action RemoveValidatedUserAction `json:"action"`
	// the user removed from validated list
	Data string `json:"data"`
}

func (o *RemoveValidatedUserResponseBody) GetResult() RemoveValidatedUserResult {
	if o == nil {
		return RemoveValidatedUserResult("")
	}
	return o.Result
}

func (o *RemoveValidatedUserResponseBody) GetAction() RemoveValidatedUserAction {
	if o == nil {
		return RemoveValidatedUserAction("")
	}
	return o.Action
}

func (o *RemoveValidatedUserResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type RemoveValidatedUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Removed user
	Object *RemoveValidatedUserResponseBody
}

func (o *RemoveValidatedUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RemoveValidatedUserResponse) GetObject() *RemoveValidatedUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
