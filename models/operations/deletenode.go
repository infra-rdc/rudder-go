package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteNodeRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
	// Deletion mode to use, either move to trash ('move', default) or erase ('erase')
	Mode *components.NodeDeleteMode `default:"move" queryParam:"style=form,explode=true,name=mode"`
}

func (d DeleteNodeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeleteNodeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeleteNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *DeleteNodeRequest) GetMode() *components.NodeDeleteMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

// DeleteNodeResult - Result of the request
type DeleteNodeResult string

const (
	DeleteNodeResultSuccess DeleteNodeResult = "success"
	DeleteNodeResultError   DeleteNodeResult = "error"
)

func (e DeleteNodeResult) ToPointer() *DeleteNodeResult {
	return &e
}
func (e *DeleteNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteNodeResult: %v", v)
	}
}

// DeleteNodeAction - The id of the action
type DeleteNodeAction string

const (
	DeleteNodeActionDeleteNode DeleteNodeAction = "deleteNode"
)

func (e DeleteNodeAction) ToPointer() *DeleteNodeAction {
	return &e
}
func (e *DeleteNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleteNode":
		*e = DeleteNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteNodeAction: %v", v)
	}
}

// DeleteNodeData - Information about the node
type DeleteNodeData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *DeleteNodeData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// DeleteNodeResponseBody - Nodes
type DeleteNodeResponseBody struct {
	// Result of the request
	Result DeleteNodeResult `json:"result"`
	// The id of the action
	Action DeleteNodeAction `json:"action"`
	// Information about the node
	Data DeleteNodeData `json:"data"`
}

func (o *DeleteNodeResponseBody) GetResult() DeleteNodeResult {
	if o == nil {
		return DeleteNodeResult("")
	}
	return o.Result
}

func (o *DeleteNodeResponseBody) GetAction() DeleteNodeAction {
	if o == nil {
		return DeleteNodeAction("")
	}
	return o.Action
}

func (o *DeleteNodeResponseBody) GetData() DeleteNodeData {
	if o == nil {
		return DeleteNodeData{}
	}
	return o.Data
}

type DeleteNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *DeleteNodeResponseBody
}

func (o *DeleteNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteNodeResponse) GetObject() *DeleteNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
