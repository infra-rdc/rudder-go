package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type NodeInheritedPropertiesRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *NodeInheritedPropertiesRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// NodeInheritedPropertiesResult - Result of the request
type NodeInheritedPropertiesResult string

const (
	NodeInheritedPropertiesResultSuccess NodeInheritedPropertiesResult = "success"
	NodeInheritedPropertiesResultError   NodeInheritedPropertiesResult = "error"
)

func (e NodeInheritedPropertiesResult) ToPointer() *NodeInheritedPropertiesResult {
	return &e
}
func (e *NodeInheritedPropertiesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = NodeInheritedPropertiesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeInheritedPropertiesResult: %v", v)
	}
}

// NodeInheritedPropertiesAction - The id of the action
type NodeInheritedPropertiesAction string

const (
	NodeInheritedPropertiesActionNodeInheritedProperties NodeInheritedPropertiesAction = "nodeInheritedProperties"
)

func (e NodeInheritedPropertiesAction) ToPointer() *NodeInheritedPropertiesAction {
	return &e
}
func (e *NodeInheritedPropertiesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "nodeInheritedProperties":
		*e = NodeInheritedPropertiesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeInheritedPropertiesAction: %v", v)
	}
}

// NodeInheritedPropertiesResponseBody - Node
type NodeInheritedPropertiesResponseBody struct {
	// Result of the request
	Result NodeInheritedPropertiesResult `json:"result"`
	// The id of the action
	Action NodeInheritedPropertiesAction `json:"action"`
	// Information about the node inherited properties
	Data []components.NodeInheritedProperties `json:"data"`
}

func (o *NodeInheritedPropertiesResponseBody) GetResult() NodeInheritedPropertiesResult {
	if o == nil {
		return NodeInheritedPropertiesResult("")
	}
	return o.Result
}

func (o *NodeInheritedPropertiesResponseBody) GetAction() NodeInheritedPropertiesAction {
	if o == nil {
		return NodeInheritedPropertiesAction("")
	}
	return o.Action
}

func (o *NodeInheritedPropertiesResponseBody) GetData() []components.NodeInheritedProperties {
	if o == nil {
		return []components.NodeInheritedProperties{}
	}
	return o.Data
}

type NodeInheritedPropertiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Node
	Object *NodeInheritedPropertiesResponseBody
}

func (o *NodeInheritedPropertiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodeInheritedPropertiesResponse) GetObject() *NodeInheritedPropertiesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
