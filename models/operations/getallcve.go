package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetAllCveResult - Result of the request
type GetAllCveResult string

const (
	GetAllCveResultSuccess GetAllCveResult = "success"
	GetAllCveResultError   GetAllCveResult = "error"
)

func (e GetAllCveResult) ToPointer() *GetAllCveResult {
	return &e
}
func (e *GetAllCveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllCveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllCveResult: %v", v)
	}
}

// GetAllCveAction - The id of the action
type GetAllCveAction string

const (
	GetAllCveActionGetAllCve GetAllCveAction = "getAllCve"
)

func (e GetAllCveAction) ToPointer() *GetAllCveAction {
	return &e
}
func (e *GetAllCveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getAllCve":
		*e = GetAllCveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllCveAction: %v", v)
	}
}

type GetAllCveData struct {
	CVEs []components.CveDetails `json:"CVEs"`
}

func (o *GetAllCveData) GetCVEs() []components.CveDetails {
	if o == nil {
		return []components.CveDetails{}
	}
	return o.CVEs
}

// GetAllCveResponseBody - CVE details result
type GetAllCveResponseBody struct {
	// Result of the request
	Result GetAllCveResult `json:"result"`
	// The id of the action
	Action GetAllCveAction `json:"action"`
	Data   GetAllCveData   `json:"data"`
}

func (o *GetAllCveResponseBody) GetResult() GetAllCveResult {
	if o == nil {
		return GetAllCveResult("")
	}
	return o.Result
}

func (o *GetAllCveResponseBody) GetAction() GetAllCveAction {
	if o == nil {
		return GetAllCveAction("")
	}
	return o.Action
}

func (o *GetAllCveResponseBody) GetData() GetAllCveData {
	if o == nil {
		return GetAllCveData{}
	}
	return o.Data
}

type GetAllCveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// CVE details result
	Object *GetAllCveResponseBody
}

func (o *GetAllCveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllCveResponse) GetObject() *GetAllCveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
