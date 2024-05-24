package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// CreateGroupResult - Result of the request
type CreateGroupResult string

const (
	CreateGroupResultSuccess CreateGroupResult = "success"
	CreateGroupResultError   CreateGroupResult = "error"
)

func (e CreateGroupResult) ToPointer() *CreateGroupResult {
	return &e
}
func (e *CreateGroupResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = CreateGroupResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateGroupResult: %v", v)
	}
}

// CreateGroupAction - The id of the action
type CreateGroupAction string

const (
	CreateGroupActionCreateGroup CreateGroupAction = "createGroup"
)

func (e CreateGroupAction) ToPointer() *CreateGroupAction {
	return &e
}
func (e *CreateGroupAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "createGroup":
		*e = CreateGroupAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateGroupAction: %v", v)
	}
}

type CreateGroupData struct {
	Groups []components.Group `json:"groups"`
}

func (o *CreateGroupData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// CreateGroupResponseBody - Group information
type CreateGroupResponseBody struct {
	// Result of the request
	Result CreateGroupResult `json:"result"`
	// The id of the action
	Action CreateGroupAction `json:"action"`
	Data   CreateGroupData   `json:"data"`
}

func (o *CreateGroupResponseBody) GetResult() CreateGroupResult {
	if o == nil {
		return CreateGroupResult("")
	}
	return o.Result
}

func (o *CreateGroupResponseBody) GetAction() CreateGroupAction {
	if o == nil {
		return CreateGroupAction("")
	}
	return o.Action
}

func (o *CreateGroupResponseBody) GetData() CreateGroupData {
	if o == nil {
		return CreateGroupData{}
	}
	return o.Data
}

type CreateGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Group information
	Object *CreateGroupResponseBody
}

func (o *CreateGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateGroupResponse) GetObject() *CreateGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
