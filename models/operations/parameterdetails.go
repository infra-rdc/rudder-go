package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ParameterDetailsRequest struct {
	// Id of the parameter to manage
	ParameterID string `pathParam:"style=simple,explode=false,name=parameterId"`
}

func (o *ParameterDetailsRequest) GetParameterID() string {
	if o == nil {
		return ""
	}
	return o.ParameterID
}

// ParameterDetailsResult - Result of the request
type ParameterDetailsResult string

const (
	ParameterDetailsResultSuccess ParameterDetailsResult = "success"
	ParameterDetailsResultError   ParameterDetailsResult = "error"
)

func (e ParameterDetailsResult) ToPointer() *ParameterDetailsResult {
	return &e
}
func (e *ParameterDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ParameterDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterDetailsResult: %v", v)
	}
}

// ParameterDetailsAction - The id of the action
type ParameterDetailsAction string

const (
	ParameterDetailsActionParameterDetails ParameterDetailsAction = "parameterDetails"
)

func (e ParameterDetailsAction) ToPointer() *ParameterDetailsAction {
	return &e
}
func (e *ParameterDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "parameterDetails":
		*e = ParameterDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ParameterDetailsAction: %v", v)
	}
}

// ParameterDetailsData - Parameters
type ParameterDetailsData struct {
	Parameters []components.Parameter `json:"parameters"`
}

func (o *ParameterDetailsData) GetParameters() []components.Parameter {
	if o == nil {
		return []components.Parameter{}
	}
	return o.Parameters
}

// ParameterDetailsResponseBody - Settings
type ParameterDetailsResponseBody struct {
	// Id of the property
	ID string `json:"id"`
	// Result of the request
	Result ParameterDetailsResult `json:"result"`
	// The id of the action
	Action ParameterDetailsAction `json:"action"`
	// Parameters
	Data ParameterDetailsData `json:"data"`
}

func (o *ParameterDetailsResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ParameterDetailsResponseBody) GetResult() ParameterDetailsResult {
	if o == nil {
		return ParameterDetailsResult("")
	}
	return o.Result
}

func (o *ParameterDetailsResponseBody) GetAction() ParameterDetailsAction {
	if o == nil {
		return ParameterDetailsAction("")
	}
	return o.Action
}

func (o *ParameterDetailsResponseBody) GetData() ParameterDetailsData {
	if o == nil {
		return ParameterDetailsData{}
	}
	return o.Data
}

type ParameterDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *ParameterDetailsResponseBody
}

func (o *ParameterDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ParameterDetailsResponse) GetObject() *ParameterDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
