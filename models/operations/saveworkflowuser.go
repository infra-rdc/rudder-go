package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type SaveWorkflowUserRequestBody struct {
	// list of user to put in validated list
	ValidatedUsers []string `json:"validatedUsers"`
}

func (o *SaveWorkflowUserRequestBody) GetValidatedUsers() []string {
	if o == nil {
		return []string{}
	}
	return o.ValidatedUsers
}

// SaveWorkflowUserResult - Result of the request
type SaveWorkflowUserResult string

const (
	SaveWorkflowUserResultSuccess SaveWorkflowUserResult = "success"
	SaveWorkflowUserResultError   SaveWorkflowUserResult = "error"
)

func (e SaveWorkflowUserResult) ToPointer() *SaveWorkflowUserResult {
	return &e
}
func (e *SaveWorkflowUserResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = SaveWorkflowUserResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveWorkflowUserResult: %v", v)
	}
}

// SaveWorkflowUserAction - The id of the action
type SaveWorkflowUserAction string

const (
	SaveWorkflowUserActionAddUser SaveWorkflowUserAction = "addUser"
)

func (e SaveWorkflowUserAction) ToPointer() *SaveWorkflowUserAction {
	return &e
}
func (e *SaveWorkflowUserAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "addUser":
		*e = SaveWorkflowUserAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SaveWorkflowUserAction: %v", v)
	}
}

// SaveWorkflowUserResponseBody - Updated
type SaveWorkflowUserResponseBody struct {
	// Result of the request
	Result SaveWorkflowUserResult `json:"result"`
	// The id of the action
	Action SaveWorkflowUserAction `json:"action"`
	// list of users with their workflow settings
	Data components.ValidatedUser `json:"data"`
}

func (o *SaveWorkflowUserResponseBody) GetResult() SaveWorkflowUserResult {
	if o == nil {
		return SaveWorkflowUserResult("")
	}
	return o.Result
}

func (o *SaveWorkflowUserResponseBody) GetAction() SaveWorkflowUserAction {
	if o == nil {
		return SaveWorkflowUserAction("")
	}
	return o.Action
}

func (o *SaveWorkflowUserResponseBody) GetData() components.ValidatedUser {
	if o == nil {
		return components.ValidatedUser{}
	}
	return o.Data
}

type SaveWorkflowUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updated
	Object *SaveWorkflowUserResponseBody
}

func (o *SaveWorkflowUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SaveWorkflowUserResponse) GetObject() *SaveWorkflowUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
