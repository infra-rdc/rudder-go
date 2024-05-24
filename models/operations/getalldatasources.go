package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetAllDataSourcesResult - Result of the request
type GetAllDataSourcesResult string

const (
	GetAllDataSourcesResultSuccess GetAllDataSourcesResult = "success"
	GetAllDataSourcesResultError   GetAllDataSourcesResult = "error"
)

func (e GetAllDataSourcesResult) ToPointer() *GetAllDataSourcesResult {
	return &e
}
func (e *GetAllDataSourcesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllDataSourcesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllDataSourcesResult: %v", v)
	}
}

// GetAllDataSourcesAction - The id of the action
type GetAllDataSourcesAction string

const (
	GetAllDataSourcesActionGetAllDataSources GetAllDataSourcesAction = "getAllDataSources"
)

func (e GetAllDataSourcesAction) ToPointer() *GetAllDataSourcesAction {
	return &e
}
func (e *GetAllDataSourcesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getAllDataSources":
		*e = GetAllDataSourcesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllDataSourcesAction: %v", v)
	}
}

type GetAllDataSourcesData struct {
	Datasources []components.Datasource `json:"datasources"`
}

func (o *GetAllDataSourcesData) GetDatasources() []components.Datasource {
	if o == nil {
		return []components.Datasource{}
	}
	return o.Datasources
}

// GetAllDataSourcesResponseBody - Data sources information
type GetAllDataSourcesResponseBody struct {
	// Result of the request
	Result GetAllDataSourcesResult `json:"result"`
	// The id of the action
	Action GetAllDataSourcesAction `json:"action"`
	Data   GetAllDataSourcesData   `json:"data"`
}

func (o *GetAllDataSourcesResponseBody) GetResult() GetAllDataSourcesResult {
	if o == nil {
		return GetAllDataSourcesResult("")
	}
	return o.Result
}

func (o *GetAllDataSourcesResponseBody) GetAction() GetAllDataSourcesAction {
	if o == nil {
		return GetAllDataSourcesAction("")
	}
	return o.Action
}

func (o *GetAllDataSourcesResponseBody) GetData() GetAllDataSourcesData {
	if o == nil {
		return GetAllDataSourcesData{}
	}
	return o.Data
}

type GetAllDataSourcesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data sources information
	Object *GetAllDataSourcesResponseBody
}

func (o *GetAllDataSourcesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllDataSourcesResponse) GetObject() *GetAllDataSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
