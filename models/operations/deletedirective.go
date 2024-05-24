package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteDirectiveRequest struct {
	// Id of the directive
	DirectiveID string `pathParam:"style=simple,explode=false,name=directiveId"`
}

func (o *DeleteDirectiveRequest) GetDirectiveID() string {
	if o == nil {
		return ""
	}
	return o.DirectiveID
}

// DeleteDirectiveResult - Result of the request
type DeleteDirectiveResult string

const (
	DeleteDirectiveResultSuccess DeleteDirectiveResult = "success"
	DeleteDirectiveResultError   DeleteDirectiveResult = "error"
)

func (e DeleteDirectiveResult) ToPointer() *DeleteDirectiveResult {
	return &e
}
func (e *DeleteDirectiveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteDirectiveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteDirectiveResult: %v", v)
	}
}

// DeleteDirectiveAction - The id of the action
type DeleteDirectiveAction string

const (
	DeleteDirectiveActionDeleteDirective DeleteDirectiveAction = "deleteDirective"
)

func (e DeleteDirectiveAction) ToPointer() *DeleteDirectiveAction {
	return &e
}
func (e *DeleteDirectiveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteDirective":
		*e = DeleteDirectiveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteDirectiveAction: %v", v)
	}
}

type DeleteDirectiveData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *DeleteDirectiveData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// DeleteDirectiveResponseBody - Directives information
type DeleteDirectiveResponseBody struct {
	// Result of the request
	Result DeleteDirectiveResult `json:"result"`
	// The id of the action
	Action DeleteDirectiveAction `json:"action"`
	Data   DeleteDirectiveData   `json:"data"`
}

func (o *DeleteDirectiveResponseBody) GetResult() DeleteDirectiveResult {
	if o == nil {
		return DeleteDirectiveResult("")
	}
	return o.Result
}

func (o *DeleteDirectiveResponseBody) GetAction() DeleteDirectiveAction {
	if o == nil {
		return DeleteDirectiveAction("")
	}
	return o.Action
}

func (o *DeleteDirectiveResponseBody) GetData() DeleteDirectiveData {
	if o == nil {
		return DeleteDirectiveData{}
	}
	return o.Data
}

type DeleteDirectiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Directives information
	Object *DeleteDirectiveResponseBody
}

func (o *DeleteDirectiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteDirectiveResponse) GetObject() *DeleteDirectiveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
