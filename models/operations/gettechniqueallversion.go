package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetTechniqueAllVersionRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
}

func (o *GetTechniqueAllVersionRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

// GetTechniqueAllVersionResult - Result of the request
type GetTechniqueAllVersionResult string

const (
	GetTechniqueAllVersionResultSuccess GetTechniqueAllVersionResult = "success"
	GetTechniqueAllVersionResultError   GetTechniqueAllVersionResult = "error"
)

func (e GetTechniqueAllVersionResult) ToPointer() *GetTechniqueAllVersionResult {
	return &e
}
func (e *GetTechniqueAllVersionResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetTechniqueAllVersionResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniqueAllVersionResult: %v", v)
	}
}

// GetTechniqueAllVersionAction - The id of the action
type GetTechniqueAllVersionAction string

const (
	GetTechniqueAllVersionActionListTechniques GetTechniqueAllVersionAction = "listTechniques"
)

func (e GetTechniqueAllVersionAction) ToPointer() *GetTechniqueAllVersionAction {
	return &e
}
func (e *GetTechniqueAllVersionAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = GetTechniqueAllVersionAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniqueAllVersionAction: %v", v)
	}
}

type GetTechniqueAllVersionTechniques struct {
	JREditorTechnique components.EditorTechnique `json:"JREditorTechnique"`
}

func (o *GetTechniqueAllVersionTechniques) GetJREditorTechnique() components.EditorTechnique {
	if o == nil {
		return components.EditorTechnique{}
	}
	return o.JREditorTechnique
}

type GetTechniqueAllVersionData struct {
	Techniques []GetTechniqueAllVersionTechniques `json:"techniques"`
}

func (o *GetTechniqueAllVersionData) GetTechniques() []GetTechniqueAllVersionTechniques {
	if o == nil {
		return []GetTechniqueAllVersionTechniques{}
	}
	return o.Techniques
}

// GetTechniqueAllVersionResponseBody - Techniques information
type GetTechniqueAllVersionResponseBody struct {
	// Result of the request
	Result GetTechniqueAllVersionResult `json:"result"`
	// The id of the action
	Action GetTechniqueAllVersionAction `json:"action"`
	Data   GetTechniqueAllVersionData   `json:"data"`
}

func (o *GetTechniqueAllVersionResponseBody) GetResult() GetTechniqueAllVersionResult {
	if o == nil {
		return GetTechniqueAllVersionResult("")
	}
	return o.Result
}

func (o *GetTechniqueAllVersionResponseBody) GetAction() GetTechniqueAllVersionAction {
	if o == nil {
		return GetTechniqueAllVersionAction("")
	}
	return o.Action
}

func (o *GetTechniqueAllVersionResponseBody) GetData() GetTechniqueAllVersionData {
	if o == nil {
		return GetTechniqueAllVersionData{}
	}
	return o.Data
}

type GetTechniqueAllVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *GetTechniqueAllVersionResponseBody
}

func (o *GetTechniqueAllVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTechniqueAllVersionResponse) GetObject() *GetTechniqueAllVersionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
