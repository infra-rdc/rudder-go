package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ReloadGroupRequest struct {
	// Id of the group
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

func (o *ReloadGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

// ReloadGroupResult - Result of the request
type ReloadGroupResult string

const (
	ReloadGroupResultSuccess ReloadGroupResult = "success"
	ReloadGroupResultError   ReloadGroupResult = "error"
)

func (e ReloadGroupResult) ToPointer() *ReloadGroupResult {
	return &e
}
func (e *ReloadGroupResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadGroupResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadGroupResult: %v", v)
	}
}

// ReloadGroupAction - The id of the action
type ReloadGroupAction string

const (
	ReloadGroupActionReloadGroup ReloadGroupAction = "reloadGroup"
)

func (e ReloadGroupAction) ToPointer() *ReloadGroupAction {
	return &e
}
func (e *ReloadGroupAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reloadGroup":
		*e = ReloadGroupAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadGroupAction: %v", v)
	}
}

type ReloadGroupData struct {
	Groups []components.Group `json:"groups"`
}

func (o *ReloadGroupData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// ReloadGroupResponseBody - Groups information
type ReloadGroupResponseBody struct {
	// Result of the request
	Result ReloadGroupResult `json:"result"`
	// The id of the action
	Action ReloadGroupAction `json:"action"`
	Data   ReloadGroupData   `json:"data"`
}

func (o *ReloadGroupResponseBody) GetResult() ReloadGroupResult {
	if o == nil {
		return ReloadGroupResult("")
	}
	return o.Result
}

func (o *ReloadGroupResponseBody) GetAction() ReloadGroupAction {
	if o == nil {
		return ReloadGroupAction("")
	}
	return o.Action
}

func (o *ReloadGroupResponseBody) GetData() ReloadGroupData {
	if o == nil {
		return ReloadGroupData{}
	}
	return o.Data
}

type ReloadGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *ReloadGroupResponseBody
}

func (o *ReloadGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadGroupResponse) GetObject() *ReloadGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
