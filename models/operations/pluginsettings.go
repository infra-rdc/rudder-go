package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// PluginSettingsResult - Result of the request
type PluginSettingsResult string

const (
	PluginSettingsResultSuccess PluginSettingsResult = "success"
	PluginSettingsResultError   PluginSettingsResult = "error"
)

func (e PluginSettingsResult) ToPointer() *PluginSettingsResult {
	return &e
}
func (e *PluginSettingsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = PluginSettingsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PluginSettingsResult: %v", v)
	}
}

// PluginSettingsAction - The id of the action
type PluginSettingsAction string

const (
	PluginSettingsActionPluginSettings PluginSettingsAction = "pluginSettings"
)

func (e PluginSettingsAction) ToPointer() *PluginSettingsAction {
	return &e
}
func (e *PluginSettingsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pluginSettings":
		*e = PluginSettingsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PluginSettingsAction: %v", v)
	}
}

// PluginSettingsData - Settings
type PluginSettingsData struct {
	Plugins []components.PluginsSettings `json:"plugins"`
}

func (o *PluginSettingsData) GetPlugins() []components.PluginsSettings {
	if o == nil {
		return []components.PluginsSettings{}
	}
	return o.Plugins
}

// PluginSettingsResponseBody - Settings
type PluginSettingsResponseBody struct {
	// Result of the request
	Result PluginSettingsResult `json:"result"`
	// The id of the action
	Action PluginSettingsAction `json:"action"`
	// Settings
	Data PluginSettingsData `json:"data"`
}

func (o *PluginSettingsResponseBody) GetResult() PluginSettingsResult {
	if o == nil {
		return PluginSettingsResult("")
	}
	return o.Result
}

func (o *PluginSettingsResponseBody) GetAction() PluginSettingsAction {
	if o == nil {
		return PluginSettingsAction("")
	}
	return o.Action
}

func (o *PluginSettingsResponseBody) GetData() PluginSettingsData {
	if o == nil {
		return PluginSettingsData{}
	}
	return o.Data
}

type PluginSettingsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *PluginSettingsResponseBody
}

func (o *PluginSettingsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PluginSettingsResponse) GetObject() *PluginSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
