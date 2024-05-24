package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetCVECheckConfigurationResult - Result of the request
type GetCVECheckConfigurationResult string

const (
	GetCVECheckConfigurationResultSuccess GetCVECheckConfigurationResult = "success"
	GetCVECheckConfigurationResultError   GetCVECheckConfigurationResult = "error"
)

func (e GetCVECheckConfigurationResult) ToPointer() *GetCVECheckConfigurationResult {
	return &e
}
func (e *GetCVECheckConfigurationResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetCVECheckConfigurationResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCVECheckConfigurationResult: %v", v)
	}
}

// GetCVECheckConfigurationAction - The id of the action
type GetCVECheckConfigurationAction string

const (
	GetCVECheckConfigurationActionGetCveCheckConfiguration GetCVECheckConfigurationAction = "getCVECheckConfiguration"
)

func (e GetCVECheckConfigurationAction) ToPointer() *GetCVECheckConfigurationAction {
	return &e
}
func (e *GetCVECheckConfigurationAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getCVECheckConfiguration":
		*e = GetCVECheckConfigurationAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCVECheckConfigurationAction: %v", v)
	}
}

type GetCVECheckConfigurationData struct {
	// Url used to check CVE
	URL *string `json:"url,omitempty"`
	// Token used by to contact the API to check CVE
	APIKey *string `json:"apiKey,omitempty"`
}

func (o *GetCVECheckConfigurationData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *GetCVECheckConfigurationData) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

// GetCVECheckConfigurationResponseBody - CVE check config
type GetCVECheckConfigurationResponseBody struct {
	// Result of the request
	Result GetCVECheckConfigurationResult `json:"result"`
	// The id of the action
	Action GetCVECheckConfigurationAction `json:"action"`
	Data   GetCVECheckConfigurationData   `json:"data"`
}

func (o *GetCVECheckConfigurationResponseBody) GetResult() GetCVECheckConfigurationResult {
	if o == nil {
		return GetCVECheckConfigurationResult("")
	}
	return o.Result
}

func (o *GetCVECheckConfigurationResponseBody) GetAction() GetCVECheckConfigurationAction {
	if o == nil {
		return GetCVECheckConfigurationAction("")
	}
	return o.Action
}

func (o *GetCVECheckConfigurationResponseBody) GetData() GetCVECheckConfigurationData {
	if o == nil {
		return GetCVECheckConfigurationData{}
	}
	return o.Data
}

type GetCVECheckConfigurationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// CVE check config
	Object *GetCVECheckConfigurationResponseBody
}

func (o *GetCVECheckConfigurationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCVECheckConfigurationResponse) GetObject() *GetCVECheckConfigurationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
