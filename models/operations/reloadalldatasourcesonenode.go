package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type ReloadAllDatasourcesOneNodeRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *ReloadAllDatasourcesOneNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// ReloadAllDatasourcesOneNodeResult - Result of the request
type ReloadAllDatasourcesOneNodeResult string

const (
	ReloadAllDatasourcesOneNodeResultSuccess ReloadAllDatasourcesOneNodeResult = "success"
	ReloadAllDatasourcesOneNodeResultError   ReloadAllDatasourcesOneNodeResult = "error"
)

func (e ReloadAllDatasourcesOneNodeResult) ToPointer() *ReloadAllDatasourcesOneNodeResult {
	return &e
}
func (e *ReloadAllDatasourcesOneNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadAllDatasourcesOneNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllDatasourcesOneNodeResult: %v", v)
	}
}

// ReloadAllDatasourcesOneNodeAction - The id of the action
type ReloadAllDatasourcesOneNodeAction string

const (
	ReloadAllDatasourcesOneNodeActionReloadAllDatasourcesOneNode ReloadAllDatasourcesOneNodeAction = "ReloadAllDatasourcesOneNode"
)

func (e ReloadAllDatasourcesOneNodeAction) ToPointer() *ReloadAllDatasourcesOneNodeAction {
	return &e
}
func (e *ReloadAllDatasourcesOneNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ReloadAllDatasourcesOneNode":
		*e = ReloadAllDatasourcesOneNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllDatasourcesOneNodeAction: %v", v)
	}
}

// ReloadAllDatasourcesOneNodeResponseBody - Data sources reloaded
type ReloadAllDatasourcesOneNodeResponseBody struct {
	// Result of the request
	Result ReloadAllDatasourcesOneNodeResult `json:"result"`
	// The id of the action
	Action ReloadAllDatasourcesOneNodeAction `json:"action"`
	Data   string                            `json:"data"`
}

func (o *ReloadAllDatasourcesOneNodeResponseBody) GetResult() ReloadAllDatasourcesOneNodeResult {
	if o == nil {
		return ReloadAllDatasourcesOneNodeResult("")
	}
	return o.Result
}

func (o *ReloadAllDatasourcesOneNodeResponseBody) GetAction() ReloadAllDatasourcesOneNodeAction {
	if o == nil {
		return ReloadAllDatasourcesOneNodeAction("")
	}
	return o.Action
}

func (o *ReloadAllDatasourcesOneNodeResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type ReloadAllDatasourcesOneNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data sources reloaded
	Object *ReloadAllDatasourcesOneNodeResponseBody
}

func (o *ReloadAllDatasourcesOneNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadAllDatasourcesOneNodeResponse) GetObject() *ReloadAllDatasourcesOneNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
