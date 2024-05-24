package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetDataSourceRequest struct {
	// Id of the data source
	DatasourceID string `pathParam:"style=simple,explode=false,name=datasourceId"`
}

func (o *GetDataSourceRequest) GetDatasourceID() string {
	if o == nil {
		return ""
	}
	return o.DatasourceID
}

// GetDataSourceResult - Result of the request
type GetDataSourceResult string

const (
	GetDataSourceResultSuccess GetDataSourceResult = "success"
	GetDataSourceResultError   GetDataSourceResult = "error"
)

func (e GetDataSourceResult) ToPointer() *GetDataSourceResult {
	return &e
}
func (e *GetDataSourceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetDataSourceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDataSourceResult: %v", v)
	}
}

// GetDataSourceAction - The id of the action
type GetDataSourceAction string

const (
	GetDataSourceActionGetDataSource GetDataSourceAction = "getDataSource"
)

func (e GetDataSourceAction) ToPointer() *GetDataSourceAction {
	return &e
}
func (e *GetDataSourceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getDataSource":
		*e = GetDataSourceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDataSourceAction: %v", v)
	}
}

type GetDataSourceData struct {
	Datasources []components.Datasource `json:"datasources"`
}

func (o *GetDataSourceData) GetDatasources() []components.Datasource {
	if o == nil {
		return []components.Datasource{}
	}
	return o.Datasources
}

// GetDataSourceResponseBody - Data source information
type GetDataSourceResponseBody struct {
	// Result of the request
	Result GetDataSourceResult `json:"result"`
	// The id of the action
	Action GetDataSourceAction `json:"action"`
	Data   GetDataSourceData   `json:"data"`
}

func (o *GetDataSourceResponseBody) GetResult() GetDataSourceResult {
	if o == nil {
		return GetDataSourceResult("")
	}
	return o.Result
}

func (o *GetDataSourceResponseBody) GetAction() GetDataSourceAction {
	if o == nil {
		return GetDataSourceAction("")
	}
	return o.Action
}

func (o *GetDataSourceResponseBody) GetData() GetDataSourceData {
	if o == nil {
		return GetDataSourceData{}
	}
	return o.Data
}

type GetDataSourceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data source information
	Object *GetDataSourceResponseBody
}

func (o *GetDataSourceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetDataSourceResponse) GetObject() *GetDataSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
