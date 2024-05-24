package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListGroupsResult - Result of the request
type ListGroupsResult string

const (
	ListGroupsResultSuccess ListGroupsResult = "success"
	ListGroupsResultError   ListGroupsResult = "error"
)

func (e ListGroupsResult) ToPointer() *ListGroupsResult {
	return &e
}
func (e *ListGroupsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListGroupsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListGroupsResult: %v", v)
	}
}

// ListGroupsAction - The id of the action
type ListGroupsAction string

const (
	ListGroupsActionListGroups ListGroupsAction = "listGroups"
)

func (e ListGroupsAction) ToPointer() *ListGroupsAction {
	return &e
}
func (e *ListGroupsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listGroups":
		*e = ListGroupsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListGroupsAction: %v", v)
	}
}

type ListGroupsData struct {
	Groups []components.Group `json:"groups"`
}

func (o *ListGroupsData) GetGroups() []components.Group {
	if o == nil {
		return []components.Group{}
	}
	return o.Groups
}

// ListGroupsResponseBody - Groups information
type ListGroupsResponseBody struct {
	// Result of the request
	Result ListGroupsResult `json:"result"`
	// The id of the action
	Action ListGroupsAction `json:"action"`
	Data   ListGroupsData   `json:"data"`
}

func (o *ListGroupsResponseBody) GetResult() ListGroupsResult {
	if o == nil {
		return ListGroupsResult("")
	}
	return o.Result
}

func (o *ListGroupsResponseBody) GetAction() ListGroupsAction {
	if o == nil {
		return ListGroupsAction("")
	}
	return o.Action
}

func (o *ListGroupsResponseBody) GetData() ListGroupsData {
	if o == nil {
		return ListGroupsData{}
	}
	return o.Data
}

type ListGroupsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *ListGroupsResponseBody
}

func (o *ListGroupsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListGroupsResponse) GetObject() *ListGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
