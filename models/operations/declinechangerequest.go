package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeclineChangeRequestRequest struct {
	ChangeRequestID int64 `pathParam:"style=simple,explode=false,name=changeRequestId"`
}

func (o *DeclineChangeRequestRequest) GetChangeRequestID() int64 {
	if o == nil {
		return 0
	}
	return o.ChangeRequestID
}

// DeclineChangeRequestResult - Result of the request
type DeclineChangeRequestResult string

const (
	DeclineChangeRequestResultSuccess DeclineChangeRequestResult = "success"
	DeclineChangeRequestResultError   DeclineChangeRequestResult = "error"
)

func (e DeclineChangeRequestResult) ToPointer() *DeclineChangeRequestResult {
	return &e
}
func (e *DeclineChangeRequestResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeclineChangeRequestResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeclineChangeRequestResult: %v", v)
	}
}

// DeclineChangeRequestAction - The id of the action
type DeclineChangeRequestAction string

const (
	DeclineChangeRequestActionDeclineChangeRequest DeclineChangeRequestAction = "declineChangeRequest"
)

func (e DeclineChangeRequestAction) ToPointer() *DeclineChangeRequestAction {
	return &e
}
func (e *DeclineChangeRequestAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "declineChangeRequest":
		*e = DeclineChangeRequestAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeclineChangeRequestAction: %v", v)
	}
}

type DeclineChangeRequestData struct {
	Rules []components.ChangeRequest `json:"rules"`
}

func (o *DeclineChangeRequestData) GetRules() []components.ChangeRequest {
	if o == nil {
		return []components.ChangeRequest{}
	}
	return o.Rules
}

// DeclineChangeRequestResponseBody - Change requests information
type DeclineChangeRequestResponseBody struct {
	// Result of the request
	Result DeclineChangeRequestResult `json:"result"`
	// The id of the action
	Action DeclineChangeRequestAction `json:"action"`
	Data   DeclineChangeRequestData   `json:"data"`
}

func (o *DeclineChangeRequestResponseBody) GetResult() DeclineChangeRequestResult {
	if o == nil {
		return DeclineChangeRequestResult("")
	}
	return o.Result
}

func (o *DeclineChangeRequestResponseBody) GetAction() DeclineChangeRequestAction {
	if o == nil {
		return DeclineChangeRequestAction("")
	}
	return o.Action
}

func (o *DeclineChangeRequestResponseBody) GetData() DeclineChangeRequestData {
	if o == nil {
		return DeclineChangeRequestData{}
	}
	return o.Data
}

type DeclineChangeRequestResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Change requests information
	Object *DeclineChangeRequestResponseBody
}

func (o *DeclineChangeRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeclineChangeRequestResponse) GetObject() *DeclineChangeRequestResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
