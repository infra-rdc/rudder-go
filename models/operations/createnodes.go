package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// CreateNodesResult - Result of the request
type CreateNodesResult string

const (
	CreateNodesResultSuccess CreateNodesResult = "success"
	CreateNodesResultError   CreateNodesResult = "error"
)

func (e CreateNodesResult) ToPointer() *CreateNodesResult {
	return &e
}
func (e *CreateNodesResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = CreateNodesResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateNodesResult: %v", v)
	}
}

// CreateNodesAction - The id of the action
type CreateNodesAction string

const (
	CreateNodesActionCreateNode CreateNodesAction = "createNode"
)

func (e CreateNodesAction) ToPointer() *CreateNodesAction {
	return &e
}
func (e *CreateNodesAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "createNode":
		*e = CreateNodesAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateNodesAction: %v", v)
	}
}

type CreateNodesData struct {
	Created []string `json:"created"`
	Failed  []string `json:"failed"`
}

func (o *CreateNodesData) GetCreated() []string {
	if o == nil {
		return []string{}
	}
	return o.Created
}

func (o *CreateNodesData) GetFailed() []string {
	if o == nil {
		return []string{}
	}
	return o.Failed
}

// CreateNodesResponseBody - Creation information
type CreateNodesResponseBody struct {
	// Result of the request
	Result CreateNodesResult `json:"result"`
	// The id of the action
	Action CreateNodesAction `json:"action"`
	Data   CreateNodesData   `json:"data"`
}

func (o *CreateNodesResponseBody) GetResult() CreateNodesResult {
	if o == nil {
		return CreateNodesResult("")
	}
	return o.Result
}

func (o *CreateNodesResponseBody) GetAction() CreateNodesAction {
	if o == nil {
		return CreateNodesAction("")
	}
	return o.Action
}

func (o *CreateNodesResponseBody) GetData() CreateNodesData {
	if o == nil {
		return CreateNodesData{}
	}
	return o.Data
}

type CreateNodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Creation information
	Object *CreateNodesResponseBody
}

func (o *CreateNodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateNodesResponse) GetObject() *CreateNodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
