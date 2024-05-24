package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListTechniquesVersionsResult - Result of the request
type ListTechniquesVersionsResult string

const (
	ListTechniquesVersionsResultSuccess ListTechniquesVersionsResult = "success"
	ListTechniquesVersionsResultError   ListTechniquesVersionsResult = "error"
)

func (e ListTechniquesVersionsResult) ToPointer() *ListTechniquesVersionsResult {
	return &e
}
func (e *ListTechniquesVersionsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListTechniquesVersionsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesVersionsResult: %v", v)
	}
}

// ListTechniquesVersionsAction - The id of the action
type ListTechniquesVersionsAction string

const (
	ListTechniquesVersionsActionListTechniques ListTechniquesVersionsAction = "listTechniques"
)

func (e ListTechniquesVersionsAction) ToPointer() *ListTechniquesVersionsAction {
	return &e
}
func (e *ListTechniquesVersionsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = ListTechniquesVersionsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesVersionsAction: %v", v)
	}
}

type ListTechniquesVersionsData struct {
	Techniques []components.TechniquesVersions `json:"techniques"`
}

func (o *ListTechniquesVersionsData) GetTechniques() []components.TechniquesVersions {
	if o == nil {
		return []components.TechniquesVersions{}
	}
	return o.Techniques
}

// ListTechniquesVersionsResponseBody - Versions information
type ListTechniquesVersionsResponseBody struct {
	// Result of the request
	Result ListTechniquesVersionsResult `json:"result"`
	// The id of the action
	Action ListTechniquesVersionsAction `json:"action"`
	Data   ListTechniquesVersionsData   `json:"data"`
}

func (o *ListTechniquesVersionsResponseBody) GetResult() ListTechniquesVersionsResult {
	if o == nil {
		return ListTechniquesVersionsResult("")
	}
	return o.Result
}

func (o *ListTechniquesVersionsResponseBody) GetAction() ListTechniquesVersionsAction {
	if o == nil {
		return ListTechniquesVersionsAction("")
	}
	return o.Action
}

func (o *ListTechniquesVersionsResponseBody) GetData() ListTechniquesVersionsData {
	if o == nil {
		return ListTechniquesVersionsData{}
	}
	return o.Data
}

type ListTechniquesVersionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Versions information
	Object *ListTechniquesVersionsResponseBody
}

func (o *ListTechniquesVersionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTechniquesVersionsResponse) GetObject() *ListTechniquesVersionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
