package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListDirectivesResult - Result of the request
type ListDirectivesResult string

const (
	ListDirectivesResultSuccess ListDirectivesResult = "success"
	ListDirectivesResultError   ListDirectivesResult = "error"
)

func (e ListDirectivesResult) ToPointer() *ListDirectivesResult {
	return &e
}
func (e *ListDirectivesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListDirectivesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectivesResult: %v", v)
	}
}

// ListDirectivesAction - The id of the action
type ListDirectivesAction string

const (
	ListDirectivesActionListDirectives ListDirectivesAction = "listDirectives"
)

func (e ListDirectivesAction) ToPointer() *ListDirectivesAction {
	return &e
}
func (e *ListDirectivesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listDirectives":
		*e = ListDirectivesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDirectivesAction: %v", v)
	}
}

type ListDirectivesData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *ListDirectivesData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// ListDirectivesResponseBody - Directives information
type ListDirectivesResponseBody struct {
	// Result of the request
	Result ListDirectivesResult `json:"result"`
	// The id of the action
	Action ListDirectivesAction `json:"action"`
	Data   ListDirectivesData   `json:"data"`
}

func (o *ListDirectivesResponseBody) GetResult() ListDirectivesResult {
	if o == nil {
		return ListDirectivesResult("")
	}
	return o.Result
}

func (o *ListDirectivesResponseBody) GetAction() ListDirectivesAction {
	if o == nil {
		return ListDirectivesAction("")
	}
	return o.Action
}

func (o *ListDirectivesResponseBody) GetData() ListDirectivesData {
	if o == nil {
		return ListDirectivesData{}
	}
	return o.Data
}

type ListDirectivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Directives information
	Object *ListDirectivesResponseBody
}

func (o *ListDirectivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListDirectivesResponse) GetObject() *ListDirectivesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
