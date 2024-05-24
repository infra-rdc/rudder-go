package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ListTechniquesDirectivesRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
}

func (o *ListTechniquesDirectivesRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

// ListTechniquesDirectivesResult - Result of the request
type ListTechniquesDirectivesResult string

const (
	ListTechniquesDirectivesResultSuccess ListTechniquesDirectivesResult = "success"
	ListTechniquesDirectivesResultError   ListTechniquesDirectivesResult = "error"
)

func (e ListTechniquesDirectivesResult) ToPointer() *ListTechniquesDirectivesResult {
	return &e
}
func (e *ListTechniquesDirectivesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListTechniquesDirectivesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesDirectivesResult: %v", v)
	}
}

// ListTechniquesDirectivesAction - The id of the action
type ListTechniquesDirectivesAction string

const (
	ListTechniquesDirectivesActionListTechniquesDirectives ListTechniquesDirectivesAction = "listTechniquesDirectives"
)

func (e ListTechniquesDirectivesAction) ToPointer() *ListTechniquesDirectivesAction {
	return &e
}
func (e *ListTechniquesDirectivesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniquesDirectives":
		*e = ListTechniquesDirectivesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesDirectivesAction: %v", v)
	}
}

type ListTechniquesDirectivesData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *ListTechniquesDirectivesData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// ListTechniquesDirectivesResponseBody - Techniques information
type ListTechniquesDirectivesResponseBody struct {
	// Result of the request
	Result ListTechniquesDirectivesResult `json:"result"`
	// The id of the action
	Action ListTechniquesDirectivesAction `json:"action"`
	Data   ListTechniquesDirectivesData   `json:"data"`
}

func (o *ListTechniquesDirectivesResponseBody) GetResult() ListTechniquesDirectivesResult {
	if o == nil {
		return ListTechniquesDirectivesResult("")
	}
	return o.Result
}

func (o *ListTechniquesDirectivesResponseBody) GetAction() ListTechniquesDirectivesAction {
	if o == nil {
		return ListTechniquesDirectivesAction("")
	}
	return o.Action
}

func (o *ListTechniquesDirectivesResponseBody) GetData() ListTechniquesDirectivesData {
	if o == nil {
		return ListTechniquesDirectivesData{}
	}
	return o.Data
}

type ListTechniquesDirectivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *ListTechniquesDirectivesResponseBody
}

func (o *ListTechniquesDirectivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTechniquesDirectivesResponse) GetObject() *ListTechniquesDirectivesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
