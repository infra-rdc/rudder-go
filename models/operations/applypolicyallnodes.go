package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// ApplyPolicyAllNodesResult - Result of the request
type ApplyPolicyAllNodesResult string

const (
	ApplyPolicyAllNodesResultSuccess ApplyPolicyAllNodesResult = "success"
	ApplyPolicyAllNodesResultError   ApplyPolicyAllNodesResult = "error"
)

func (e ApplyPolicyAllNodesResult) ToPointer() *ApplyPolicyAllNodesResult {
	return &e
}
func (e *ApplyPolicyAllNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = ApplyPolicyAllNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplyPolicyAllNodesResult: %v", v)
	}
}

// ApplyPolicyAllNodesAction - The id of the action
type ApplyPolicyAllNodesAction string

const (
	ApplyPolicyAllNodesActionApplyPolicyAllNodes ApplyPolicyAllNodesAction = "applyPolicyAllNodes"
)

func (e ApplyPolicyAllNodesAction) ToPointer() *ApplyPolicyAllNodesAction {
	return &e
}
func (e *ApplyPolicyAllNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "applyPolicyAllNodes":
		*e = ApplyPolicyAllNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplyPolicyAllNodesAction: %v", v)
	}
}

type ApplyPolicyAllNodesData struct {
	// Rudder id of the node
	ID *string `json:"id,omitempty"`
	// Node hostname
	Hostname *string `json:"hostname,omitempty"`
	// Result or policy application trigger
	Result *string `json:"result,omitempty"`
}

func (o *ApplyPolicyAllNodesData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ApplyPolicyAllNodesData) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *ApplyPolicyAllNodesData) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

// ApplyPolicyAllNodesResponseBody - Result
type ApplyPolicyAllNodesResponseBody struct {
	// Result of the request
	Result ApplyPolicyAllNodesResult `json:"result"`
	// The id of the action
	Action ApplyPolicyAllNodesAction `json:"action"`
	Data   []ApplyPolicyAllNodesData `json:"data"`
}

func (o *ApplyPolicyAllNodesResponseBody) GetResult() ApplyPolicyAllNodesResult {
	if o == nil {
		return ApplyPolicyAllNodesResult("")
	}
	return o.Result
}

func (o *ApplyPolicyAllNodesResponseBody) GetAction() ApplyPolicyAllNodesAction {
	if o == nil {
		return ApplyPolicyAllNodesAction("")
	}
	return o.Action
}

func (o *ApplyPolicyAllNodesResponseBody) GetData() []ApplyPolicyAllNodesData {
	if o == nil {
		return []ApplyPolicyAllNodesData{}
	}
	return o.Data
}

type ApplyPolicyAllNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Result
	Object *ApplyPolicyAllNodesResponseBody
}

func (o *ApplyPolicyAllNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ApplyPolicyAllNodesResponse) GetObject() *ApplyPolicyAllNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
