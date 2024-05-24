package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetTechniqueAllVersionIDRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
	// Technique version
	TechniqueVersion string `pathParam:"style=simple,explode=false,name=techniqueVersion"`
}

func (o *GetTechniqueAllVersionIDRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

func (o *GetTechniqueAllVersionIDRequest) GetTechniqueVersion() string {
	if o == nil {
		return ""
	}
	return o.TechniqueVersion
}

// GetTechniqueAllVersionIDResult - Result of the request
type GetTechniqueAllVersionIDResult string

const (
	GetTechniqueAllVersionIDResultSuccess GetTechniqueAllVersionIDResult = "success"
	GetTechniqueAllVersionIDResultError   GetTechniqueAllVersionIDResult = "error"
)

func (e GetTechniqueAllVersionIDResult) ToPointer() *GetTechniqueAllVersionIDResult {
	return &e
}
func (e *GetTechniqueAllVersionIDResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetTechniqueAllVersionIDResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniqueAllVersionIDResult: %v", v)
	}
}

// GetTechniqueAllVersionIDAction - The id of the action
type GetTechniqueAllVersionIDAction string

const (
	GetTechniqueAllVersionIDActionListTechniques GetTechniqueAllVersionIDAction = "listTechniques"
)

func (e GetTechniqueAllVersionIDAction) ToPointer() *GetTechniqueAllVersionIDAction {
	return &e
}
func (e *GetTechniqueAllVersionIDAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = GetTechniqueAllVersionIDAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniqueAllVersionIDAction: %v", v)
	}
}

type GetTechniqueAllVersionIDTechniques struct {
	JREditorTechnique components.EditorTechnique `json:"JREditorTechnique"`
}

func (o *GetTechniqueAllVersionIDTechniques) GetJREditorTechnique() components.EditorTechnique {
	if o == nil {
		return components.EditorTechnique{}
	}
	return o.JREditorTechnique
}

type GetTechniqueAllVersionIDData struct {
	Techniques []GetTechniqueAllVersionIDTechniques `json:"techniques"`
}

func (o *GetTechniqueAllVersionIDData) GetTechniques() []GetTechniqueAllVersionIDTechniques {
	if o == nil {
		return []GetTechniqueAllVersionIDTechniques{}
	}
	return o.Techniques
}

// GetTechniqueAllVersionIDResponseBody - Techniques information
type GetTechniqueAllVersionIDResponseBody struct {
	// Result of the request
	Result GetTechniqueAllVersionIDResult `json:"result"`
	// The id of the action
	Action GetTechniqueAllVersionIDAction `json:"action"`
	Data   GetTechniqueAllVersionIDData   `json:"data"`
}

func (o *GetTechniqueAllVersionIDResponseBody) GetResult() GetTechniqueAllVersionIDResult {
	if o == nil {
		return GetTechniqueAllVersionIDResult("")
	}
	return o.Result
}

func (o *GetTechniqueAllVersionIDResponseBody) GetAction() GetTechniqueAllVersionIDAction {
	if o == nil {
		return GetTechniqueAllVersionIDAction("")
	}
	return o.Action
}

func (o *GetTechniqueAllVersionIDResponseBody) GetData() GetTechniqueAllVersionIDData {
	if o == nil {
		return GetTechniqueAllVersionIDData{}
	}
	return o.Data
}

type GetTechniqueAllVersionIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Techniques information
	Object *GetTechniqueAllVersionIDResponseBody
}

func (o *GetTechniqueAllVersionIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTechniqueAllVersionIDResponse) GetObject() *GetTechniqueAllVersionIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
