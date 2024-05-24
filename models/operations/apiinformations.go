package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type APIInformationsRequest struct {
	// Name of the endpoint for which one wants information
	EndpointName string `pathParam:"style=simple,explode=false,name=endpointName"`
}

func (o *APIInformationsRequest) GetEndpointName() string {
	if o == nil {
		return ""
	}
	return o.EndpointName
}

// APIInformationsResult - Result of the request
type APIInformationsResult string

const (
	APIInformationsResultSuccess APIInformationsResult = "success"
	APIInformationsResultError   APIInformationsResult = "error"
)

func (e APIInformationsResult) ToPointer() *APIInformationsResult {
	return &e
}
func (e *APIInformationsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = APIInformationsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIInformationsResult: %v", v)
	}
}

// APIInformationsAction - The id of the action
type APIInformationsAction string

const (
	APIInformationsActionAPIInformations APIInformationsAction = "apiInformations"
)

func (e APIInformationsAction) ToPointer() *APIInformationsAction {
	return &e
}
func (e *APIInformationsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apiInformations":
		*e = APIInformationsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIInformationsAction: %v", v)
	}
}

// Endpoints - Supported path and version for that endpoint
type Endpoints struct {
	HTTPVerb any                    `json:"httpVerb,omitempty"`
	Version  *components.APIVersion `json:"version,omitempty"`
}

func (o *Endpoints) GetHTTPVerb() any {
	if o == nil {
		return nil
	}
	return o.HTTPVerb
}

func (o *Endpoints) GetVersion() *components.APIVersion {
	if o == nil {
		return nil
	}
	return o.Version
}

type APIInformationsData struct {
	// Link to Rudder API description
	Documentation string `json:"documentation"`
	// The endpoint name as key and the endpoint description as value
	EndpointName string      `json:"endpointName"`
	Endpoints    []Endpoints `json:"endpoints"`
}

func (o *APIInformationsData) GetDocumentation() string {
	if o == nil {
		return ""
	}
	return o.Documentation
}

func (o *APIInformationsData) GetEndpointName() string {
	if o == nil {
		return ""
	}
	return o.EndpointName
}

func (o *APIInformationsData) GetEndpoints() []Endpoints {
	if o == nil {
		return []Endpoints{}
	}
	return o.Endpoints
}

// APIInformationsResponseBody - API Endpoint information
type APIInformationsResponseBody struct {
	// Result of the request
	Result APIInformationsResult `json:"result"`
	// The id of the action
	Action APIInformationsAction `json:"action"`
	Data   APIInformationsData   `json:"data"`
}

func (o *APIInformationsResponseBody) GetResult() APIInformationsResult {
	if o == nil {
		return APIInformationsResult("")
	}
	return o.Result
}

func (o *APIInformationsResponseBody) GetAction() APIInformationsAction {
	if o == nil {
		return APIInformationsAction("")
	}
	return o.Action
}

func (o *APIInformationsResponseBody) GetData() APIInformationsData {
	if o == nil {
		return APIInformationsData{}
	}
	return o.Data
}

type APIInformationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// API Endpoint information
	Object *APIInformationsResponseBody
}

func (o *APIInformationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *APIInformationsResponse) GetObject() *APIInformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
