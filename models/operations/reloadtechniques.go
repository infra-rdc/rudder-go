package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadTechniquesResult - Result of the request
type ReloadTechniquesResult string

const (
	ReloadTechniquesResultSuccess ReloadTechniquesResult = "success"
	ReloadTechniquesResultError   ReloadTechniquesResult = "error"
)

func (e ReloadTechniquesResult) ToPointer() *ReloadTechniquesResult {
	return &e
}
func (e *ReloadTechniquesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadTechniquesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadTechniquesResult: %v", v)
	}
}

// ReloadTechniquesAction - The id of the action
type ReloadTechniquesAction string

const (
	ReloadTechniquesActionReloadTechniques ReloadTechniquesAction = "reloadTechniques"
)

func (e ReloadTechniquesAction) ToPointer() *ReloadTechniquesAction {
	return &e
}
func (e *ReloadTechniquesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reloadTechniques":
		*e = ReloadTechniquesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadTechniquesAction: %v", v)
	}
}

type ReloadTechniquesTechniques string

const (
	ReloadTechniquesTechniquesStarted ReloadTechniquesTechniques = "Started"
)

func (e ReloadTechniquesTechniques) ToPointer() *ReloadTechniquesTechniques {
	return &e
}
func (e *ReloadTechniquesTechniques) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Started":
		*e = ReloadTechniquesTechniques(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadTechniquesTechniques: %v", v)
	}
}

type ReloadTechniquesData struct {
	Techniques ReloadTechniquesTechniques `json:"techniques"`
}

func (o *ReloadTechniquesData) GetTechniques() ReloadTechniquesTechniques {
	if o == nil {
		return ReloadTechniquesTechniques("")
	}
	return o.Techniques
}

// ReloadTechniquesResponseBody - Service reload
type ReloadTechniquesResponseBody struct {
	// Result of the request
	Result ReloadTechniquesResult `json:"result"`
	// The id of the action
	Action ReloadTechniquesAction `json:"action"`
	Data   ReloadTechniquesData   `json:"data"`
}

func (o *ReloadTechniquesResponseBody) GetResult() ReloadTechniquesResult {
	if o == nil {
		return ReloadTechniquesResult("")
	}
	return o.Result
}

func (o *ReloadTechniquesResponseBody) GetAction() ReloadTechniquesAction {
	if o == nil {
		return ReloadTechniquesAction("")
	}
	return o.Action
}

func (o *ReloadTechniquesResponseBody) GetData() ReloadTechniquesData {
	if o == nil {
		return ReloadTechniquesData{}
	}
	return o.Data
}

type ReloadTechniquesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Service reload
	Object *ReloadTechniquesResponseBody
}

func (o *ReloadTechniquesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadTechniquesResponse) GetObject() *ReloadTechniquesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
