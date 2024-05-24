package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// MethodsResult - Result of the request
type MethodsResult string

const (
	MethodsResultSuccess MethodsResult = "success"
	MethodsResultError   MethodsResult = "error"
)

func (e MethodsResult) ToPointer() *MethodsResult {
	return &e
}
func (e *MethodsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = MethodsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MethodsResult: %v", v)
	}
}

// MethodsAction - The id of the action
type MethodsAction string

const (
	MethodsActionListTechniques MethodsAction = "listTechniques"
)

func (e MethodsAction) ToPointer() *MethodsAction {
	return &e
}
func (e *MethodsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = MethodsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MethodsAction: %v", v)
	}
}

type MethodsData struct {
	// List of available generic methods
	Methods components.Methods `json:"methods"`
}

func (o *MethodsData) GetMethods() components.Methods {
	if o == nil {
		return components.Methods{}
	}
	return o.Methods
}

// MethodsResponseBody - Methods information
type MethodsResponseBody struct {
	// Result of the request
	Result MethodsResult `json:"result"`
	// The id of the action
	Action MethodsAction `json:"action"`
	Data   MethodsData   `json:"data"`
}

func (o *MethodsResponseBody) GetResult() MethodsResult {
	if o == nil {
		return MethodsResult("")
	}
	return o.Result
}

func (o *MethodsResponseBody) GetAction() MethodsAction {
	if o == nil {
		return MethodsAction("")
	}
	return o.Action
}

func (o *MethodsResponseBody) GetData() MethodsData {
	if o == nil {
		return MethodsData{}
	}
	return o.Data
}

type MethodsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Methods information
	Object *MethodsResponseBody
}

func (o *MethodsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MethodsResponse) GetObject() *MethodsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
