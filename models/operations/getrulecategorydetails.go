package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/models/components"
)

type GetRuleCategoryDetailsRequest struct {
	RuleCategoryID string `pathParam:"style=simple,explode=false,name=ruleCategoryId"`
}

func (o *GetRuleCategoryDetailsRequest) GetRuleCategoryID() string {
	if o == nil {
		return ""
	}
	return o.RuleCategoryID
}

// GetRuleCategoryDetailsResult - Result of the request
type GetRuleCategoryDetailsResult string

const (
	GetRuleCategoryDetailsResultSuccess GetRuleCategoryDetailsResult = "success"
	GetRuleCategoryDetailsResultError   GetRuleCategoryDetailsResult = "error"
)

func (e GetRuleCategoryDetailsResult) ToPointer() *GetRuleCategoryDetailsResult {
	return &e
}
func (e *GetRuleCategoryDetailsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetRuleCategoryDetailsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleCategoryDetailsResult: %v", v)
	}
}

// GetRuleCategoryDetailsAction - The id of the action
type GetRuleCategoryDetailsAction string

const (
	GetRuleCategoryDetailsActionGetRuleCategoryDetails GetRuleCategoryDetailsAction = "GetRuleCategoryDetails"
)

func (e GetRuleCategoryDetailsAction) ToPointer() *GetRuleCategoryDetailsAction {
	return &e
}
func (e *GetRuleCategoryDetailsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GetRuleCategoryDetails":
		*e = GetRuleCategoryDetailsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuleCategoryDetailsAction: %v", v)
	}
}

type GetRuleCategoryDetailsData struct {
	RulesCategories []components.RuleCategory `json:"rulesCategories"`
}

func (o *GetRuleCategoryDetailsData) GetRulesCategories() []components.RuleCategory {
	if o == nil {
		return []components.RuleCategory{}
	}
	return o.RulesCategories
}

// GetRuleCategoryDetailsResponseBody - Rules category information
type GetRuleCategoryDetailsResponseBody struct {
	// Result of the request
	Result GetRuleCategoryDetailsResult `json:"result"`
	// The id of the action
	Action GetRuleCategoryDetailsAction `json:"action"`
	Data   GetRuleCategoryDetailsData   `json:"data"`
}

func (o *GetRuleCategoryDetailsResponseBody) GetResult() GetRuleCategoryDetailsResult {
	if o == nil {
		return GetRuleCategoryDetailsResult("")
	}
	return o.Result
}

func (o *GetRuleCategoryDetailsResponseBody) GetAction() GetRuleCategoryDetailsAction {
	if o == nil {
		return GetRuleCategoryDetailsAction("")
	}
	return o.Action
}

func (o *GetRuleCategoryDetailsResponseBody) GetData() GetRuleCategoryDetailsData {
	if o == nil {
		return GetRuleCategoryDetailsData{}
	}
	return o.Data
}

type GetRuleCategoryDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Rules category information
	Object *GetRuleCategoryDetailsResponseBody
}

func (o *GetRuleCategoryDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRuleCategoryDetailsResponse) GetObject() *GetRuleCategoryDetailsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
