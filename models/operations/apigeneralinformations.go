package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// APIGeneralInformationsResult - Result of the request
type APIGeneralInformationsResult string

const (
	APIGeneralInformationsResultSuccess APIGeneralInformationsResult = "success"
	APIGeneralInformationsResultError   APIGeneralInformationsResult = "error"
)

func (e APIGeneralInformationsResult) ToPointer() *APIGeneralInformationsResult {
	return &e
}
func (e *APIGeneralInformationsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = APIGeneralInformationsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIGeneralInformationsResult: %v", v)
	}
}

// APIGeneralInformationsAction - The id of the action
type APIGeneralInformationsAction string

const (
	APIGeneralInformationsActionAPIGeneralInformations APIGeneralInformationsAction = "apiGeneralInformations"
)

func (e APIGeneralInformationsAction) ToPointer() *APIGeneralInformationsAction {
	return &e
}
func (e *APIGeneralInformationsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apiGeneralInformations":
		*e = APIGeneralInformationsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIGeneralInformationsAction: %v", v)
	}
}

type APIGeneralInformationsData struct {
	// Link to Rudder API description
	Documentation     string                      `json:"documentation"`
	AvailableVersions []components.APIVersions    `json:"availableVersions"`
	Endpoints         [][]components.APIEndpoints `json:"endpoints"`
}

func (o *APIGeneralInformationsData) GetDocumentation() string {
	if o == nil {
		return ""
	}
	return o.Documentation
}

func (o *APIGeneralInformationsData) GetAvailableVersions() []components.APIVersions {
	if o == nil {
		return []components.APIVersions{}
	}
	return o.AvailableVersions
}

func (o *APIGeneralInformationsData) GetEndpoints() [][]components.APIEndpoints {
	if o == nil {
		return [][]components.APIEndpoints{}
	}
	return o.Endpoints
}

// APIGeneralInformationsResponseBody - API endpoints
type APIGeneralInformationsResponseBody struct {
	// Result of the request
	Result APIGeneralInformationsResult `json:"result"`
	// The id of the action
	Action APIGeneralInformationsAction `json:"action"`
	Data   APIGeneralInformationsData   `json:"data"`
}

func (o *APIGeneralInformationsResponseBody) GetResult() APIGeneralInformationsResult {
	if o == nil {
		return APIGeneralInformationsResult("")
	}
	return o.Result
}

func (o *APIGeneralInformationsResponseBody) GetAction() APIGeneralInformationsAction {
	if o == nil {
		return APIGeneralInformationsAction("")
	}
	return o.Action
}

func (o *APIGeneralInformationsResponseBody) GetData() APIGeneralInformationsData {
	if o == nil {
		return APIGeneralInformationsData{}
	}
	return o.Data
}

type APIGeneralInformationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// API endpoints
	Object *APIGeneralInformationsResponseBody
}

func (o *APIGeneralInformationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *APIGeneralInformationsResponse) GetObject() *APIGeneralInformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
