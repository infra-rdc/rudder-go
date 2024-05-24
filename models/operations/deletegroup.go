package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteGroupRequest struct {
	// Id of the group
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

func (o *DeleteGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

// DeleteGroupResult - Result of the request
type DeleteGroupResult string

const (
	DeleteGroupResultSuccess DeleteGroupResult = "success"
	DeleteGroupResultError   DeleteGroupResult = "error"
)

func (e DeleteGroupResult) ToPointer() *DeleteGroupResult {
	return &e
}
func (e *DeleteGroupResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteGroupResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteGroupResult: %v", v)
	}
}

// DeleteGroupAction - The id of the action
type DeleteGroupAction string

const (
	DeleteGroupActionDeleteGroup DeleteGroupAction = "deleteGroup"
)

func (e DeleteGroupAction) ToPointer() *DeleteGroupAction {
	return &e
}
func (e *DeleteGroupAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteGroup":
		*e = DeleteGroupAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteGroupAction: %v", v)
	}
}

type DeleteGroupData struct {
	Groups []components.Group `json:"groups"`
}

func (o *DeleteGroupData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// DeleteGroupResponseBody - Groups information
type DeleteGroupResponseBody struct {
	// Result of the request
	Result DeleteGroupResult `json:"result"`
	// The id of the action
	Action DeleteGroupAction `json:"action"`
	Data   DeleteGroupData   `json:"data"`
}

func (o *DeleteGroupResponseBody) GetResult() DeleteGroupResult {
	if o == nil {
		return DeleteGroupResult("")
	}
	return o.Result
}

func (o *DeleteGroupResponseBody) GetAction() DeleteGroupAction {
	if o == nil {
		return DeleteGroupAction("")
	}
	return o.Action
}

func (o *DeleteGroupResponseBody) GetData() DeleteGroupData {
	if o == nil {
		return DeleteGroupData{}
	}
	return o.Data
}

type DeleteGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *DeleteGroupResponseBody
}

func (o *DeleteGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteGroupResponse) GetObject() *DeleteGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
