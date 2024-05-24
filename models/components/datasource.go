package components

import (
	"encoding/json"
	"fmt"
)

// DatasourceType - `scheduled` enables periodic update, `notscheduled` disables them
type DatasourceType string

const (
	DatasourceTypeScheduled    DatasourceType = "scheduled"
	DatasourceTypeNotscheduled DatasourceType = "notscheduled"
)

func (e DatasourceType) ToPointer() *DatasourceType {
	return &e
}
func (e *DatasourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "scheduled":
		fallthrough
	case "notscheduled":
		*e = DatasourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatasourceType: %v", v)
	}
}

// Schedule - Configure if data source should be fetch periodically
type Schedule struct {
	// `scheduled` enables periodic update, `notscheduled` disables them
	Type *DatasourceType `json:"type,omitempty"`
}

func (o *Schedule) GetType() *DatasourceType {
	if o == nil {
		return nil
	}
	return o.Type
}

// RunParameters - Parameters to configure when the data source is fetched to update node properties.
type RunParameters struct {
	// Trigger a fetch at the beginning of a policy generation
	OnGeneration *bool `json:"onGeneration,omitempty"`
	// Trigger a fetch when a new node is accepted, for that node
	OnNewNode *bool `json:"onNewNode,omitempty"`
	// Configure if data source should be fetch periodically
	Schedule *Schedule `json:"schedule,omitempty"`
}

func (o *RunParameters) GetOnGeneration() *bool {
	if o == nil {
		return nil
	}
	return o.OnGeneration
}

func (o *RunParameters) GetOnNewNode() *bool {
	if o == nil {
		return nil
	}
	return o.OnNewNode
}

func (o *RunParameters) GetSchedule() *Schedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

// DatasourceName - Data source type name
type DatasourceName string

const (
	DatasourceNameHTTP DatasourceName = "HTTP"
)

func (e DatasourceName) ToPointer() *DatasourceName {
	return &e
}
func (e *DatasourceName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTP":
		*e = DatasourceName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatasourceName: %v", v)
	}
}

// RequestMethod - HTTP method to use to contact the URL.
type RequestMethod string

const (
	RequestMethodGet  RequestMethod = "GET"
	RequestMethodPost RequestMethod = "POST"
)

func (e RequestMethod) ToPointer() *RequestMethod {
	return &e
}
func (e *RequestMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		*e = RequestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestMethod: %v", v)
	}
}

type Headers struct {
	// Name of the header
	Name *string `json:"name,omitempty"`
	// Value of the header
	Value *string `json:"value,omitempty"`
}

func (o *Headers) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Headers) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// DatasourceTypeName - Node by node strategy
type DatasourceTypeName string

const (
	DatasourceTypeNameByNode DatasourceTypeName = "byNode"
)

func (e DatasourceTypeName) ToPointer() *DatasourceTypeName {
	return &e
}
func (e *DatasourceTypeName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "byNode":
		*e = DatasourceTypeName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatasourceTypeName: %v", v)
	}
}

// RequestMode - Configure the strategy used to query the HTTP data source.
type RequestMode struct {
	// Node by node strategy
	Name *DatasourceTypeName `json:"name,omitempty"`
}

func (o *RequestMode) GetName() *DatasourceTypeName {
	if o == nil {
		return nil
	}
	return o.Name
}

// DatasourceParameters - You can use Rudder variable expansion (`${rudder.node`, `${node.properties...}`)
type DatasourceParameters struct {
	// URL to contact. Rudder expansion available.
	URL *string `json:"url,omitempty"`
	// HTTP method to use to contact the URL.
	RequestMethod *RequestMethod `json:"requestMethod,omitempty"`
	// Represent HTTP headers for the query. Rudder expansion available.
	Headers []Headers `json:"headers,omitempty"`
	// JSON path (as defined in [the specification](https://github.com/jayway/JsonPath/), without the leading `$.`) to find the interesting sub-json or string/number/boolean value in the answer. Let empty to use the whole answer as value.
	Path *string `json:"path,omitempty"`
	// Check SSL certificate validity for https. Must be set to false for self-signed certificate
	CheckSsl *bool `json:"checkSsl,omitempty"`
	// Timeout in seconds for each HTTP request
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Configure the strategy used to query the HTTP data source.
	RequestMode *RequestMode `json:"requestMode,omitempty"`
}

func (o *DatasourceParameters) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *DatasourceParameters) GetRequestMethod() *RequestMethod {
	if o == nil {
		return nil
	}
	return o.RequestMethod
}

func (o *DatasourceParameters) GetHeaders() []Headers {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *DatasourceParameters) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *DatasourceParameters) GetCheckSsl() *bool {
	if o == nil {
		return nil
	}
	return o.CheckSsl
}

func (o *DatasourceParameters) GetRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *DatasourceParameters) GetRequestMode() *RequestMode {
	if o == nil {
		return nil
	}
	return o.RequestMode
}

// Type - Define and configure data source type.
type Type struct {
	// Data source type name
	Name *DatasourceName `json:"name,omitempty"`
	// You can use Rudder variable expansion (`${rudder.node`, `${node.properties...}`)
	Parameters *DatasourceParameters `json:"parameters,omitempty"`
}

func (o *Type) GetName() *DatasourceName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Type) GetParameters() *DatasourceParameters {
	if o == nil {
		return nil
	}
	return o.Parameters
}

type Datasource struct {
	// Unique identifier of the data source to create.
	ID *string `json:"id,omitempty"`
	// The human readable name of the data source to create.
	Name *string `json:"name,omitempty"`
	// Description of the goal of the data source to create.
	Description *string `json:"description,omitempty"`
	// Enable or disable data source.
	Enabled *bool `json:"enabled,omitempty"`
	// Duration in seconds before aborting data source update. The main goal is to prevent never ending requests. If a periodicity if configured, you should set that timeout at a lower value.
	UpdateTimeout *int64 `json:"updateTimeout,omitempty"`
	// Parameters to configure when the data source is fetched to update node properties.
	RunParameters *RunParameters `json:"runParameters,omitempty"`
	// Define and configure data source type.
	Type *Type `json:"type,omitempty"`
}

func (o *Datasource) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Datasource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Datasource) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Datasource) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Datasource) GetUpdateTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdateTimeout
}

func (o *Datasource) GetRunParameters() *RunParameters {
	if o == nil {
		return nil
	}
	return o.RunParameters
}

func (o *Datasource) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
