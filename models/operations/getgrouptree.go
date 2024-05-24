package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetGroupTreeResult - Result of the request
type GetGroupTreeResult string

const (
	GetGroupTreeResultSuccess GetGroupTreeResult = "success"
	GetGroupTreeResultError   GetGroupTreeResult = "error"
)

func (e GetGroupTreeResult) ToPointer() *GetGroupTreeResult {
	return &e
}
func (e *GetGroupTreeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetGroupTreeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGroupTreeResult: %v", v)
	}
}

// GetGroupTreeAction - The id of the action
type GetGroupTreeAction string

const (
	GetGroupTreeActionGetGroupTree GetGroupTreeAction = "GetGroupTree"
)

func (e GetGroupTreeAction) ToPointer() *GetGroupTreeAction {
	return &e
}
func (e *GetGroupTreeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GetGroupTree":
		*e = GetGroupTreeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGroupTreeAction: %v", v)
	}
}

// GroupCategories - Group tree
type GroupCategories struct {
}

type GetGroupTreeData struct {
	// Group tree
	GroupCategories GroupCategories `json:"groupCategories"`
}

func (o *GetGroupTreeData) GetGroupCategories() GroupCategories {
	if o == nil {
		return GroupCategories{}
	}
	return o.GroupCategories
}

// GetGroupTreeResponseBody - Groups information
type GetGroupTreeResponseBody struct {
	// Result of the request
	Result GetGroupTreeResult `json:"result"`
	// The id of the action
	Action GetGroupTreeAction `json:"action"`
	Data   GetGroupTreeData   `json:"data"`
}

func (o *GetGroupTreeResponseBody) GetResult() GetGroupTreeResult {
	if o == nil {
		return GetGroupTreeResult("")
	}
	return o.Result
}

func (o *GetGroupTreeResponseBody) GetAction() GetGroupTreeAction {
	if o == nil {
		return GetGroupTreeAction("")
	}
	return o.Action
}

func (o *GetGroupTreeResponseBody) GetData() GetGroupTreeData {
	if o == nil {
		return GetGroupTreeData{}
	}
	return o.Data
}

type GetGroupTreeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups information
	Object *GetGroupTreeResponseBody
}

func (o *GetGroupTreeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetGroupTreeResponse) GetObject() *GetGroupTreeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
