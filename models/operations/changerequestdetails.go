package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ChangeRequestDetailsRequest struct {
	ChangeRequestID int64 `pathParam:"style=simple,explode=false,name=changeRequestId"`
}

func (o *ChangeRequestDetailsRequest) GetChangeRequestID() int64 {
	if o == nil {
		return 0
	}
	return o.ChangeRequestID
}

// ChangeRequestDetailsResult - Result of the request
type ChangeRequestDetailsResult string

const (
	ChangeRequestDetailsResultSuccess ChangeRequestDetailsResult = "success"
	ChangeRequestDetailsResultError   ChangeRequestDetailsResult = "error"
)

func (e ChangeRequestDetailsResult) ToPointer() *ChangeRequestDetailsResult {
	return &e
}
func (e *ChangeRequestDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ChangeRequestDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChangeRequestDetailsResult: %v", v)
	}
}

// ChangeRequestDetailsAction - The id of the action
type ChangeRequestDetailsAction string

const (
	ChangeRequestDetailsActionChangeRequestDetails ChangeRequestDetailsAction = "changeRequestDetails"
)

func (e ChangeRequestDetailsAction) ToPointer() *ChangeRequestDetailsAction {
	return &e
}
func (e *ChangeRequestDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "changeRequestDetails":
		*e = ChangeRequestDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChangeRequestDetailsAction: %v", v)
	}
}

type ChangeRequestDetailsData struct {
	Rules []components.ChangeRequest `json:"rules"`
}

func (o *ChangeRequestDetailsData) GetRules() []components.ChangeRequest {
	if o == nil {
		return []components.ChangeRequest{}
	}
	return o.Rules
}

// ChangeRequestDetailsResponseBody - Change requests information
type ChangeRequestDetailsResponseBody struct {
	// Result of the request
	Result ChangeRequestDetailsResult `json:"result"`
	// The id of the action
	Action ChangeRequestDetailsAction `json:"action"`
	Data   ChangeRequestDetailsData   `json:"data"`
}

func (o *ChangeRequestDetailsResponseBody) GetResult() ChangeRequestDetailsResult {
	if o == nil {
		return ChangeRequestDetailsResult("")
	}
	return o.Result
}

func (o *ChangeRequestDetailsResponseBody) GetAction() ChangeRequestDetailsAction {
	if o == nil {
		return ChangeRequestDetailsAction("")
	}
	return o.Action
}

func (o *ChangeRequestDetailsResponseBody) GetData() ChangeRequestDetailsData {
	if o == nil {
		return ChangeRequestDetailsData{}
	}
	return o.Data
}

type ChangeRequestDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Change requests information
	Object *ChangeRequestDetailsResponseBody
}

func (o *ChangeRequestDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ChangeRequestDetailsResponse) GetObject() *ChangeRequestDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
