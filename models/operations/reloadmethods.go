package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadMethodsResult - Result of the request
type ReloadMethodsResult string

const (
	ReloadMethodsResultSuccess ReloadMethodsResult = "success"
	ReloadMethodsResultError   ReloadMethodsResult = "error"
)

func (e ReloadMethodsResult) ToPointer() *ReloadMethodsResult {
	return &e
}
func (e *ReloadMethodsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadMethodsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadMethodsResult: %v", v)
	}
}

// ReloadMethodsAction - The id of the action
type ReloadMethodsAction string

const (
	ReloadMethodsActionListTechniques ReloadMethodsAction = "listTechniques"
)

func (e ReloadMethodsAction) ToPointer() *ReloadMethodsAction {
	return &e
}
func (e *ReloadMethodsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = ReloadMethodsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadMethodsAction: %v", v)
	}
}

type ReloadMethodsData struct {
	// List of available generic methods
	Methods components.Methods `json:"methods"`
}

func (o *ReloadMethodsData) GetMethods() components.Methods {
	if o == nil {
		return components.Methods{}
	}
	return o.Methods
}

// ReloadMethodsResponseBody - Methods information
type ReloadMethodsResponseBody struct {
	// Result of the request
	Result ReloadMethodsResult `json:"result"`
	// The id of the action
	Action ReloadMethodsAction `json:"action"`
	Data   ReloadMethodsData   `json:"data"`
}

func (o *ReloadMethodsResponseBody) GetResult() ReloadMethodsResult {
	if o == nil {
		return ReloadMethodsResult("")
	}
	return o.Result
}

func (o *ReloadMethodsResponseBody) GetAction() ReloadMethodsAction {
	if o == nil {
		return ReloadMethodsAction("")
	}
	return o.Action
}

func (o *ReloadMethodsResponseBody) GetData() ReloadMethodsData {
	if o == nil {
		return ReloadMethodsData{}
	}
	return o.Data
}

type ReloadMethodsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Methods information
	Object *ReloadMethodsResponseBody
}

func (o *ReloadMethodsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadMethodsResponse) GetObject() *ReloadMethodsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
