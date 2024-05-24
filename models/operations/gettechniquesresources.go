package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetTechniquesResourcesRequest struct {
	// Technique ID
	TechniqueID string `pathParam:"style=simple,explode=false,name=techniqueId"`
	// Technique version
	TechniqueVersion string `pathParam:"style=simple,explode=false,name=techniqueVersion"`
}

func (o *GetTechniquesResourcesRequest) GetTechniqueID() string {
	if o == nil {
		return ""
	}
	return o.TechniqueID
}

func (o *GetTechniquesResourcesRequest) GetTechniqueVersion() string {
	if o == nil {
		return ""
	}
	return o.TechniqueVersion
}

// GetTechniquesResourcesResult - Result of the request
type GetTechniquesResourcesResult string

const (
	GetTechniquesResourcesResultSuccess GetTechniquesResourcesResult = "success"
	GetTechniquesResourcesResultError   GetTechniquesResourcesResult = "error"
)

func (e GetTechniquesResourcesResult) ToPointer() *GetTechniquesResourcesResult {
	return &e
}
func (e *GetTechniquesResourcesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetTechniquesResourcesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniquesResourcesResult: %v", v)
	}
}

// GetTechniquesResourcesAction - The id of the action
type GetTechniquesResourcesAction string

const (
	GetTechniquesResourcesActionListTechniques GetTechniquesResourcesAction = "listTechniques"
)

func (e GetTechniquesResourcesAction) ToPointer() *GetTechniquesResourcesAction {
	return &e
}
func (e *GetTechniquesResourcesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "listTechniques":
		*e = GetTechniquesResourcesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTechniquesResourcesAction: %v", v)
	}
}

type GetTechniquesResourcesData struct {
	Resources []components.TechniquesResources `json:"resources"`
}

func (o *GetTechniquesResourcesData) GetResources() []components.TechniquesResources {
	if o == nil {
		return []components.TechniquesResources{}
	}
	return o.Resources
}

// GetTechniquesResourcesResponseBody - Resources information
type GetTechniquesResourcesResponseBody struct {
	// Result of the request
	Result GetTechniquesResourcesResult `json:"result"`
	// The id of the action
	Action GetTechniquesResourcesAction `json:"action"`
	Data   GetTechniquesResourcesData   `json:"data"`
}

func (o *GetTechniquesResourcesResponseBody) GetResult() GetTechniquesResourcesResult {
	if o == nil {
		return GetTechniquesResourcesResult("")
	}
	return o.Result
}

func (o *GetTechniquesResourcesResponseBody) GetAction() GetTechniquesResourcesAction {
	if o == nil {
		return GetTechniquesResourcesAction("")
	}
	return o.Action
}

func (o *GetTechniquesResourcesResponseBody) GetData() GetTechniquesResourcesData {
	if o == nil {
		return GetTechniquesResourcesData{}
	}
	return o.Data
}

type GetTechniquesResourcesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Resources information
	Object *GetTechniquesResourcesResponseBody
}

func (o *GetTechniquesResourcesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTechniquesResourcesResponse) GetObject() *GetTechniquesResourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
