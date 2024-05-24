package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type UpdateRuleCategoryRequest struct {
	RuleCategoryID     string                        `pathParam:"style=simple,explode=false,name=ruleCategoryId"`
	RuleCategoryUpdate components.RuleCategoryUpdate `request:"mediaType=application/json"`
}

func (o *UpdateRuleCategoryRequest) GetRuleCategoryID() string {
	if o == nil {
		return ""
	}
	return o.RuleCategoryID
}

func (o *UpdateRuleCategoryRequest) GetRuleCategoryUpdate() components.RuleCategoryUpdate {
	if o == nil {
		return components.RuleCategoryUpdate{}
	}
	return o.RuleCategoryUpdate
}

// UpdateRuleCategoryResult - Result of the request
type UpdateRuleCategoryResult string

const (
	UpdateRuleCategoryResultSuccess UpdateRuleCategoryResult = "success"
	UpdateRuleCategoryResultError   UpdateRuleCategoryResult = "error"
)

func (e UpdateRuleCategoryResult) ToPointer() *UpdateRuleCategoryResult {
	return &e
}
func (e *UpdateRuleCategoryResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = UpdateRuleCategoryResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateRuleCategoryResult: %v", v)
	}
}

// UpdateRuleCategoryAction - The id of the action
type UpdateRuleCategoryAction string

const (
	UpdateRuleCategoryActionUpdateRuleCategory UpdateRuleCategoryAction = "UpdateRuleCategory"
)

func (e UpdateRuleCategoryAction) ToPointer() *UpdateRuleCategoryAction {
	return &e
}
func (e *UpdateRuleCategoryAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UpdateRuleCategory":
		*e = UpdateRuleCategoryAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateRuleCategoryAction: %v", v)
	}
}

type UpdateRuleCategoryData struct {
	RuleCategories []components.RuleCategory `json:"ruleCategories"`
}

func (o *UpdateRuleCategoryData) GetRuleCategories() []components.RuleCategory {
	if o == nil {
		return []components.RuleCategory{}
	}
	return o.RuleCategories
}

// UpdateRuleCategoryResponseBody - Rules category information
type UpdateRuleCategoryResponseBody struct {
	// Result of the request
	Result UpdateRuleCategoryResult `json:"result"`
	// The id of the action
	Action UpdateRuleCategoryAction `json:"action"`
	Data   UpdateRuleCategoryData   `json:"data"`
}

func (o *UpdateRuleCategoryResponseBody) GetResult() UpdateRuleCategoryResult {
	if o == nil {
		return UpdateRuleCategoryResult("")
	}
	return o.Result
}

func (o *UpdateRuleCategoryResponseBody) GetAction() UpdateRuleCategoryAction {
	if o == nil {
		return UpdateRuleCategoryAction("")
	}
	return o.Action
}

func (o *UpdateRuleCategoryResponseBody) GetData() UpdateRuleCategoryData {
	if o == nil {
		return UpdateRuleCategoryData{}
	}
	return o.Data
}

type UpdateRuleCategoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Rules category information
	Object *UpdateRuleCategoryResponseBody
}

func (o *UpdateRuleCategoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateRuleCategoryResponse) GetObject() *UpdateRuleCategoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
