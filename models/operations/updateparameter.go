package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateParameterRequest struct {
	// Id of the parameter to manage
	ParameterID string `pathParam:"style=simple,explode=false,name=parameterId"`
}

func (o *UpdateParameterRequest) GetParameterID() string {
	if o == nil {
		return ""
	}
	return o.ParameterID
}

// UpdateParameterResult - Result of the request
type UpdateParameterResult string

const (
	UpdateParameterResultSuccess UpdateParameterResult = "success"
	UpdateParameterResultError   UpdateParameterResult = "error"
)

func (e UpdateParameterResult) ToPointer() *UpdateParameterResult {
	return &e
}
func (e *UpdateParameterResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateParameterResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateParameterResult: %v", v)
	}
}

// UpdateParameterAction - The id of the action
type UpdateParameterAction string

const (
	UpdateParameterActionUpdateParameter UpdateParameterAction = "updateParameter"
)

func (e UpdateParameterAction) ToPointer() *UpdateParameterAction {
	return &e
}
func (e *UpdateParameterAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateParameter":
		*e = UpdateParameterAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateParameterAction: %v", v)
	}
}

// UpdateParameterData - Parameters
type UpdateParameterData struct {
	Parameters []components.Parameter `json:"parameters"`
}

func (o *UpdateParameterData) GetParameters() []components.Parameter {
	if o == nil {
		return []components.Parameter{}
	}
	return o.Parameters
}

// UpdateParameterResponseBody - Settings
type UpdateParameterResponseBody struct {
	// Id of the global property
	ID string `json:"id"`
	// Result of the request
	Result UpdateParameterResult `json:"result"`
	// The id of the action
	Action UpdateParameterAction `json:"action"`
	// Parameters
	Data UpdateParameterData `json:"data"`
}

func (o *UpdateParameterResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateParameterResponseBody) GetResult() UpdateParameterResult {
	if o == nil {
		return UpdateParameterResult("")
	}
	return o.Result
}

func (o *UpdateParameterResponseBody) GetAction() UpdateParameterAction {
	if o == nil {
		return UpdateParameterAction("")
	}
	return o.Action
}

func (o *UpdateParameterResponseBody) GetData() UpdateParameterData {
	if o == nil {
		return UpdateParameterData{}
	}
	return o.Data
}

type UpdateParameterResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *UpdateParameterResponseBody
}

func (o *UpdateParameterResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateParameterResponse) GetObject() *UpdateParameterResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
