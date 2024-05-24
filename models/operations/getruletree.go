package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

// GetRuleTreeResult - Result of the request
type GetRuleTreeResult string

const (
	GetRuleTreeResultSuccess GetRuleTreeResult = "success"
	GetRuleTreeResultError   GetRuleTreeResult = "error"
)

func (e GetRuleTreeResult) ToPointer() *GetRuleTreeResult {
	return &e
}
func (e *GetRuleTreeResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetRuleTreeResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleTreeResult: %v", v)
	}
}

// GetRuleTreeAction - The id of the action
type GetRuleTreeAction string

const (
	GetRuleTreeActionGetRuleTree GetRuleTreeAction = "GetRuleTree"
)

func (e GetRuleTreeAction) ToPointer() *GetRuleTreeAction {
	return &e
}
func (e *GetRuleTreeAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GetRuleTree":
		*e = GetRuleTreeAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleTreeAction: %v", v)
	}
}

// RuleCategories - Rule tree
type RuleCategories struct {
}

type GetRuleTreeData struct {
	// Rule tree
	RuleCategories RuleCategories `json:"ruleCategories"`
}

func (o *GetRuleTreeData) GetRuleCategories() RuleCategories {
	if o == nil {
		return RuleCategories{}
	}
	return o.RuleCategories
}

// GetRuleTreeResponseBody - Rules information
type GetRuleTreeResponseBody struct {
	// Result of the request
	Result GetRuleTreeResult `json:"result"`
	// The id of the action
	Action GetRuleTreeAction `json:"action"`
	Data   GetRuleTreeData   `json:"data"`
}

func (o *GetRuleTreeResponseBody) GetResult() GetRuleTreeResult {
	if o == nil {
		return GetRuleTreeResult("")
	}
	return o.Result
}

func (o *GetRuleTreeResponseBody) GetAction() GetRuleTreeAction {
	if o == nil {
		return GetRuleTreeAction("")
	}
	return o.Action
}

func (o *GetRuleTreeResponseBody) GetData() GetRuleTreeData {
	if o == nil {
		return GetRuleTreeData{}
	}
	return o.Data
}

type GetRuleTreeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Rules information
	Object *GetRuleTreeResponseBody
}

func (o *GetRuleTreeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRuleTreeResponse) GetObject() *GetRuleTreeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
