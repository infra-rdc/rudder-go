package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ReloadOneDatasourceOneNodeRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
	// Id of the data source
	DatasourceID string `pathParam:"style=simple,explode=false,name=datasourceId"`
}

func (o *ReloadOneDatasourceOneNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *ReloadOneDatasourceOneNodeRequest) GetDatasourceID() string {
	if o == nil {
		return ""
	}
	return o.DatasourceID
}

// ReloadOneDatasourceOneNodeResult - Result of the request
type ReloadOneDatasourceOneNodeResult string

const (
	ReloadOneDatasourceOneNodeResultSuccess ReloadOneDatasourceOneNodeResult = "success"
	ReloadOneDatasourceOneNodeResultError   ReloadOneDatasourceOneNodeResult = "error"
)

func (e ReloadOneDatasourceOneNodeResult) ToPointer() *ReloadOneDatasourceOneNodeResult {
	return &e
}
func (e *ReloadOneDatasourceOneNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadOneDatasourceOneNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadOneDatasourceOneNodeResult: %v", v)
	}
}

// ReloadOneDatasourceOneNodeAction - The id of the action
type ReloadOneDatasourceOneNodeAction string

const (
	ReloadOneDatasourceOneNodeActionReloadOneDatasourceOneNode ReloadOneDatasourceOneNodeAction = "ReloadOneDatasourceOneNode"
)

func (e ReloadOneDatasourceOneNodeAction) ToPointer() *ReloadOneDatasourceOneNodeAction {
	return &e
}
func (e *ReloadOneDatasourceOneNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ReloadOneDatasourceOneNode":
		*e = ReloadOneDatasourceOneNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadOneDatasourceOneNodeAction: %v", v)
	}
}

// ReloadOneDatasourceOneNodeResponseBody - Data sources reloaded
type ReloadOneDatasourceOneNodeResponseBody struct {
	// Result of the request
	Result ReloadOneDatasourceOneNodeResult `json:"result"`
	// The id of the action
	Action ReloadOneDatasourceOneNodeAction `json:"action"`
	Data   string                           `json:"data"`
}

func (o *ReloadOneDatasourceOneNodeResponseBody) GetResult() ReloadOneDatasourceOneNodeResult {
	if o == nil {
		return ReloadOneDatasourceOneNodeResult("")
	}
	return o.Result
}

func (o *ReloadOneDatasourceOneNodeResponseBody) GetAction() ReloadOneDatasourceOneNodeAction {
	if o == nil {
		return ReloadOneDatasourceOneNodeAction("")
	}
	return o.Action
}

func (o *ReloadOneDatasourceOneNodeResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type ReloadOneDatasourceOneNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data sources reloaded
	Object *ReloadOneDatasourceOneNodeResponseBody
}

func (o *ReloadOneDatasourceOneNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadOneDatasourceOneNodeResponse) GetObject() *ReloadOneDatasourceOneNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
