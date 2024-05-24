package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListParametersResult - Result of the request
type ListParametersResult string

const (
	ListParametersResultSuccess ListParametersResult = "success"
	ListParametersResultError   ListParametersResult = "error"
)

func (e ListParametersResult) ToPointer() *ListParametersResult {
	return &e
}
func (e *ListParametersResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListParametersResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListParametersResult: %v", v)
	}
}

// ListParametersAction - The id of the action
type ListParametersAction string

const (
	ListParametersActionListParameters ListParametersAction = "listParameters"
)

func (e ListParametersAction) ToPointer() *ListParametersAction {
	return &e
}
func (e *ListParametersAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listParameters":
		*e = ListParametersAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListParametersAction: %v", v)
	}
}

// ListParametersData - Parameters
type ListParametersData struct {
	Parameters []components.Parameter `json:"parameters"`
}

func (o *ListParametersData) GetParameters() []components.Parameter {
	if o == nil {
		return []components.Parameter{}
	}
	return o.Parameters
}

// ListParametersResponseBody - Settings
type ListParametersResponseBody struct {
	// Result of the request
	Result ListParametersResult `json:"result"`
	// The id of the action
	Action ListParametersAction `json:"action"`
	// Parameters
	Data ListParametersData `json:"data"`
}

func (o *ListParametersResponseBody) GetResult() ListParametersResult {
	if o == nil {
		return ListParametersResult("")
	}
	return o.Result
}

func (o *ListParametersResponseBody) GetAction() ListParametersAction {
	if o == nil {
		return ListParametersAction("")
	}
	return o.Action
}

func (o *ListParametersResponseBody) GetData() ListParametersData {
	if o == nil {
		return ListParametersData{}
	}
	return o.Data
}

type ListParametersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *ListParametersResponseBody
}

func (o *ListParametersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListParametersResponse) GetObject() *ListParametersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
