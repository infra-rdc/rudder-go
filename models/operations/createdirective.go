package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// CreateDirectiveResult - Result of the request
type CreateDirectiveResult string

const (
	CreateDirectiveResultSuccess CreateDirectiveResult = "success"
	CreateDirectiveResultError   CreateDirectiveResult = "error"
)

func (e CreateDirectiveResult) ToPointer() *CreateDirectiveResult {
	return &e
}
func (e *CreateDirectiveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = CreateDirectiveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDirectiveResult: %v", v)
	}
}

// CreateDirectiveAction - The id of the action
type CreateDirectiveAction string

const (
	CreateDirectiveActionCreateDirective CreateDirectiveAction = "createDirective"
)

func (e CreateDirectiveAction) ToPointer() *CreateDirectiveAction {
	return &e
}
func (e *CreateDirectiveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "createDirective":
		*e = CreateDirectiveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDirectiveAction: %v", v)
	}
}

type CreateDirectiveData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *CreateDirectiveData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// CreateDirectiveResponseBody - Directives information
type CreateDirectiveResponseBody struct {
	// Result of the request
	Result CreateDirectiveResult `json:"result"`
	// The id of the action
	Action CreateDirectiveAction `json:"action"`
	Data   CreateDirectiveData   `json:"data"`
}

func (o *CreateDirectiveResponseBody) GetResult() CreateDirectiveResult {
	if o == nil {
		return CreateDirectiveResult("")
	}
	return o.Result
}

func (o *CreateDirectiveResponseBody) GetAction() CreateDirectiveAction {
	if o == nil {
		return CreateDirectiveAction("")
	}
	return o.Action
}

func (o *CreateDirectiveResponseBody) GetData() CreateDirectiveData {
	if o == nil {
		return CreateDirectiveData{}
	}
	return o.Data
}

type CreateDirectiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Directives information
	Object *CreateDirectiveResponseBody
}

func (o *CreateDirectiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateDirectiveResponse) GetObject() *CreateDirectiveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
