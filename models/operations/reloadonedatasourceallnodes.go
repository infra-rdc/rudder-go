package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ReloadOneDatasourceAllNodesRequest struct {
	// Id of the data source
	DatasourceID string `pathParam:"style=simple,explode=false,name=datasourceId"`
}

func (o *ReloadOneDatasourceAllNodesRequest) GetDatasourceID() string {
	if o == nil {
		return ""
	}
	return o.DatasourceID
}

// ReloadOneDatasourceAllNodesResult - Result of the request
type ReloadOneDatasourceAllNodesResult string

const (
	ReloadOneDatasourceAllNodesResultSuccess ReloadOneDatasourceAllNodesResult = "success"
	ReloadOneDatasourceAllNodesResultError   ReloadOneDatasourceAllNodesResult = "error"
)

func (e ReloadOneDatasourceAllNodesResult) ToPointer() *ReloadOneDatasourceAllNodesResult {
	return &e
}
func (e *ReloadOneDatasourceAllNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadOneDatasourceAllNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadOneDatasourceAllNodesResult: %v", v)
	}
}

// ReloadOneDatasourceAllNodesAction - The id of the action
type ReloadOneDatasourceAllNodesAction string

const (
	ReloadOneDatasourceAllNodesActionReloadOneDatasourceAllNodes ReloadOneDatasourceAllNodesAction = "ReloadOneDatasourceAllNodes"
)

func (e ReloadOneDatasourceAllNodesAction) ToPointer() *ReloadOneDatasourceAllNodesAction {
	return &e
}
func (e *ReloadOneDatasourceAllNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ReloadOneDatasourceAllNodes":
		*e = ReloadOneDatasourceAllNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadOneDatasourceAllNodesAction: %v", v)
	}
}

// ReloadOneDatasourceAllNodesResponseBody - Data source reloaded
type ReloadOneDatasourceAllNodesResponseBody struct {
	// Result of the request
	Result ReloadOneDatasourceAllNodesResult `json:"result"`
	// The id of the action
	Action ReloadOneDatasourceAllNodesAction `json:"action"`
	Data   string                            `json:"data"`
}

func (o *ReloadOneDatasourceAllNodesResponseBody) GetResult() ReloadOneDatasourceAllNodesResult {
	if o == nil {
		return ReloadOneDatasourceAllNodesResult("")
	}
	return o.Result
}

func (o *ReloadOneDatasourceAllNodesResponseBody) GetAction() ReloadOneDatasourceAllNodesAction {
	if o == nil {
		return ReloadOneDatasourceAllNodesAction("")
	}
	return o.Action
}

func (o *ReloadOneDatasourceAllNodesResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type ReloadOneDatasourceAllNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data source reloaded
	Object *ReloadOneDatasourceAllNodesResponseBody
}

func (o *ReloadOneDatasourceAllNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadOneDatasourceAllNodesResponse) GetObject() *ReloadOneDatasourceAllNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
