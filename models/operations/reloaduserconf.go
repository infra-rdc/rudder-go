package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadUserConfResult - Result of the request
type ReloadUserConfResult string

const (
	ReloadUserConfResultSuccess ReloadUserConfResult = "success"
	ReloadUserConfResultError   ReloadUserConfResult = "error"
)

func (e ReloadUserConfResult) ToPointer() *ReloadUserConfResult {
	return &e
}
func (e *ReloadUserConfResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadUserConfResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadUserConfResult: %v", v)
	}
}

// ReloadUserConfAction - The id of the action
type ReloadUserConfAction string

const (
	ReloadUserConfActionReloadUserConf ReloadUserConfAction = "reloadUserConf"
)

func (e ReloadUserConfAction) ToPointer() *ReloadUserConfAction {
	return &e
}
func (e *ReloadUserConfAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reloadUserConf":
		*e = ReloadUserConfAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadUserConfAction: %v", v)
	}
}

type Reload struct {
	Status string `json:"status"`
}

func (o *Reload) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type ReloadUserConfData struct {
	Reload Reload `json:"reload"`
}

func (o *ReloadUserConfData) GetReload() Reload {
	if o == nil {
		return Reload{}
	}
	return o.Reload
}

// ReloadUserConfResponseBody - Reload information
type ReloadUserConfResponseBody struct {
	// Result of the request
	Result ReloadUserConfResult `json:"result"`
	// The id of the action
	Action ReloadUserConfAction `json:"action"`
	Data   ReloadUserConfData   `json:"data"`
}

func (o *ReloadUserConfResponseBody) GetResult() ReloadUserConfResult {
	if o == nil {
		return ReloadUserConfResult("")
	}
	return o.Result
}

func (o *ReloadUserConfResponseBody) GetAction() ReloadUserConfAction {
	if o == nil {
		return ReloadUserConfAction("")
	}
	return o.Action
}

func (o *ReloadUserConfResponseBody) GetData() ReloadUserConfData {
	if o == nil {
		return ReloadUserConfData{}
	}
	return o.Data
}

type ReloadUserConfResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Reload information
	Object *ReloadUserConfResponseBody
}

func (o *ReloadUserConfResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadUserConfResponse) GetObject() *ReloadUserConfResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
