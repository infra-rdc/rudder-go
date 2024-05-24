package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ListTechniqueVersionDirectivesRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
	// Technique version
	TechniqueVersion string `pathParam:"style=simple,explode=false,name=techniqueVersion"`
}

func (o *ListTechniqueVersionDirectivesRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

func (o *ListTechniqueVersionDirectivesRequest) GetTechniqueVersion() string {
	if o == nil {
		return ""
	}
	return o.TechniqueVersion
}

// ListTechniqueVersionDirectivesResult - Result of the request
type ListTechniqueVersionDirectivesResult string

const (
	ListTechniqueVersionDirectivesResultSuccess ListTechniqueVersionDirectivesResult = "success"
	ListTechniqueVersionDirectivesResultError   ListTechniqueVersionDirectivesResult = "error"
)

func (e ListTechniqueVersionDirectivesResult) ToPointer() *ListTechniqueVersionDirectivesResult {
	return &e
}
func (e *ListTechniqueVersionDirectivesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListTechniqueVersionDirectivesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniqueVersionDirectivesResult: %v", v)
	}
}

// ListTechniqueVersionDirectivesAction - The id of the action
type ListTechniqueVersionDirectivesAction string

const (
	ListTechniqueVersionDirectivesActionListTechniqueDirectives ListTechniqueVersionDirectivesAction = "listTechniqueDirectives"
)

func (e ListTechniqueVersionDirectivesAction) ToPointer() *ListTechniqueVersionDirectivesAction {
	return &e
}
func (e *ListTechniqueVersionDirectivesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniqueDirectives":
		*e = ListTechniqueVersionDirectivesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniqueVersionDirectivesAction: %v", v)
	}
}

type ListTechniqueVersionDirectivesData struct {
	Directives []components.Directive `json:"directives"`
}

func (o *ListTechniqueVersionDirectivesData) GetDirectives() []components.Directive {
	if o == nil {
		return []components.Directive{}
	}
	return o.Directives
}

// ListTechniqueVersionDirectivesResponseBody - Techniques information
type ListTechniqueVersionDirectivesResponseBody struct {
	// Result of the request
	Result ListTechniqueVersionDirectivesResult `json:"result"`
	// The id of the action
	Action ListTechniqueVersionDirectivesAction `json:"action"`
	Data   ListTechniqueVersionDirectivesData   `json:"data"`
}

func (o *ListTechniqueVersionDirectivesResponseBody) GetResult() ListTechniqueVersionDirectivesResult {
	if o == nil {
		return ListTechniqueVersionDirectivesResult("")
	}
	return o.Result
}

func (o *ListTechniqueVersionDirectivesResponseBody) GetAction() ListTechniqueVersionDirectivesAction {
	if o == nil {
		return ListTechniqueVersionDirectivesAction("")
	}
	return o.Action
}

func (o *ListTechniqueVersionDirectivesResponseBody) GetData() ListTechniqueVersionDirectivesData {
	if o == nil {
		return ListTechniqueVersionDirectivesData{}
	}
	return o.Data
}

type ListTechniqueVersionDirectivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *ListTechniqueVersionDirectivesResponseBody
}

func (o *ListTechniqueVersionDirectivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTechniqueVersionDirectivesResponse) GetObject() *ListTechniqueVersionDirectivesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
