package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetSystemUpdateResultForNodeRequest struct {
	// Id of the campaign event
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *GetSystemUpdateResultForNodeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetSystemUpdateResultForNodeRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// GetSystemUpdateResultForNodeResult - Result of the request
type GetSystemUpdateResultForNodeResult string

const (
	GetSystemUpdateResultForNodeResultSuccess GetSystemUpdateResultForNodeResult = "success"
	GetSystemUpdateResultForNodeResultError   GetSystemUpdateResultForNodeResult = "error"
)

func (e GetSystemUpdateResultForNodeResult) ToPointer() *GetSystemUpdateResultForNodeResult {
	return &e
}
func (e *GetSystemUpdateResultForNodeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetSystemUpdateResultForNodeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSystemUpdateResultForNodeResult: %v", v)
	}
}

// GetSystemUpdateResultForNodeAction - The id of the action
type GetSystemUpdateResultForNodeAction string

const (
	GetSystemUpdateResultForNodeActionGetSystemUpdateResultForNode GetSystemUpdateResultForNodeAction = "getSystemUpdateResultForNode"
)

func (e GetSystemUpdateResultForNodeAction) ToPointer() *GetSystemUpdateResultForNodeAction {
	return &e
}
func (e *GetSystemUpdateResultForNodeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getSystemUpdateResultForNode":
		*e = GetSystemUpdateResultForNodeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSystemUpdateResultForNodeAction: %v", v)
	}
}

type GetSystemUpdateResultForNodeData struct {
	Noderesult []components.CampaignEventNodeResult `json:"noderesult,omitempty"`
}

func (o *GetSystemUpdateResultForNodeData) GetNoderesult() []components.CampaignEventNodeResult {
	if o == nil {
		return nil
	}
	return o.Noderesult
}

// GetSystemUpdateResultForNodeResponseBody - Campaign event result for a node
type GetSystemUpdateResultForNodeResponseBody struct {
	// Result of the request
	Result GetSystemUpdateResultForNodeResult `json:"result"`
	// The id of the action
	Action GetSystemUpdateResultForNodeAction `json:"action"`
	// Campaign event id
	ID   *string                          `json:"id,omitempty"`
	Data GetSystemUpdateResultForNodeData `json:"data"`
}

func (o *GetSystemUpdateResultForNodeResponseBody) GetResult() GetSystemUpdateResultForNodeResult {
	if o == nil {
		return GetSystemUpdateResultForNodeResult("")
	}
	return o.Result
}

func (o *GetSystemUpdateResultForNodeResponseBody) GetAction() GetSystemUpdateResultForNodeAction {
	if o == nil {
		return GetSystemUpdateResultForNodeAction("")
	}
	return o.Action
}

func (o *GetSystemUpdateResultForNodeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSystemUpdateResultForNodeResponseBody) GetData() GetSystemUpdateResultForNodeData {
	if o == nil {
		return GetSystemUpdateResultForNodeData{}
	}
	return o.Data
}

type GetSystemUpdateResultForNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Campaign event result for a node
	Object *GetSystemUpdateResultForNodeResponseBody
}

func (o *GetSystemUpdateResultForNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSystemUpdateResultForNodeResponse) GetObject() *GetSystemUpdateResultForNodeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
