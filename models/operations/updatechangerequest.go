package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateChangeRequestRequestBody struct {
	// Change request name
	Name *string `json:"name,omitempty"`
	// Change request description
	Description *string `json:"description,omitempty"`
}

func (o *UpdateChangeRequestRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateChangeRequestRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type UpdateChangeRequestRequest struct {
	ChangeRequestID int64                          `pathParam:"style=simple,explode=false,name=changeRequestId"`
	RequestBody     UpdateChangeRequestRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateChangeRequestRequest) GetChangeRequestID() int64 {
	if o == nil {
		return 0
	}
	return o.ChangeRequestID
}

func (o *UpdateChangeRequestRequest) GetRequestBody() UpdateChangeRequestRequestBody {
	if o == nil {
		return UpdateChangeRequestRequestBody{}
	}
	return o.RequestBody
}

// UpdateChangeRequestResult - Result of the request
type UpdateChangeRequestResult string

const (
	UpdateChangeRequestResultSuccess UpdateChangeRequestResult = "success"
	UpdateChangeRequestResultError   UpdateChangeRequestResult = "error"
)

func (e UpdateChangeRequestResult) ToPointer() *UpdateChangeRequestResult {
	return &e
}
func (e *UpdateChangeRequestResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateChangeRequestResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateChangeRequestResult: %v", v)
	}
}

// UpdateChangeRequestAction - The id of the action
type UpdateChangeRequestAction string

const (
	UpdateChangeRequestActionUpdateChangeRequest UpdateChangeRequestAction = "updateChangeRequest"
)

func (e UpdateChangeRequestAction) ToPointer() *UpdateChangeRequestAction {
	return &e
}
func (e *UpdateChangeRequestAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateChangeRequest":
		*e = UpdateChangeRequestAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateChangeRequestAction: %v", v)
	}
}

type UpdateChangeRequestData struct {
	Rules []components.ChangeRequest `json:"rules"`
}

func (o *UpdateChangeRequestData) GetRules() []components.ChangeRequest {
	if o == nil {
		return []components.ChangeRequest{}
	}
	return o.Rules
}

// UpdateChangeRequestResponseBody - Change requests information
type UpdateChangeRequestResponseBody struct {
	// Result of the request
	Result UpdateChangeRequestResult `json:"result"`
	// The id of the action
	Action UpdateChangeRequestAction `json:"action"`
	Data   UpdateChangeRequestData   `json:"data"`
}

func (o *UpdateChangeRequestResponseBody) GetResult() UpdateChangeRequestResult {
	if o == nil {
		return UpdateChangeRequestResult("")
	}
	return o.Result
}

func (o *UpdateChangeRequestResponseBody) GetAction() UpdateChangeRequestAction {
	if o == nil {
		return UpdateChangeRequestAction("")
	}
	return o.Action
}

func (o *UpdateChangeRequestResponseBody) GetData() UpdateChangeRequestData {
	if o == nil {
		return UpdateChangeRequestData{}
	}
	return o.Data
}

type UpdateChangeRequestResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Change requests information
	Object *UpdateChangeRequestResponseBody
}

func (o *UpdateChangeRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateChangeRequestResponse) GetObject() *UpdateChangeRequestResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
