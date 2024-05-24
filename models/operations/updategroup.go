package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateGroupRequest struct {
	// Id of the group
	GroupID     string                 `pathParam:"style=simple,explode=false,name=groupId"`
	GroupUpdate components.GroupUpdate `request:"mediaType=application/json"`
}

func (o *UpdateGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *UpdateGroupRequest) GetGroupUpdate() components.GroupUpdate {
	if o == nil {
		return components.GroupUpdate{}
	}
	return o.GroupUpdate
}

// UpdateGroupResult - Result of the request
type UpdateGroupResult string

const (
	UpdateGroupResultSuccess UpdateGroupResult = "success"
	UpdateGroupResultError   UpdateGroupResult = "error"
)

func (e UpdateGroupResult) ToPointer() *UpdateGroupResult {
	return &e
}
func (e *UpdateGroupResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateGroupResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateGroupResult: %v", v)
	}
}

// UpdateGroupAction - The id of the action
type UpdateGroupAction string

const (
	UpdateGroupActionUpdateGroup UpdateGroupAction = "updateGroup"
)

func (e UpdateGroupAction) ToPointer() *UpdateGroupAction {
	return &e
}
func (e *UpdateGroupAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateGroup":
		*e = UpdateGroupAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateGroupAction: %v", v)
	}
}

type UpdateGroupData struct {
	Groups []components.Group `json:"groups"`
}

func (o *UpdateGroupData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// UpdateGroupResponseBody - Groups information
type UpdateGroupResponseBody struct {
	// Result of the request
	Result UpdateGroupResult `json:"result"`
	// The id of the action
	Action UpdateGroupAction `json:"action"`
	Data   UpdateGroupData   `json:"data"`
}

func (o *UpdateGroupResponseBody) GetResult() UpdateGroupResult {
	if o == nil {
		return UpdateGroupResult("")
	}
	return o.Result
}

func (o *UpdateGroupResponseBody) GetAction() UpdateGroupAction {
	if o == nil {
		return UpdateGroupAction("")
	}
	return o.Action
}

func (o *UpdateGroupResponseBody) GetData() UpdateGroupData {
	if o == nil {
		return UpdateGroupData{}
	}
	return o.Data
}

type UpdateGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *UpdateGroupResponseBody
}

func (o *UpdateGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateGroupResponse) GetObject() *UpdateGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
