package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
	"github.com/infra-rdc/rudder-go/types"
)

// GetSystemInfoResult - Result of the request
type GetSystemInfoResult string

const (
	GetSystemInfoResultSuccess GetSystemInfoResult = "success"
	GetSystemInfoResultError   GetSystemInfoResult = "error"
)

func (e GetSystemInfoResult) ToPointer() *GetSystemInfoResult {
	return &e
}
func (e *GetSystemInfoResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetSystemInfoResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSystemInfoResult: %v", v)
	}
}

// GetSystemInfoAction - The id of the action
type GetSystemInfoAction string

const (
	GetSystemInfoActionGetSystemInfo GetSystemInfoAction = "getSystemInfo"
)

func (e GetSystemInfoAction) ToPointer() *GetSystemInfoAction {
	return &e
}
func (e *GetSystemInfoAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getSystemInfo":
		*e = GetSystemInfoAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSystemInfoAction: %v", v)
	}
}

type Rudder struct {
	MajorVersion string     `json:"major-version"`
	FullVersion  string     `json:"full-version"`
	BuildTime    types.Date `json:"build-time"`
}

func (r Rudder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Rudder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Rudder) GetMajorVersion() string {
	if o == nil {
		return ""
	}
	return o.MajorVersion
}

func (o *Rudder) GetFullVersion() string {
	if o == nil {
		return ""
	}
	return o.FullVersion
}

func (o *Rudder) GetBuildTime() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.BuildTime
}

// GetSystemInfoData - Information about the service
type GetSystemInfoData struct {
	Rudder Rudder `json:"rudder"`
}

func (o *GetSystemInfoData) GetRudder() Rudder {
	if o == nil {
		return Rudder{}
	}
	return o.Rudder
}

// GetSystemInfoResponseBody - Service information
type GetSystemInfoResponseBody struct {
	// Result of the request
	Result GetSystemInfoResult `json:"result"`
	// The id of the action
	Action GetSystemInfoAction `json:"action"`
	// Information about the service
	Data GetSystemInfoData `json:"data"`
}

func (o *GetSystemInfoResponseBody) GetResult() GetSystemInfoResult {
	if o == nil {
		return GetSystemInfoResult("")
	}
	return o.Result
}

func (o *GetSystemInfoResponseBody) GetAction() GetSystemInfoAction {
	if o == nil {
		return GetSystemInfoAction("")
	}
	return o.Action
}

func (o *GetSystemInfoResponseBody) GetData() GetSystemInfoData {
	if o == nil {
		return GetSystemInfoData{}
	}
	return o.Data
}

type GetSystemInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Service information
	Object *GetSystemInfoResponseBody
}

func (o *GetSystemInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSystemInfoResponse) GetObject() *GetSystemInfoResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
