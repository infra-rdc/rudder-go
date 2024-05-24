package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteDataSourceRequest struct {
	// Id of the data source
	DatasourceID string `pathParam:"style=simple,explode=false,name=datasourceId"`
}

func (o *DeleteDataSourceRequest) GetDatasourceID() string {
	if o == nil {
		return ""
	}
	return o.DatasourceID
}

// DeleteDataSourceResult - Result of the request
type DeleteDataSourceResult string

const (
	DeleteDataSourceResultSuccess DeleteDataSourceResult = "success"
	DeleteDataSourceResultError   DeleteDataSourceResult = "error"
)

func (e DeleteDataSourceResult) ToPointer() *DeleteDataSourceResult {
	return &e
}
func (e *DeleteDataSourceResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteDataSourceResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteDataSourceResult: %v", v)
	}
}

// DeleteDataSourceAction - The id of the action
type DeleteDataSourceAction string

const (
	DeleteDataSourceActionDeleteDataSource DeleteDataSourceAction = "deleteDataSource"
)

func (e DeleteDataSourceAction) ToPointer() *DeleteDataSourceAction {
	return &e
}
func (e *DeleteDataSourceAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteDataSource":
		*e = DeleteDataSourceAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteDataSourceAction: %v", v)
	}
}

type DeleteDataSourceData struct {
	Datasources []components.Datasource `json:"datasources"`
}

func (o *DeleteDataSourceData) GetDatasources() []components.Datasource {
	if o == nil {
		return []components.Datasource{}
	}
	return o.Datasources
}

// DeleteDataSourceResponseBody - Data source information
type DeleteDataSourceResponseBody struct {
	// Result of the request
	Result DeleteDataSourceResult `json:"result"`
	// The id of the action
	Action DeleteDataSourceAction `json:"action"`
	Data   DeleteDataSourceData   `json:"data"`
}

func (o *DeleteDataSourceResponseBody) GetResult() DeleteDataSourceResult {
	if o == nil {
		return DeleteDataSourceResult("")
	}
	return o.Result
}

func (o *DeleteDataSourceResponseBody) GetAction() DeleteDataSourceAction {
	if o == nil {
		return DeleteDataSourceAction("")
	}
	return o.Action
}

func (o *DeleteDataSourceResponseBody) GetData() DeleteDataSourceData {
	if o == nil {
		return DeleteDataSourceData{}
	}
	return o.Data
}

type DeleteDataSourceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data source information
	Object *DeleteDataSourceResponseBody
}

func (o *DeleteDataSourceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteDataSourceResponse) GetObject() *DeleteDataSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
