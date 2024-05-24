package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetStatusResult - Result of the request
type GetStatusResult string

const (
	GetStatusResultSuccess GetStatusResult = "success"
	GetStatusResultError   GetStatusResult = "error"
)

func (e GetStatusResult) ToPointer() *GetStatusResult {
	return &e
}
func (e *GetStatusResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetStatusResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetStatusResult: %v", v)
	}
}

// GetStatusAction - The id of the action
type GetStatusAction string

const (
	GetStatusActionGetStatus GetStatusAction = "getStatus"
)

func (e GetStatusAction) ToPointer() *GetStatusAction {
	return &e
}
func (e *GetStatusAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getStatus":
		*e = GetStatusAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetStatusAction: %v", v)
	}
}

type Global string

const (
	GlobalOk Global = "OK"
)

func (e Global) ToPointer() *Global {
	return &e
}
func (e *Global) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OK":
		*e = Global(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Global: %v", v)
	}
}

// GetStatusData - Status of the service
type GetStatusData struct {
	Global *Global `json:"global,omitempty"`
}

func (o *GetStatusData) GetGlobal() *Global {
	if o == nil {
		return nil
	}
	return o.Global
}

// GetStatusResponseBody - Service status
type GetStatusResponseBody struct {
	// Result of the request
	Result GetStatusResult `json:"result"`
	// The id of the action
	Action GetStatusAction `json:"action"`
	// Status of the service
	Data GetStatusData `json:"data"`
}

func (o *GetStatusResponseBody) GetResult() GetStatusResult {
	if o == nil {
		return GetStatusResult("")
	}
	return o.Result
}

func (o *GetStatusResponseBody) GetAction() GetStatusAction {
	if o == nil {
		return GetStatusAction("")
	}
	return o.Action
}

func (o *GetStatusResponseBody) GetData() GetStatusData {
	if o == nil {
		return GetStatusData{}
	}
	return o.Data
}

type GetStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Service status
	Object *GetStatusResponseBody
}

func (o *GetStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetStatusResponse) GetObject() *GetStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
