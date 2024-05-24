package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

type GetNodesStatusRequest struct {
	// Comma separated list of node Ids
	Ids string `default:"default" queryParam:"style=form,explode=true,name=ids"`
}

func (g GetNodesStatusRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetNodesStatusRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetNodesStatusRequest) GetIds() string {
	if o == nil {
		return ""
	}
	return o.Ids
}

// GetNodesStatusResult - Result of the request
type GetNodesStatusResult string

const (
	GetNodesStatusResultSuccess GetNodesStatusResult = "success"
	GetNodesStatusResultError   GetNodesStatusResult = "error"
)

func (e GetNodesStatusResult) ToPointer() *GetNodesStatusResult {
	return &e
}
func (e *GetNodesStatusResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetNodesStatusResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesStatusResult: %v", v)
	}
}

// GetNodesStatusAction - The id of the action
type GetNodesStatusAction string

const (
	GetNodesStatusActionGetNodesStatus GetNodesStatusAction = "getNodesStatus"
)

func (e GetNodesStatusAction) ToPointer() *GetNodesStatusAction {
	return &e
}
func (e *GetNodesStatusAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getNodesStatus":
		*e = GetNodesStatusAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesStatusAction: %v", v)
	}
}

// GetNodesStatusStatus - Status of the node
type GetNodesStatusStatus string

const (
	GetNodesStatusStatusPending  GetNodesStatusStatus = "pending"
	GetNodesStatusStatusAccepted GetNodesStatusStatus = "accepted"
	GetNodesStatusStatusDeleted  GetNodesStatusStatus = "deleted"
	GetNodesStatusStatusUnknown  GetNodesStatusStatus = "unknown"
)

func (e GetNodesStatusStatus) ToPointer() *GetNodesStatusStatus {
	return &e
}
func (e *GetNodesStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "accepted":
		fallthrough
	case "deleted":
		fallthrough
	case "unknown":
		*e = GetNodesStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNodesStatusStatus: %v", v)
	}
}

type GetNodesStatusNodes struct {
	// Unique id of the node
	ID string `json:"id"`
	// Status of the node
	Status GetNodesStatusStatus `json:"status"`
}

func (o *GetNodesStatusNodes) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetNodesStatusNodes) GetStatus() GetNodesStatusStatus {
	if o == nil {
		return GetNodesStatusStatus("")
	}
	return o.Status
}

// GetNodesStatusData - List of nodeId and associated status
type GetNodesStatusData struct {
	// List of nodeId and associated status
	Nodes []GetNodesStatusNodes `json:"nodes"`
}

func (o *GetNodesStatusData) GetNodes() []GetNodesStatusNodes {
	if o == nil {
		return []GetNodesStatusNodes{}
	}
	return o.Nodes
}

// GetNodesStatusResponseBody - nodes
type GetNodesStatusResponseBody struct {
	// Result of the request
	Result GetNodesStatusResult `json:"result"`
	// The id of the action
	Action GetNodesStatusAction `json:"action"`
	// List of nodeId and associated status
	Data GetNodesStatusData `json:"data"`
}

func (o *GetNodesStatusResponseBody) GetResult() GetNodesStatusResult {
	if o == nil {
		return GetNodesStatusResult("")
	}
	return o.Result
}

func (o *GetNodesStatusResponseBody) GetAction() GetNodesStatusAction {
	if o == nil {
		return GetNodesStatusAction("")
	}
	return o.Action
}

func (o *GetNodesStatusResponseBody) GetData() GetNodesStatusData {
	if o == nil {
		return GetNodesStatusData{}
	}
	return o.Data
}

type GetNodesStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// nodes
	Object *GetNodesStatusResponseBody
}

func (o *GetNodesStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetNodesStatusResponse) GetObject() *GetNodesStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
