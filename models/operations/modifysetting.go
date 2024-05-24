package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ModifySettingRequestBody struct {
	// New value of the setting
	Value *string `json:"value,omitempty"`
}

func (o *ModifySettingRequestBody) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type ModifySettingRequest struct {
	// Id of the setting to set
	SettingID   string                   `pathParam:"style=simple,explode=false,name=settingId"`
	RequestBody ModifySettingRequestBody `request:"mediaType=application/json"`
}

func (o *ModifySettingRequest) GetSettingID() string {
	if o == nil {
		return ""
	}
	return o.SettingID
}

func (o *ModifySettingRequest) GetRequestBody() ModifySettingRequestBody {
	if o == nil {
		return ModifySettingRequestBody{}
	}
	return o.RequestBody
}

// ModifySettingResult - Result of the request
type ModifySettingResult string

const (
	ModifySettingResultSuccess ModifySettingResult = "success"
	ModifySettingResultError   ModifySettingResult = "error"
)

func (e ModifySettingResult) ToPointer() *ModifySettingResult {
	return &e
}
func (e *ModifySettingResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ModifySettingResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ModifySettingResult: %v", v)
	}
}

// ModifySettingAction - The id of the action
type ModifySettingAction string

const (
	ModifySettingActionModifySetting ModifySettingAction = "modifySetting"
)

func (e ModifySettingAction) ToPointer() *ModifySettingAction {
	return &e
}
func (e *ModifySettingAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "modifySetting":
		*e = ModifySettingAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ModifySettingAction: %v", v)
	}
}

// ModifySettingData - Information about the setting
type ModifySettingData struct {
	// Id and value of the property
	SettingID *string `json:"settingId,omitempty"`
}

func (o *ModifySettingData) GetSettingID() *string {
	if o == nil {
		return nil
	}
	return o.SettingID
}

// ModifySettingResponseBody - Settings
type ModifySettingResponseBody struct {
	// Id of the setting
	ID string `json:"id"`
	// Result of the request
	Result ModifySettingResult `json:"result"`
	// The id of the action
	Action ModifySettingAction `json:"action"`
	// Information about the setting
	Data ModifySettingData `json:"data"`
}

func (o *ModifySettingResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ModifySettingResponseBody) GetResult() ModifySettingResult {
	if o == nil {
		return ModifySettingResult("")
	}
	return o.Result
}

func (o *ModifySettingResponseBody) GetAction() ModifySettingAction {
	if o == nil {
		return ModifySettingAction("")
	}
	return o.Action
}

func (o *ModifySettingResponseBody) GetData() ModifySettingData {
	if o == nil {
		return ModifySettingData{}
	}
	return o.Data
}

type ModifySettingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *ModifySettingResponseBody
}

func (o *ModifySettingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ModifySettingResponse) GetObject() *ModifySettingResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
