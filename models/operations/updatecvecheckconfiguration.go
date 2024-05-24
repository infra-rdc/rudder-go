package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// UpdateCVECheckConfigurationRequestBody - CVE check config
type UpdateCVECheckConfigurationRequestBody struct {
	// Url used to check CVE
	URL *string `json:"url,omitempty"`
	// Token used by to contact the API to check CVE
	APIKey *string `json:"apiKey,omitempty"`
}

func (o *UpdateCVECheckConfigurationRequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UpdateCVECheckConfigurationRequestBody) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

// UpdateCVECheckConfigurationResult - Result of the request
type UpdateCVECheckConfigurationResult string

const (
	UpdateCVECheckConfigurationResultSuccess UpdateCVECheckConfigurationResult = "success"
	UpdateCVECheckConfigurationResultError   UpdateCVECheckConfigurationResult = "error"
)

func (e UpdateCVECheckConfigurationResult) ToPointer() *UpdateCVECheckConfigurationResult {
	return &e
}
func (e *UpdateCVECheckConfigurationResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateCVECheckConfigurationResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCVECheckConfigurationResult: %v", v)
	}
}

// UpdateCVECheckConfigurationAction - The id of the action
type UpdateCVECheckConfigurationAction string

const (
	UpdateCVECheckConfigurationActionUpdateCveCheckConfiguration UpdateCVECheckConfigurationAction = "updateCVECheckConfiguration"
)

func (e UpdateCVECheckConfigurationAction) ToPointer() *UpdateCVECheckConfigurationAction {
	return &e
}
func (e *UpdateCVECheckConfigurationAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateCVECheckConfiguration":
		*e = UpdateCVECheckConfigurationAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCVECheckConfigurationAction: %v", v)
	}
}

type UpdateCVECheckConfigurationData struct {
	// Url used to check CVE
	URL *string `json:"url,omitempty"`
	// Token used by to contact the API to check CVE
	APIKey *string `json:"apiKey,omitempty"`
}

func (o *UpdateCVECheckConfigurationData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UpdateCVECheckConfigurationData) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

// UpdateCVECheckConfigurationResponseBody - new CVE check config
type UpdateCVECheckConfigurationResponseBody struct {
	// Result of the request
	Result UpdateCVECheckConfigurationResult `json:"result"`
	// The id of the action
	Action UpdateCVECheckConfigurationAction `json:"action"`
	Data   UpdateCVECheckConfigurationData   `json:"data"`
}

func (o *UpdateCVECheckConfigurationResponseBody) GetResult() UpdateCVECheckConfigurationResult {
	if o == nil {
		return UpdateCVECheckConfigurationResult("")
	}
	return o.Result
}

func (o *UpdateCVECheckConfigurationResponseBody) GetAction() UpdateCVECheckConfigurationAction {
	if o == nil {
		return UpdateCVECheckConfigurationAction("")
	}
	return o.Action
}

func (o *UpdateCVECheckConfigurationResponseBody) GetData() UpdateCVECheckConfigurationData {
	if o == nil {
		return UpdateCVECheckConfigurationData{}
	}
	return o.Data
}

type UpdateCVECheckConfigurationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// new CVE check config
	Object *UpdateCVECheckConfigurationResponseBody
}

func (o *UpdateCVECheckConfigurationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateCVECheckConfigurationResponse) GetObject() *UpdateCVECheckConfigurationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
