package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteTechniqueRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
	// Technique version
	TechniqueVersion string `pathParam:"style=simple,explode=false,name=techniqueVersion"`
}

func (o *DeleteTechniqueRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

func (o *DeleteTechniqueRequest) GetTechniqueVersion() string {
	if o == nil {
		return ""
	}
	return o.TechniqueVersion
}

// DeleteTechniqueResult - Result of the request
type DeleteTechniqueResult string

const (
	DeleteTechniqueResultSuccess DeleteTechniqueResult = "success"
	DeleteTechniqueResultError   DeleteTechniqueResult = "error"
)

func (e DeleteTechniqueResult) ToPointer() *DeleteTechniqueResult {
	return &e
}
func (e *DeleteTechniqueResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteTechniqueResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteTechniqueResult: %v", v)
	}
}

// DeleteTechniqueAction - The id of the action
type DeleteTechniqueAction string

const (
	DeleteTechniqueActionListTechniques DeleteTechniqueAction = "listTechniques"
)

func (e DeleteTechniqueAction) ToPointer() *DeleteTechniqueAction {
	return &e
}
func (e *DeleteTechniqueAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = DeleteTechniqueAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteTechniqueAction: %v", v)
	}
}

type DeleteTechniqueTechniques struct {
	// Deleted technique ID
	ID string `json:"id"`
	// Deleted technique version
	Version string `json:"version"`
}

func (o *DeleteTechniqueTechniques) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteTechniqueTechniques) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

type DeleteTechniqueData struct {
	Techniques *DeleteTechniqueTechniques `json:"techniques,omitempty"`
}

func (o *DeleteTechniqueData) GetTechniques() *DeleteTechniqueTechniques {
	if o == nil {
		return nil
	}
	return o.Techniques
}

// DeleteTechniqueResponseBody - Deletion information
type DeleteTechniqueResponseBody struct {
	// Result of the request
	Result DeleteTechniqueResult `json:"result"`
	// The id of the action
	Action DeleteTechniqueAction `json:"action"`
	Data   DeleteTechniqueData   `json:"data"`
}

func (o *DeleteTechniqueResponseBody) GetResult() DeleteTechniqueResult {
	if o == nil {
		return DeleteTechniqueResult("")
	}
	return o.Result
}

func (o *DeleteTechniqueResponseBody) GetAction() DeleteTechniqueAction {
	if o == nil {
		return DeleteTechniqueAction("")
	}
	return o.Action
}

func (o *DeleteTechniqueResponseBody) GetData() DeleteTechniqueData {
	if o == nil {
		return DeleteTechniqueData{}
	}
	return o.Data
}

type DeleteTechniqueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deletion information
	Object *DeleteTechniqueResponseBody
}

func (o *DeleteTechniqueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteTechniqueResponse) GetObject() *DeleteTechniqueResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
