package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// UpdateSettingsResult - Result of the request
type UpdateSettingsResult string

const (
	UpdateSettingsResultSuccess UpdateSettingsResult = "success"
	UpdateSettingsResultError   UpdateSettingsResult = "error"
)

func (e UpdateSettingsResult) ToPointer() *UpdateSettingsResult {
	return &e
}
func (e *UpdateSettingsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateSettingsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSettingsResult: %v", v)
	}
}

// UpdateSettingsAction - The id of the action
type UpdateSettingsAction string

const (
	UpdateSettingsActionUpdateSettings UpdateSettingsAction = "updateSettings"
)

func (e UpdateSettingsAction) ToPointer() *UpdateSettingsAction {
	return &e
}
func (e *UpdateSettingsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateSettings":
		*e = UpdateSettingsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSettingsAction: %v", v)
	}
}

// UpdateSettingsData - Parameters
type UpdateSettingsData struct {
	Parameters []components.PluginsSettings `json:"parameters,omitempty"`
}

func (o *UpdateSettingsData) GetParameters() []components.PluginsSettings {
	if o == nil {
		return nil
	}
	return o.Parameters
}

// UpdateSettingsResponseBody - Settings
type UpdateSettingsResponseBody struct {
	// Result of the request
	Result UpdateSettingsResult `json:"result"`
	// The id of the action
	Action UpdateSettingsAction `json:"action"`
	// Parameters
	Data UpdateSettingsData `json:"data"`
}

func (o *UpdateSettingsResponseBody) GetResult() UpdateSettingsResult {
	if o == nil {
		return UpdateSettingsResult("")
	}
	return o.Result
}

func (o *UpdateSettingsResponseBody) GetAction() UpdateSettingsAction {
	if o == nil {
		return UpdateSettingsAction("")
	}
	return o.Action
}

func (o *UpdateSettingsResponseBody) GetData() UpdateSettingsData {
	if o == nil {
		return UpdateSettingsData{}
	}
	return o.Data
}

type UpdateSettingsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *UpdateSettingsResponseBody
}

func (o *UpdateSettingsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateSettingsResponse) GetObject() *UpdateSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
