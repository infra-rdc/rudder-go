package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GroupDetailsRequest struct {
	// Id of the group
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

func (o *GroupDetailsRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

// GroupDetailsResult - Result of the request
type GroupDetailsResult string

const (
	GroupDetailsResultSuccess GroupDetailsResult = "success"
	GroupDetailsResultError   GroupDetailsResult = "error"
)

func (e GroupDetailsResult) ToPointer() *GroupDetailsResult {
	return &e
}
func (e *GroupDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GroupDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupDetailsResult: %v", v)
	}
}

// GroupDetailsAction - The id of the action
type GroupDetailsAction string

const (
	GroupDetailsActionGroupDetails GroupDetailsAction = "groupDetails"
)

func (e GroupDetailsAction) ToPointer() *GroupDetailsAction {
	return &e
}
func (e *GroupDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "groupDetails":
		*e = GroupDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupDetailsAction: %v", v)
	}
}

type GroupDetailsData struct {
	Groups []components.Group `json:"groups"`
}

func (o *GroupDetailsData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// GroupDetailsResponseBody - Groups information
type GroupDetailsResponseBody struct {
	// Result of the request
	Result GroupDetailsResult `json:"result"`
	// The id of the action
	Action GroupDetailsAction `json:"action"`
	Data   GroupDetailsData   `json:"data"`
}

func (o *GroupDetailsResponseBody) GetResult() GroupDetailsResult {
	if o == nil {
		return GroupDetailsResult("")
	}
	return o.Result
}

func (o *GroupDetailsResponseBody) GetAction() GroupDetailsAction {
	if o == nil {
		return GroupDetailsAction("")
	}
	return o.Action
}

func (o *GroupDetailsResponseBody) GetData() GroupDetailsData {
	if o == nil {
		return GroupDetailsData{}
	}
	return o.Data
}

type GroupDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *GroupDetailsResponseBody
}

func (o *GroupDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GroupDetailsResponse) GetObject() *GroupDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
