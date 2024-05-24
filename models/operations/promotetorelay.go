package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type PromoteToRelayRequest struct {
	// Id of the target node
	NodeID string `pathParam:"style=simple,explode=false,name=nodeId"`
}

func (o *PromoteToRelayRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

// PromoteToRelayAction - The id of the action
type PromoteToRelayAction string

const (
	PromoteToRelayActionPromoteToRelay PromoteToRelayAction = "promoteToRelay"
)

func (e PromoteToRelayAction) ToPointer() *PromoteToRelayAction {
	return &e
}
func (e *PromoteToRelayAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "promoteToRelay":
		*e = PromoteToRelayAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PromoteToRelayAction: %v", v)
	}
}

// PromoteToRelayResult - Result of the request
type PromoteToRelayResult string

const (
	PromoteToRelayResultSuccess PromoteToRelayResult = "success"
	PromoteToRelayResultError   PromoteToRelayResult = "error"
)

func (e PromoteToRelayResult) ToPointer() *PromoteToRelayResult {
	return &e
}
func (e *PromoteToRelayResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = PromoteToRelayResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PromoteToRelayResult: %v", v)
	}
}

// PromoteToRelayResponseBody - Promote response
type PromoteToRelayResponseBody struct {
	// The id of the action
	Action PromoteToRelayAction `json:"action"`
	// Result of the request
	Result PromoteToRelayResult `json:"result"`
	// Success or error message
	Data string `json:"data"`
}

func (o *PromoteToRelayResponseBody) GetAction() PromoteToRelayAction {
	if o == nil {
		return PromoteToRelayAction("")
	}
	return o.Action
}

func (o *PromoteToRelayResponseBody) GetResult() PromoteToRelayResult {
	if o == nil {
		return PromoteToRelayResult("")
	}
	return o.Result
}

func (o *PromoteToRelayResponseBody) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

type PromoteToRelayResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Promote response
	Object *PromoteToRelayResponseBody
}

func (o *PromoteToRelayResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PromoteToRelayResponse) GetObject() *PromoteToRelayResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
