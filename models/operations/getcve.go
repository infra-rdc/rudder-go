package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetCveRequest struct {
	// Id of the CVE
	CveID string `pathParam:"style=simple,explode=false,name=cveId"`
}

func (o *GetCveRequest) GetCveID() string {
	if o == nil {
		return ""
	}
	return o.CveID
}

// GetCveResult - Result of the request
type GetCveResult string

const (
	GetCveResultSuccess GetCveResult = "success"
	GetCveResultError   GetCveResult = "error"
)

func (e GetCveResult) ToPointer() *GetCveResult {
	return &e
}
func (e *GetCveResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCveResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCveResult: %v", v)
	}
}

// GetCveAction - The id of the action
type GetCveAction string

const (
	GetCveActionGetCve GetCveAction = "getCve"
)

func (e GetCveAction) ToPointer() *GetCveAction {
	return &e
}
func (e *GetCveAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCve":
		*e = GetCveAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCveAction: %v", v)
	}
}

type GetCveData struct {
	CVEs []components.CveDetails `json:"CVEs"`
}

func (o *GetCveData) GetCVEs() []components.CveDetails {
	if o == nil {
		return []components.CveDetails{}
	}
	return o.CVEs
}

// GetCveResponseBody - CVE details result
type GetCveResponseBody struct {
	// Result of the request
	Result GetCveResult `json:"result"`
	// The id of the action
	Action GetCveAction `json:"action"`
	Data   GetCveData   `json:"data"`
}

func (o *GetCveResponseBody) GetResult() GetCveResult {
	if o == nil {
		return GetCveResult("")
	}
	return o.Result
}

func (o *GetCveResponseBody) GetAction() GetCveAction {
	if o == nil {
		return GetCveAction("")
	}
	return o.Action
}

func (o *GetCveResponseBody) GetData() GetCveData {
	if o == nil {
		return GetCveData{}
	}
	return o.Data
}

type GetCveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// CVE details result
	Object *GetCveResponseBody
}

func (o *GetCveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCveResponse) GetObject() *GetCveResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
