package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// Result of the request
type Result string

const (
	ResultSuccess Result = "success"
	ResultError   Result = "error"
)

func (e Result) ToPointer() *Result {
	return &e
}
func (e *Result) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = Result(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Result: %v", v)
	}
}

// Action - The id of the action
type Action string

const (
	ActionListChangeRequests Action = "listChangeRequests"
)

func (e Action) ToPointer() *Action {
	return &e
}
func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listChangeRequests":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

type Data struct {
	Rules []components.ChangeRequest `json:"rules"`
}

func (o *Data) GetRules() []components.ChangeRequest {
	if o == nil {
		return []components.ChangeRequest{}
	}
	return o.Rules
}

// ListChangeRequestsResponseBody - Change requests information
type ListChangeRequestsResponseBody struct {
	// Result of the request
	Result Result `json:"result"`
	// The id of the action
	Action Action `json:"action"`
	Data   Data   `json:"data"`
}

func (o *ListChangeRequestsResponseBody) GetResult() Result {
	if o == nil {
		return Result("")
	}
	return o.Result
}

func (o *ListChangeRequestsResponseBody) GetAction() Action {
	if o == nil {
		return Action("")
	}
	return o.Action
}

func (o *ListChangeRequestsResponseBody) GetData() Data {
	if o == nil {
		return Data{}
	}
	return o.Data
}

type ListChangeRequestsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Change requests information
	Object *ListChangeRequestsResponseBody
}

func (o *ListChangeRequestsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListChangeRequestsResponse) GetObject() *ListChangeRequestsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
