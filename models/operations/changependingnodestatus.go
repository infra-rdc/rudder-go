package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ChangePendingNodeStatusStatus - New status of the pending node
type ChangePendingNodeStatusStatus string

const (
	ChangePendingNodeStatusStatusAccepted ChangePendingNodeStatusStatus = "accepted"
	ChangePendingNodeStatusStatusRefused  ChangePendingNodeStatusStatus = "refused"
)

func (e ChangePendingNodeStatusStatus) ToPointer() *ChangePendingNodeStatusStatus {
	return &e
}
func (e *ChangePendingNodeStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accepted":
		fallthrough
	case "refused":
		*e = ChangePendingNodeStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChangePendingNodeStatusStatus: %v", v)
	}
}

type ChangePendingNodeStatusRequestBody struct {
	// New status of the pending node
	Status *ChangePendingNodeStatusStatus `json:"status,omitempty"`
}

func (o *ChangePendingNodeStatusRequestBody) GetStatus() *ChangePendingNodeStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type ChangePendingNodeStatusRequest struct {
	// Id of the target node
	NodeID      string                              `pathParam:"style=simple,explode=false,name=nodeId"`
	RequestBody *ChangePendingNodeStatusRequestBody `request:"mediaType=application/json"`
}

func (o *ChangePendingNodeStatusRequest) GetNodeID() string {
	if o == nil {
		return ""
	}
	return o.NodeID
}

func (o *ChangePendingNodeStatusRequest) GetRequestBody() *ChangePendingNodeStatusRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// ChangePendingNodeStatusResult - Result of the request
type ChangePendingNodeStatusResult string

const (
	ChangePendingNodeStatusResultSuccess ChangePendingNodeStatusResult = "success"
	ChangePendingNodeStatusResultError   ChangePendingNodeStatusResult = "error"
)

func (e ChangePendingNodeStatusResult) ToPointer() *ChangePendingNodeStatusResult {
	return &e
}
func (e *ChangePendingNodeStatusResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ChangePendingNodeStatusResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChangePendingNodeStatusResult: %v", v)
	}
}

// ChangePendingNodeStatusAction - The id of the action
type ChangePendingNodeStatusAction string

const (
	ChangePendingNodeStatusActionChangePendingNodeStatus ChangePendingNodeStatusAction = "changePendingNodeStatus"
)

func (e ChangePendingNodeStatusAction) ToPointer() *ChangePendingNodeStatusAction {
	return &e
}
func (e *ChangePendingNodeStatusAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "changePendingNodeStatus":
		*e = ChangePendingNodeStatusAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChangePendingNodeStatusAction: %v", v)
	}
}

// ChangePendingNodeStatusData - Information about the node
type ChangePendingNodeStatusData struct {
	Nodes []components.NodeFull `json:"nodes"`
}

func (o *ChangePendingNodeStatusData) GetNodes() []components.NodeFull {
	if o == nil {
		return []components.NodeFull{}
	}
	return o.Nodes
}

// ChangePendingNodeStatusResponseBody - Nodes
type ChangePendingNodeStatusResponseBody struct {
	// Result of the request
	Result ChangePendingNodeStatusResult `json:"result"`
	// The id of the action
	Action ChangePendingNodeStatusAction `json:"action"`
	// Information about the node
	Data ChangePendingNodeStatusData `json:"data"`
}

func (o *ChangePendingNodeStatusResponseBody) GetResult() ChangePendingNodeStatusResult {
	if o == nil {
		return ChangePendingNodeStatusResult("")
	}
	return o.Result
}

func (o *ChangePendingNodeStatusResponseBody) GetAction() ChangePendingNodeStatusAction {
	if o == nil {
		return ChangePendingNodeStatusAction("")
	}
	return o.Action
}

func (o *ChangePendingNodeStatusResponseBody) GetData() ChangePendingNodeStatusData {
	if o == nil {
		return ChangePendingNodeStatusData{}
	}
	return o.Data
}

type ChangePendingNodeStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Nodes
	Object *ChangePendingNodeStatusResponseBody
}

func (o *ChangePendingNodeStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ChangePendingNodeStatusResponse) GetObject() *ChangePendingNodeStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
