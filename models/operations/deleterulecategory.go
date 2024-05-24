package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type DeleteRuleCategoryRequest struct {
	RuleCategoryID string `pathParam:"style=simple,explode=false,name=ruleCategoryId"`
}

func (o *DeleteRuleCategoryRequest) GetRuleCategoryID() string {
	if o == nil {
		return ""
	}
	return o.RuleCategoryID
}

// DeleteRuleCategoryResult - Result of the request
type DeleteRuleCategoryResult string

const (
	DeleteRuleCategoryResultSuccess DeleteRuleCategoryResult = "success"
	DeleteRuleCategoryResultError   DeleteRuleCategoryResult = "error"
)

func (e DeleteRuleCategoryResult) ToPointer() *DeleteRuleCategoryResult {
	return &e
}
func (e *DeleteRuleCategoryResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteRuleCategoryResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteRuleCategoryResult: %v", v)
	}
}

// DeleteRuleCategoryAction - The id of the action
type DeleteRuleCategoryAction string

const (
	DeleteRuleCategoryActionDeleteRuleCategory DeleteRuleCategoryAction = "DeleteRuleCategory"
)

func (e DeleteRuleCategoryAction) ToPointer() *DeleteRuleCategoryAction {
	return &e
}
func (e *DeleteRuleCategoryAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DeleteRuleCategory":
		*e = DeleteRuleCategoryAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteRuleCategoryAction: %v", v)
	}
}

type DeleteRuleCategoryData struct {
	GroupCategories []components.RuleCategory `json:"groupCategories"`
}

func (o *DeleteRuleCategoryData) GetGroupCategories() []components.RuleCategory {
	if o == nil {
		return []components.RuleCategory{}
	}
	return o.GroupCategories
}

// DeleteRuleCategoryResponseBody - Groups category information
type DeleteRuleCategoryResponseBody struct {
	// Result of the request
	Result DeleteRuleCategoryResult `json:"result"`
	// The id of the action
	Action DeleteRuleCategoryAction `json:"action"`
	Data   DeleteRuleCategoryData   `json:"data"`
}

func (o *DeleteRuleCategoryResponseBody) GetResult() DeleteRuleCategoryResult {
	if o == nil {
		return DeleteRuleCategoryResult("")
	}
	return o.Result
}

func (o *DeleteRuleCategoryResponseBody) GetAction() DeleteRuleCategoryAction {
	if o == nil {
		return DeleteRuleCategoryAction("")
	}
	return o.Action
}

func (o *DeleteRuleCategoryResponseBody) GetData() DeleteRuleCategoryData {
	if o == nil {
		return DeleteRuleCategoryData{}
	}
	return o.Data
}

type DeleteRuleCategoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Groups category information
	Object *DeleteRuleCategoryResponseBody
}

func (o *DeleteRuleCategoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteRuleCategoryResponse) GetObject() *DeleteRuleCategoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
