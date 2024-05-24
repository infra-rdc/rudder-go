package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// TechniquesResult - Result of the request
type TechniquesResult string

const (
	TechniquesResultSuccess TechniquesResult = "success"
	TechniquesResultError   TechniquesResult = "error"
)

func (e TechniquesResult) ToPointer() *TechniquesResult {
	return &e
}
func (e *TechniquesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = TechniquesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TechniquesResult: %v", v)
	}
}

// TechniquesAction - The id of the action
type TechniquesAction string

const (
	TechniquesActionListTechniques TechniquesAction = "listTechniques"
)

func (e TechniquesAction) ToPointer() *TechniquesAction {
	return &e
}
func (e *TechniquesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = TechniquesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TechniquesAction: %v", v)
	}
}

type TechniquesData struct {
	Techniques []components.EditorTechnique `json:"techniques"`
}

func (o *TechniquesData) GetTechniques() []components.EditorTechnique {
	if o == nil {
		return []components.EditorTechnique{}
	}
	return o.Techniques
}

// TechniquesResponseBody - Techniques information
type TechniquesResponseBody struct {
	// Result of the request
	Result TechniquesResult `json:"result"`
	// The id of the action
	Action TechniquesAction `json:"action"`
	Data   TechniquesData   `json:"data"`
}

func (o *TechniquesResponseBody) GetResult() TechniquesResult {
	if o == nil {
		return TechniquesResult("")
	}
	return o.Result
}

func (o *TechniquesResponseBody) GetAction() TechniquesAction {
	if o == nil {
		return TechniquesAction("")
	}
	return o.Action
}

func (o *TechniquesResponseBody) GetData() TechniquesData {
	if o == nil {
		return TechniquesData{}
	}
	return o.Data
}

type TechniquesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *TechniquesResponseBody
}

func (o *TechniquesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TechniquesResponse) GetObject() *TechniquesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
