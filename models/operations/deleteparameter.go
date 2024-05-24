package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteParameterRequest struct {
	// Id of the parameter to manage
	ParameterID string `pathParam:"style=simple,explode=false,name=parameterId"`
}

func (o *DeleteParameterRequest) GetParameterID() string {
	if o == nil {
		return ""
	}
	return o.ParameterID
}

// DeleteParameterResult - Result of the request
type DeleteParameterResult string

const (
	DeleteParameterResultSuccess DeleteParameterResult = "success"
	DeleteParameterResultError   DeleteParameterResult = "error"
)

func (e DeleteParameterResult) ToPointer() *DeleteParameterResult {
	return &e
}
func (e *DeleteParameterResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteParameterResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteParameterResult: %v", v)
	}
}

// DeleteParameterAction - The id of the action
type DeleteParameterAction string

const (
	DeleteParameterActionDeleteParameter DeleteParameterAction = "deleteParameter"
)

func (e DeleteParameterAction) ToPointer() *DeleteParameterAction {
	return &e
}
func (e *DeleteParameterAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteParameter":
		*e = DeleteParameterAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteParameterAction: %v", v)
	}
}

// DeleteParameterData - Parameters
type DeleteParameterData struct {
	Parameters []components.Parameter `json:"parameters"`
}

func (o *DeleteParameterData) GetParameters() []components.Parameter {
	if o == nil {
		return []components.Parameter{}
	}
	return o.Parameters
}

// DeleteParameterResponseBody - Settings
type DeleteParameterResponseBody struct {
	// Id of the global property
	ID string `json:"id"`
	// Result of the request
	Result DeleteParameterResult `json:"result"`
	// The id of the action
	Action DeleteParameterAction `json:"action"`
	// Parameters
	Data DeleteParameterData `json:"data"`
}

func (o *DeleteParameterResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteParameterResponseBody) GetResult() DeleteParameterResult {
	if o == nil {
		return DeleteParameterResult("")
	}
	return o.Result
}

func (o *DeleteParameterResponseBody) GetAction() DeleteParameterAction {
	if o == nil {
		return DeleteParameterAction("")
	}
	return o.Action
}

func (o *DeleteParameterResponseBody) GetData() DeleteParameterData {
	if o == nil {
		return DeleteParameterData{}
	}
	return o.Data
}

type DeleteParameterResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *DeleteParameterResponseBody
}

func (o *DeleteParameterResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteParameterResponse) GetObject() *DeleteParameterResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
