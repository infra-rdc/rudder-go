package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ListTechniquesResult - Result of the request
type ListTechniquesResult string

const (
	ListTechniquesResultSuccess ListTechniquesResult = "success"
	ListTechniquesResultError   ListTechniquesResult = "error"
)

func (e ListTechniquesResult) ToPointer() *ListTechniquesResult {
	return &e
}
func (e *ListTechniquesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ListTechniquesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesResult: %v", v)
	}
}

// ListTechniquesAction - The id of the action
type ListTechniquesAction string

const (
	ListTechniquesActionListTechniques ListTechniquesAction = "listTechniques"
)

func (e ListTechniquesAction) ToPointer() *ListTechniquesAction {
	return &e
}
func (e *ListTechniquesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = ListTechniquesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTechniquesAction: %v", v)
	}
}

type ListTechniquesData struct {
	Techniques []components.EditorTechnique `json:"techniques"`
}

func (o *ListTechniquesData) GetTechniques() []components.EditorTechnique {
	if o == nil {
		return []components.EditorTechnique{}
	}
	return o.Techniques
}

// ListTechniquesResponseBody - Techniques information
type ListTechniquesResponseBody struct {
	// Result of the request
	Result ListTechniquesResult `json:"result"`
	// The id of the action
	Action ListTechniquesAction `json:"action"`
	Data   ListTechniquesData   `json:"data"`
}

func (o *ListTechniquesResponseBody) GetResult() ListTechniquesResult {
	if o == nil {
		return ListTechniquesResult("")
	}
	return o.Result
}

func (o *ListTechniquesResponseBody) GetAction() ListTechniquesAction {
	if o == nil {
		return ListTechniquesAction("")
	}
	return o.Action
}

func (o *ListTechniquesResponseBody) GetData() ListTechniquesData {
	if o == nil {
		return ListTechniquesData{}
	}
	return o.Data
}

type ListTechniquesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *ListTechniquesResponseBody
}

func (o *ListTechniquesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTechniquesResponse) GetObject() *ListTechniquesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
