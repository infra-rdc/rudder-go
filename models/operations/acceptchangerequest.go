package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// Status - New status of the change request
type Status string

const (
	StatusPendingDeployment Status = "pending deployment"
	StatusDeployed          Status = "deployed"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending deployment":
		fallthrough
	case "deployed":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type AcceptChangeRequestRequestBody struct {
	// New status of the change request
	Status *Status `json:"status,omitempty"`
}

func (o *AcceptChangeRequestRequestBody) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type AcceptChangeRequestRequest struct {
	ChangeRequestID int64                          `pathParam:"style=simple,explode=false,name=changeRequestId"`
	RequestBody     AcceptChangeRequestRequestBody `request:"mediaType=application/json"`
}

func (o *AcceptChangeRequestRequest) GetChangeRequestID() int64 {
	if o == nil {
		return 0
	}
	return o.ChangeRequestID
}

func (o *AcceptChangeRequestRequest) GetRequestBody() AcceptChangeRequestRequestBody {
	if o == nil {
		return AcceptChangeRequestRequestBody{}
	}
	return o.RequestBody
}

// AcceptChangeRequestResult - Result of the request
type AcceptChangeRequestResult string

const (
	AcceptChangeRequestResultSuccess AcceptChangeRequestResult = "success"
	AcceptChangeRequestResultError   AcceptChangeRequestResult = "error"
)

func (e AcceptChangeRequestResult) ToPointer() *AcceptChangeRequestResult {
	return &e
}
func (e *AcceptChangeRequestResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = AcceptChangeRequestResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AcceptChangeRequestResult: %v", v)
	}
}

// AcceptChangeRequestAction - The id of the action
type AcceptChangeRequestAction string

const (
	AcceptChangeRequestActionAcceptChangeRequest AcceptChangeRequestAction = "acceptChangeRequest"
)

func (e AcceptChangeRequestAction) ToPointer() *AcceptChangeRequestAction {
	return &e
}
func (e *AcceptChangeRequestAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "acceptChangeRequest":
		*e = AcceptChangeRequestAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AcceptChangeRequestAction: %v", v)
	}
}

type AcceptChangeRequestData struct {
	Rules []components.ChangeRequest `json:"rules"`
}

func (o *AcceptChangeRequestData) GetRules() []components.ChangeRequest {
	if o == nil {
		return []components.ChangeRequest{}
	}
	return o.Rules
}

// AcceptChangeRequestResponseBody - Change requests information
type AcceptChangeRequestResponseBody struct {
	// Result of the request
	Result AcceptChangeRequestResult `json:"result"`
	// The id of the action
	Action AcceptChangeRequestAction `json:"action"`
	Data   AcceptChangeRequestData   `json:"data"`
}

func (o *AcceptChangeRequestResponseBody) GetResult() AcceptChangeRequestResult {
	if o == nil {
		return AcceptChangeRequestResult("")
	}
	return o.Result
}

func (o *AcceptChangeRequestResponseBody) GetAction() AcceptChangeRequestAction {
	if o == nil {
		return AcceptChangeRequestAction("")
	}
	return o.Action
}

func (o *AcceptChangeRequestResponseBody) GetData() AcceptChangeRequestData {
	if o == nil {
		return AcceptChangeRequestData{}
	}
	return o.Data
}

type AcceptChangeRequestResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Change requests information
	Object *AcceptChangeRequestResponseBody
}

func (o *AcceptChangeRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AcceptChangeRequestResponse) GetObject() *AcceptChangeRequestResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
