package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// UpdateCVERequestBody - CVE update config
type UpdateCVERequestBody struct {
	// Url used to update CVE, will default to one set in config
	URL   *string  `json:"url,omitempty"`
	Years []string `json:"years,omitempty"`
}

func (o *UpdateCVERequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UpdateCVERequestBody) GetYears() []string {
	if o == nil {
		return nil
	}
	return o.Years
}

// UpdateCVEResult - Result of the request
type UpdateCVEResult string

const (
	UpdateCVEResultSuccess UpdateCVEResult = "success"
	UpdateCVEResultError   UpdateCVEResult = "error"
)

func (e UpdateCVEResult) ToPointer() *UpdateCVEResult {
	return &e
}
func (e *UpdateCVEResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateCVEResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCVEResult: %v", v)
	}
}

// UpdateCVEAction - The id of the action
type UpdateCVEAction string

const (
	UpdateCVEActionUpdateCve UpdateCVEAction = "updateCVE"
)

func (e UpdateCVEAction) ToPointer() *UpdateCVEAction {
	return &e
}
func (e *UpdateCVEAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateCVE":
		*e = UpdateCVEAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateCVEAction: %v", v)
	}
}

type UpdateCVEData struct {
	CVEs int64 `json:"CVEs"`
}

func (o *UpdateCVEData) GetCVEs() int64 {
	if o == nil {
		return 0
	}
	return o.CVEs
}

// UpdateCVEResponseBody - updated CVE count
type UpdateCVEResponseBody struct {
	// Result of the request
	Result UpdateCVEResult `json:"result"`
	// The id of the action
	Action UpdateCVEAction `json:"action"`
	Data   UpdateCVEData   `json:"data"`
}

func (o *UpdateCVEResponseBody) GetResult() UpdateCVEResult {
	if o == nil {
		return UpdateCVEResult("")
	}
	return o.Result
}

func (o *UpdateCVEResponseBody) GetAction() UpdateCVEAction {
	if o == nil {
		return UpdateCVEAction("")
	}
	return o.Action
}

func (o *UpdateCVEResponseBody) GetData() UpdateCVEData {
	if o == nil {
		return UpdateCVEData{}
	}
	return o.Data
}

type UpdateCVEResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// updated CVE count
	Object *UpdateCVEResponseBody
}

func (o *UpdateCVEResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateCVEResponse) GetObject() *UpdateCVEResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
