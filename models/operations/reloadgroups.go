package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadGroupsResult - Result of the request
type ReloadGroupsResult string

const (
	ReloadGroupsResultSuccess ReloadGroupsResult = "success"
	ReloadGroupsResultError   ReloadGroupsResult = "error"
)

func (e ReloadGroupsResult) ToPointer() *ReloadGroupsResult {
	return &e
}
func (e *ReloadGroupsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadGroupsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadGroupsResult: %v", v)
	}
}

// ReloadGroupsAction - The id of the action
type ReloadGroupsAction string

const (
	ReloadGroupsActionReloadGroups ReloadGroupsAction = "reloadGroups"
)

func (e ReloadGroupsAction) ToPointer() *ReloadGroupsAction {
	return &e
}
func (e *ReloadGroupsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reloadGroups":
		*e = ReloadGroupsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadGroupsAction: %v", v)
	}
}

type ReloadGroupsGroups string

const (
	ReloadGroupsGroupsStarted ReloadGroupsGroups = "Started"
)

func (e ReloadGroupsGroups) ToPointer() *ReloadGroupsGroups {
	return &e
}
func (e *ReloadGroupsGroups) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = ReloadGroupsGroups(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadGroupsGroups: %v", v)
	}
}

type ReloadGroupsData struct {
	Groups ReloadGroupsGroups `json:"groups"`
}

func (o *ReloadGroupsData) GetGroups() ReloadGroupsGroups {
	if o == nil {
		return ReloadGroupsGroups("")
	}
	return o.Groups
}

// ReloadGroupsResponseBody - Service reload
type ReloadGroupsResponseBody struct {
	// Result of the request
	Result ReloadGroupsResult `json:"result"`
	// The id of the action
	Action ReloadGroupsAction `json:"action"`
	Data   ReloadGroupsData   `json:"data"`
}

func (o *ReloadGroupsResponseBody) GetResult() ReloadGroupsResult {
	if o == nil {
		return ReloadGroupsResult("")
	}
	return o.Result
}

func (o *ReloadGroupsResponseBody) GetAction() ReloadGroupsAction {
	if o == nil {
		return ReloadGroupsAction("")
	}
	return o.Action
}

func (o *ReloadGroupsResponseBody) GetData() ReloadGroupsData {
	if o == nil {
		return ReloadGroupsData{}
	}
	return o.Data
}

type ReloadGroupsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Service reload
	Object *ReloadGroupsResponseBody
}

func (o *ReloadGroupsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadGroupsResponse) GetObject() *ReloadGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
