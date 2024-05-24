package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ReloadAllDatasourcesAllNodesResult - Result of the request
type ReloadAllDatasourcesAllNodesResult string

const (
	ReloadAllDatasourcesAllNodesResultSuccess ReloadAllDatasourcesAllNodesResult = "success"
	ReloadAllDatasourcesAllNodesResultError   ReloadAllDatasourcesAllNodesResult = "error"
)

func (e ReloadAllDatasourcesAllNodesResult) ToPointer() *ReloadAllDatasourcesAllNodesResult {
	return &e
}
func (e *ReloadAllDatasourcesAllNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ReloadAllDatasourcesAllNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllDatasourcesAllNodesResult: %v", v)
	}
}

// ReloadAllDatasourcesAllNodesAction - The id of the action
type ReloadAllDatasourcesAllNodesAction string

const (
	ReloadAllDatasourcesAllNodesActionReloadAllDatasourcesAllNodes ReloadAllDatasourcesAllNodesAction = "ReloadAllDatasourcesAllNodes"
)

func (e ReloadAllDatasourcesAllNodesAction) ToPointer() *ReloadAllDatasourcesAllNodesAction {
	return &e
}
func (e *ReloadAllDatasourcesAllNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ReloadAllDatasourcesAllNodes":
		*e = ReloadAllDatasourcesAllNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReloadAllDatasourcesAllNodesAction: %v", v)
	}
}

// ReloadAllDatasourcesAllNodesResponseBody - Data source reloaded
type ReloadAllDatasourcesAllNodesResponseBody struct {
	// Result of the request
	Result ReloadAllDatasourcesAllNodesResult `json:"result"`
	// The id of the action
	Action ReloadAllDatasourcesAllNodesAction `json:"action"`
	Data   string                             `json:"data"`
}

func (o *ReloadAllDatasourcesAllNodesResponseBody) GetResult() ReloadAllDatasourcesAllNodesResult {
	if o == nil {
		return ReloadAllDatasourcesAllNodesResult("")
	}
	return o.Result
}

func (o *ReloadAllDatasourcesAllNodesResponseBody) GetAction() ReloadAllDatasourcesAllNodesAction {
	if o == nil {
		return ReloadAllDatasourcesAllNodesAction("")
	}
	return o.Action
}

func (o *ReloadAllDatasourcesAllNodesResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type ReloadAllDatasourcesAllNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Data source reloaded
	Object *ReloadAllDatasourcesAllNodesResponseBody
}

func (o *ReloadAllDatasourcesAllNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReloadAllDatasourcesAllNodesResponse) GetObject() *ReloadAllDatasourcesAllNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
