package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DemoteToNodeRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *DemoteToNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// DemoteToNodeAction - The id of the action
type DemoteToNodeAction string

const (
	DemoteToNodeActionDemoteToNode DemoteToNodeAction = "demoteToNode"
)

func (e DemoteToNodeAction) ToPointer() *DemoteToNodeAction {
	return &e
}
func (e *DemoteToNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "demoteToNode":
		*e = DemoteToNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DemoteToNodeAction: %v", v)
	}
}

// DemoteToNodeResult - Result of the request
type DemoteToNodeResult string

const (
	DemoteToNodeResultSuccess DemoteToNodeResult = "success"
	DemoteToNodeResultError   DemoteToNodeResult = "error"
)

func (e DemoteToNodeResult) ToPointer() *DemoteToNodeResult {
	return &e
}
func (e *DemoteToNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DemoteToNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DemoteToNodeResult: %v", v)
	}
}

// DemoteToNodeResponseBody - Demote to node response
type DemoteToNodeResponseBody struct {
	// The id of the action
	Action DemoteToNodeAction `json:"action"`
	// Result of the request
	Result DemoteToNodeResult `json:"result"`
	// Success or error message
	Data string `json:"data"`
}

func (o *DemoteToNodeResponseBody) GetAction() DemoteToNodeAction {
	if o == nil {
		return DemoteToNodeAction("")
	}
	return o.Action
}

func (o *DemoteToNodeResponseBody) GetResult() DemoteToNodeResult {
	if o == nil {
		return DemoteToNodeResult("")
	}
	return o.Result
}

func (o *DemoteToNodeResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type DemoteToNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Demote to node response
	Object *DemoteToNodeResponseBody
}

func (o *DemoteToNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DemoteToNodeResponse) GetObject() *DemoteToNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
