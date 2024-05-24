package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type APISubInformationsRequest struct {
	// Id of the API section
	SectionID string `pathParam:"style=simple,explode=false,name=sectionId"`
}

func (o *APISubInformationsRequest) GetSectionID() string {
	if o == nil {
		return ""
	}
	return o.SectionID
}

// APISubInformationsResult - Result of the request
type APISubInformationsResult string

const (
	APISubInformationsResultSuccess APISubInformationsResult = "success"
	APISubInformationsResultError   APISubInformationsResult = "error"
)

func (e APISubInformationsResult) ToPointer() *APISubInformationsResult {
	return &e
}
func (e *APISubInformationsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = APISubInformationsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APISubInformationsResult: %v", v)
	}
}

// APISubInformationsAction - The id of the action
type APISubInformationsAction string

const (
	APISubInformationsActionAPISubInformations APISubInformationsAction = "apiSubInformations"
)

func (e APISubInformationsAction) ToPointer() *APISubInformationsAction {
	return &e
}
func (e *APISubInformationsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apiSubInformations":
		*e = APISubInformationsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APISubInformationsAction: %v", v)
	}
}

type APISubInformationsData struct {
	// Link to Rudder API description
	Documentation     string                      `json:"documentation"`
	AvailableVersions []components.APIVersions    `json:"availableVersions"`
	Endpoints         [][]components.APIEndpoints `json:"endpoints"`
}

func (o *APISubInformationsData) GetDocumentation() string {
	if o == nil {
		return ""
	}
	return o.Documentation
}

func (o *APISubInformationsData) GetAvailableVersions() []components.APIVersions {
	if o == nil {
		return []components.APIVersions{}
	}
	return o.AvailableVersions
}

func (o *APISubInformationsData) GetEndpoints() [][]components.APIEndpoints {
	if o == nil {
		return [][]components.APIEndpoints{}
	}
	return o.Endpoints
}

// APISubInformationsResponseBody - Endpoint information
type APISubInformationsResponseBody struct {
	// Result of the request
	Result APISubInformationsResult `json:"result"`
	// The id of the action
	Action APISubInformationsAction `json:"action"`
	Data   APISubInformationsData   `json:"data"`
}

func (o *APISubInformationsResponseBody) GetResult() APISubInformationsResult {
	if o == nil {
		return APISubInformationsResult("")
	}
	return o.Result
}

func (o *APISubInformationsResponseBody) GetAction() APISubInformationsAction {
	if o == nil {
		return APISubInformationsAction("")
	}
	return o.Action
}

func (o *APISubInformationsResponseBody) GetData() APISubInformationsData {
	if o == nil {
		return APISubInformationsData{}
	}
	return o.Data
}

type APISubInformationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Endpoint information
	Object *APISubInformationsResponseBody
}

func (o *APISubInformationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *APISubInformationsResponse) GetObject() *APISubInformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
