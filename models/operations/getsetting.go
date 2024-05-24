package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetSettingRequest struct {
	// Id of the setting to set
	SettingID string `pathParam:"style=simple,explode=false,name=settingId"`
}

func (o *GetSettingRequest) GetSettingID() string {
	if o == nil {
		return ""
	}
	return o.SettingID
}

// GetSettingResult - Result of the request
type GetSettingResult string

const (
	GetSettingResultSuccess GetSettingResult = "success"
	GetSettingResultError   GetSettingResult = "error"
)

func (e GetSettingResult) ToPointer() *GetSettingResult {
	return &e
}
func (e *GetSettingResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetSettingResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSettingResult: %v", v)
	}
}

// GetSettingAction - The id of the action
type GetSettingAction string

const (
	GetSettingActionGetSetting GetSettingAction = "getSetting"
)

func (e GetSettingAction) ToPointer() *GetSettingAction {
	return &e
}
func (e *GetSettingAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getSetting":
		*e = GetSettingAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSettingAction: %v", v)
	}
}

// GetSettingData - Information about the setting
type GetSettingData struct {
	// Id and value of the property
	SettingID *string `json:"settingId,omitempty"`
}

func (o *GetSettingData) GetSettingID() *string {
	if o == nil {
		return nil
	}
	return o.SettingID
}

// GetSettingResponseBody - Settings
type GetSettingResponseBody struct {
	// Id of the setting
	ID string `json:"id"`
	// Result of the request
	Result GetSettingResult `json:"result"`
	// The id of the action
	Action GetSettingAction `json:"action"`
	// Information about the setting
	Data GetSettingData `json:"data"`
}

func (o *GetSettingResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetSettingResponseBody) GetResult() GetSettingResult {
	if o == nil {
		return GetSettingResult("")
	}
	return o.Result
}

func (o *GetSettingResponseBody) GetAction() GetSettingAction {
	if o == nil {
		return GetSettingAction("")
	}
	return o.Action
}

func (o *GetSettingResponseBody) GetData() GetSettingData {
	if o == nil {
		return GetSettingData{}
	}
	return o.Data
}

type GetSettingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *GetSettingResponseBody
}

func (o *GetSettingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSettingResponse) GetObject() *GetSettingResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
