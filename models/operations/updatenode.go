package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateNodeRequest struct {
	// Id of the target node
	NodeID       string                   `pathParam:"style=simple,explode=false,name=nodeId"`
	NodeSettings *components.NodeSettings `request:"mediaType=application/json"`
}

func (o *UpdateNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *UpdateNodeRequest) GetNodeSettings() *components.NodeSettings {
	if o == nil {
		return nil
	}
	return o.NodeSettings
}

// UpdateNodeResult - Result of the request
type UpdateNodeResult string

const (
	UpdateNodeResultSuccess UpdateNodeResult = "success"
	UpdateNodeResultError   UpdateNodeResult = "error"
)

func (e UpdateNodeResult) ToPointer() *UpdateNodeResult {
	return &e
}
func (e *UpdateNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateNodeResult: %v", v)
	}
}

// UpdateNodeAction - The id of the action
type UpdateNodeAction string

const (
	UpdateNodeActionUpdateNode UpdateNodeAction = "updateNode"
)

func (e UpdateNodeAction) ToPointer() *UpdateNodeAction {
	return &e
}
func (e *UpdateNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updateNode":
		*e = UpdateNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateNodeAction: %v", v)
	}
}

// UpdateNodeData - Information about the node
type UpdateNodeData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *UpdateNodeData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// UpdateNodeResponseBody - Nodes
type UpdateNodeResponseBody struct {
	// Result of the request
	Result UpdateNodeResult `json:"result"`
	// The id of the action
	Action UpdateNodeAction `json:"action"`
	// Information about the node
	Data UpdateNodeData `json:"data"`
}

func (o *UpdateNodeResponseBody) GetResult() UpdateNodeResult {
	if o == nil {
		return UpdateNodeResult("")
	}
	return o.Result
}

func (o *UpdateNodeResponseBody) GetAction() UpdateNodeAction {
	if o == nil {
		return UpdateNodeAction("")
	}
	return o.Action
}

func (o *UpdateNodeResponseBody) GetData() UpdateNodeData {
	if o == nil {
		return UpdateNodeData{}
	}
	return o.Data
}

type UpdateNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *UpdateNodeResponseBody
}

func (o *UpdateNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateNodeResponse) GetObject() *UpdateNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
