package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type TechniqueRevisionsRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
	// Technique version
	TechniqueVersion string `pathParam:"style=simple,explode=false,name=techniqueVersion"`
}

func (o *TechniqueRevisionsRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

func (o *TechniqueRevisionsRequest) GetTechniqueVersion() string {
	if o == nil {
		return ""
	}
	return o.TechniqueVersion
}

// TechniqueRevisionsResult - Result of the request
type TechniqueRevisionsResult string

const (
	TechniqueRevisionsResultSuccess TechniqueRevisionsResult = "success"
	TechniqueRevisionsResultError   TechniqueRevisionsResult = "error"
)

func (e TechniqueRevisionsResult) ToPointer() *TechniqueRevisionsResult {
	return &e
}
func (e *TechniqueRevisionsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = TechniqueRevisionsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TechniqueRevisionsResult: %v", v)
	}
}

// TechniqueRevisionsAction - The id of the action
type TechniqueRevisionsAction string

const (
	TechniqueRevisionsActionListTechniques TechniqueRevisionsAction = "listTechniques"
)

func (e TechniqueRevisionsAction) ToPointer() *TechniqueRevisionsAction {
	return &e
}
func (e *TechniqueRevisionsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = TechniqueRevisionsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TechniqueRevisionsAction: %v", v)
	}
}

type TechniqueRevisionsData struct {
	Techniques []components.TechniquesRevisions `json:"techniques"`
}

func (o *TechniqueRevisionsData) GetTechniques() []components.TechniquesRevisions {
	if o == nil {
		return []components.TechniquesRevisions{}
	}
	return o.Techniques
}

// TechniqueRevisionsResponseBody - Versions information
type TechniqueRevisionsResponseBody struct {
	// Result of the request
	Result TechniqueRevisionsResult `json:"result"`
	// The id of the action
	Action TechniqueRevisionsAction `json:"action"`
	Data   TechniqueRevisionsData   `json:"data"`
}

func (o *TechniqueRevisionsResponseBody) GetResult() TechniqueRevisionsResult {
	if o == nil {
		return TechniqueRevisionsResult("")
	}
	return o.Result
}

func (o *TechniqueRevisionsResponseBody) GetAction() TechniqueRevisionsAction {
	if o == nil {
		return TechniqueRevisionsAction("")
	}
	return o.Action
}

func (o *TechniqueRevisionsResponseBody) GetData() TechniqueRevisionsData {
	if o == nil {
		return TechniqueRevisionsData{}
	}
	return o.Data
}

type TechniqueRevisionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Versions information
	Object *TechniqueRevisionsResponseBody
}

func (o *TechniqueRevisionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TechniqueRevisionsResponse) GetObject() *TechniqueRevisionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
